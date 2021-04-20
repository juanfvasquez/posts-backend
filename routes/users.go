package routes

import (
	. "postsbackend/handlers"
	. "postsbackend/middlewares"
)

func registerUsersRoutes() {
	router.HandleFunc("/users", Auth(GetUsers)).Methods("GET") // Get All or search
	router.HandleFunc("/users/{id}", Auth(GetUser)).Methods("GET") // Get All or search
	router.HandleFunc("/users", CreateUser).Methods("POST") // Create
	// router.HandleFunc("/users", nil).Methods("PATCH") // Update
}