package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jojimanos/redux-go-tutorial/models"
	"github.com/jojimanos/redux-go-tutorial/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with a secure secret

// JWT Claims structure
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the incoming JSON request body into the user model
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	collection := utils.GetCollection("users")

	// Check if the username or email already exists in the database
	var existingUser models.User
	err = collection.FindOne(context.TODO(), bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}).Decode(&existingUser)

	if err == nil {
		// If a user with the same username or email exists, return an error
		http.Error(w, "Username or email already in use", http.StatusBadRequest)
		return
	} else if err != mongo.ErrNoDocuments {
		// If an error other than no documents occurs during the lookup
		http.Error(w, "Error checking existing users", http.StatusInternalServerError)
		return
	}

	// Ensure the password meets the minimum security requirements
	if !isValidPassword(user.Password) {
		http.Error(w, "Password must be at least 8 characters long, include one symbol, one capital letter, and one number", http.StatusBadRequest)
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Insert the new user into the database
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Respond with the newly created user (without returning the password)
	user.Password = "" // Clear password field before sending the response
	json.NewEncoder(w).Encode(user)
}

// Helper function to validate the password
func isValidPassword(password string) bool {
	var hasMinLen, hasUpper, hasNumber, hasSpecial bool
	const minPassLength = 8
	specialChars := "!@#$%^&*"

	if len(password) >= minPassLength {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			if contains(specialChars, char) {
				hasSpecial = true
			}
		}
	}

	return hasMinLen && hasUpper && hasNumber && hasSpecial
}

// Helper function to check if a character is in a string of special characters
func contains(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

// Middleware to extract and validate the JWT token
func validateToken(r *http.Request) (*jwt.Token, error) {
	// Get the token from the Authorization header
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return nil, http.ErrNoCookie // No token found
	}

	// Token comes in the format "Bearer <token>", split it
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return nil, http.ErrNoCookie // Invalid token format
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenParts[1], &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNoCookie
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err // Token is invalid
	}

	return token, nil
}

// GetUser handler with JWT validation
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token
	token, err := validateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims from the token if valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Access claims data if needed
		// Example: username := claims.Username
		_ = claims // Placeholder for any claim usage
	} else {
		http.Error(w, "Unauthorized: token claims invalid", http.StatusUnauthorized)
		return
	}

	// Get the username from query params
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// MongoDB collection
	collection := utils.GetCollection("users")

	// Find the user by username
	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUser handler with JWT validation and extracting username from JWT claims
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token
	token, err := validateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims from the token if valid
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		http.Error(w, "Unauthorized: token claims invalid", http.StatusUnauthorized)
		return
	}

	// Get the username from the JWT claims
	username := claims.Username
	if username == "" {
		http.Error(w, "Username not found in token", http.StatusBadRequest)
		return
	}

	// MongoDB collection
	collection := utils.GetCollection("users")

	// Find the user by username from the claims
	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUsers handler with JWT validation
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Validate the JWT token
	token, err := validateToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	// Extract claims from the token if valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Optionally, use token claims like username, roles, etc.
		_ = claims // Placeholder for claims usage, such as user role verification
	} else {
		http.Error(w, "Unauthorized: token claims invalid", http.StatusUnauthorized)
		return
	}

	// MongoDB collection
	collection := utils.GetCollection("users")

	// Find all users (empty filter `{}`)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	// Prepare a slice to store users
	var users []models.User

	// Iterate through the cursor and decode each user
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, "Error decoding user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Check if there were any errors during the iteration
	if err := cursor.Err(); err != nil {
		http.Error(w, "Error iterating through users", http.StatusInternalServerError)
		return
	}

	// Return the list of users as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	var dbUser models.User

	// Decode the login request to get the provided username and password
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	collection := utils.GetCollection("users")

	// Find user by username in the database
	err = collection.FindOne(context.TODO(), bson.M{"username": loginUser.Username}).Decode(&dbUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Compare the provided password with the hashed password stored in the database
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password))
	if err != nil {
		// If passwords don't match, return an error
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	// If password is valid, generate the JWT token
	tokenString, err := generateJWT(dbUser.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the JWT token
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Function to generate JWT token
func generateJWT(username string) (string, error) {
	// Create the JWT claims, which includes the username and expiry time
	expirationTime := time.Now().Add(24 * time.Hour) // Token is valid for 24 hours
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token using the signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Create a map to hold incoming fields (like email) rather than decoding the entire user struct
	var updateFields map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateFields)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract username from the query params
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Get the users collection
	collection := utils.GetCollection("users")

	// Find the user by username
	var existingUser models.User
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}

	// Check if the provided email is the same as the current one
	if newEmail, emailExists := updateFields["email"]; emailExists {
		if newEmail == existingUser.Email {
			http.Error(w, "The provided email is the same as the current one", http.StatusBadRequest)
			return
		}
	}

	// Check if the provided username is the same as the current one
	if newUsername, usernameExists := updateFields["username"]; usernameExists {
		if newUsername == existingUser.Username {
			http.Error(w, "The provided username is the same as the current one", http.StatusBadRequest)
			return
		}
	}

	// Prepare the update document
	update := bson.M{
		"$set": updateFields, // Use the decoded fields to update only the provided ones
	}

	// Filter to find the correct user document by username
	filter := bson.M{"username": username}

	// Perform the update operation
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User updated successfully")
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Parse request body for old and new password
	var request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	_ = json.NewDecoder(r.Body).Decode(&request)

	// Extract JWT token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
		return
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate and parse the JWT token (assuming you have a function for that)
	claims, err := utils.ParseJWT(tokenStr)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract user information from the JWT claims (for example, username or userID)
	username := claims.Username // Assuming the token contains the username

	// Retrieve the user from the database
	var user models.User
	collection := utils.GetCollection("users")
	err = collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	// Compare the old password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	if err != nil {
		// Old password is incorrect
		http.Error(w, "Current password is incorrect", http.StatusUnauthorized)
		return
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing new password", http.StatusInternalServerError)
		return
	}

	// Update the user's password in the database
	update := bson.M{
		"$set": bson.M{
			"password": string(hashedPassword),
		},
	}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"username": username}, update)
	if err != nil {
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	json.NewEncoder(w).Encode("Password updated successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	collection := utils.GetCollection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("User deleted")
}
