package routes

import (
	. "postsbackend/handlers"
	. "postsbackend/middlewares"
)

func registerReactionsRoutes() {
	router.HandleFunc("/reactions/{post}", Auth(GetCommentsByPost)).Methods("GET")           // Get All or paginate
	router.HandleFunc("/reactions", Auth(CreateComment)).Methods("POST")                     // Create
	router.HandleFunc("/reactions/latest/{post}", Auth(GetLastCommentByPost)).Methods("GET") // Get All or paginate
}
