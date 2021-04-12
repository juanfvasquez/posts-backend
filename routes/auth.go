package routes

import (
	. "../handlers"
)

func registerAuthRoutes() {
	router.HandleFunc("/auth", Login).Methods("POST")
}