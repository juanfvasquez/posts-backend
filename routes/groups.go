package routes

import (
	. "postsbackend/handlers"
	. "postsbackend/middlewares"
)

func registerGroupsRoutes() {
	router.HandleFunc("/groups/{post}", Auth(GetCommentsByPost)).Methods("GET")           // Get All or paginate
	router.HandleFunc("/groups", Auth(CreateComment)).Methods("POST")                     // Create
	router.HandleFunc("/groups/latest/{post}", Auth(GetLastCommentByPost)).Methods("GET") // Get All or paginate
}
