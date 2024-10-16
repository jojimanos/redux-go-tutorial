package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/jojimanos/redux-go-tutorial/models"
    "github.com/jojimanos/redux-go-tutorial/utils"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    // Hash the password before storing it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    collection := utils.GetCollection("users")
    _, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    collection := utils.GetCollection("users")

    username := r.URL.Query().Get("username")

    err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
    var loginUser models.User
    var dbUser models.User

    _ = json.NewDecoder(r.Body).Decode(&loginUser)

    collection := utils.GetCollection("users")
    err := collection.FindOne(context.TODO(), bson.M{"username": loginUser.Username}).Decode(&dbUser)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Compare hashed password with the provided password
    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password))
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(dbUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    collection := utils.GetCollection("users")
    filter := bson.M{"username": user.Username}

    update := bson.M{
        "$set": bson.M{
            "email": user.Email,
            // Update other fields if necessary
        },
    }

    _, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode("User updated")
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
