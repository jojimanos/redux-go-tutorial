package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jojimanos/redux-go-tutorial/models"
	"github.com/jojimanos/redux-go-tutorial/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTranslations retrieves all translation documents from MongoDB.
func GetTranslations(w http.ResponseWriter, r *http.Request) {
	// Get the MongoDB collection
	collection := utils.GetCollection("translations")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch all documents in the collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	// Decode documents into the Translations struct
	var translations []models.Translations
	for cursor.Next(ctx) {
		var translation models.Translations
		if err := cursor.Decode(&translation); err != nil {
			http.Error(w, "Error decoding document", http.StatusInternalServerError)
			return
		}
		translations = append(translations, translation)
	}

	// Check if any documents were found
	if len(translations) == 0 {
		http.Error(w, "Translations not found", http.StatusNotFound)
		return
	}

	// Set the content type and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(translations); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

