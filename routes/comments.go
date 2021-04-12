package routes

import (
	. "../handlers"
	. "../middlewares"
)

func registerCommentsRoutes() {
	router.HandleFunc("/comments/{post}", Auth(GetCommentsByPost)).Methods("GET") // Get All or paginate
	router.HandleFunc("/comments", Auth(CreateComment)).Methods("POST") // Create
	router.HandleFunc("/comments/latest/{post}", Auth(GetLastCommentByPost)).Methods("GET") // Get All or paginate
}