package main

import (
    "log"
    "net/http"
    "github.com/jojimanos/redux-go-tutorial/handlers"
    "github.com/jojimanos/redux-go-tutorial/utils"
    "github.com/gorilla/mux"
)

func main() {
    // Connect to MongoDB
    utils.ConnectDB()

    // Set up the router
    r := mux.NewRouter()

    // Routes for CRUD operations
    r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/user", handlers.GetUser).Methods("GET")
    r.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
    r.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")

    // Start the server
    log.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
