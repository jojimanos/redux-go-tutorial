package main

import (
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers" // Import the handlers package for CORS
	"github.com/gorilla/mux"
	"github.com/jojimanos/redux-go-tutorial/handlers"
	"github.com/jojimanos/redux-go-tutorial/utils"
)

func main() {
	// Connect to MongoDB
	utils.ConnectDB()

	// Set up the router
	r := mux.NewRouter()

	// Routes for CRUD operations
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/user", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user/current", handlers.GetCurrentUser).Methods("GET")
	r.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/user/change_password", handlers.ChangePassword).Methods("PUT")
	r.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/order", handlers.CreateOrder).Methods("POST")

	// CORS configuration
	corsOptions := gorillaHandlers.AllowedOrigins([]string{"*"})                             // Allow all origins for development
	corsMethods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})  // Allow specific HTTP methods
	corsHeaders := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}) // Allow specific headers

	// Combine all CORS options
	corsMiddleware := gorillaHandlers.CORS(corsOptions, corsMethods, corsHeaders)

	// Start the server with CORS middleware
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", corsMiddleware(r)))
}
