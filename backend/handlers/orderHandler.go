package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jojimanos/redux-go-tutorial/models"
	"github.com/jojimanos/redux-go-tutorial/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	var user models.User

	// Decode the JSON request body into the 'order' variable
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Printf("Error decoding request body: %v", err) // Log decoding error
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if the provided user ID is a valid MongoDB ObjectID
	if order.UserID.IsZero() {
		log.Println("Error: Invalid user ID (zero value)") // Log invalid user ID error
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get the users collection
	usersCollection := utils.GetCollection("users")

	// Check if the user exists in the "users" collection
	err = usersCollection.FindOne(context.TODO(), bson.M{"_id": order.UserID}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user with ID %s: %v", order.UserID.Hex(), err) // Log error finding user
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}

	// If the user exists, proceed to insert the order
	ordersCollection := utils.GetCollection("orders")
	_, err = ordersCollection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Printf("Error creating order: %v", err) // Log error during order creation
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	// Return the created order as JSON
	json.NewEncoder(w).Encode(order)
}

// func GetUser(w http.ResponseWriter, r *http.Request) {
// var user models.User
// collection := utils.GetCollection("users")
//
// username := r.URL.Query().Get("username")
//
// err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
// if err == mongo.ErrNoDocuments {
// http.Error(w, "User not found", http.StatusNotFound)
// return
// }
//
// json.NewEncoder(w).Encode(user)
// }
//
// func LoginUser(w http.ResponseWriter, r *http.Request) {
// var loginUser models.User
// var dbUser models.User
//
// _ = json.NewDecoder(r.Body).Decode(&loginUser)
//
// collection := utils.GetCollection("users")
// err := collection.FindOne(context.TODO(), bson.M{"username": loginUser.Username}).Decode(&dbUser)
// if err == mongo.ErrNoDocuments {
// http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// return
// }
//
// Compare hashed password with the provided password
// err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password))
// if err != nil {
// http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// return
// }
//
// json.NewEncoder(w).Encode(dbUser)
// }
//
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// var user models.User
// _ = json.NewDecoder(r.Body).Decode(&user)
//
// collection := utils.GetCollection("users")
// filter := bson.M{"username": user.Username}
//
// update := bson.M{
// "$set": bson.M{
// "email": user.Email,
// Update other fields if necessary
// },
// }
//
// _, err := collection.UpdateOne(context.TODO(), filter, update)
// if err != nil {
// http.Error(w, "Error updating user", http.StatusInternalServerError)
// return
// }
//
// json.NewEncoder(w).Encode("User updated")
// }
//
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// username := r.URL.Query().Get("username")
//
// collection := utils.GetCollection("users")
// _, err := collection.DeleteOne(context.TODO(), bson.M{"username": username})
// if err != nil {
// http.Error(w, "Error deleting user", http.StatusInternalServerError)
// return
// }
//
// json.NewEncoder(w).Encode("User deleted")
// }
//
