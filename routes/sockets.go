package routes

import (
	. "postsbackend/sockets"
)

func registerSocketRoutes() {
	router.HandleFunc("/ws", CreateWebSocket)
}