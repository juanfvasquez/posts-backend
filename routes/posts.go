package routes

import (
	. "../handlers"
	. "../middlewares"
)

func registerPostsRoutes() {
	router.HandleFunc("/posts", Auth(GetPosts)).Methods("GET") // Get All or paginated
	router.HandleFunc("/posts/{id}", Auth(GetPost)).Methods("GET") // Get All or paginated
	router.HandleFunc("/posts", Auth(CreatePost)).Methods("POST") // Create
	router.HandleFunc("/posts/{user}", Auth(GetPostsByUser)).Methods("GET") // Get by user
}