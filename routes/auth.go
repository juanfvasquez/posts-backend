package routes

import (
	. "postsbackend/handlers"
)

func registerAuthRoutes() {
	router.HandleFunc("/auth", Login).Methods("POST")
}