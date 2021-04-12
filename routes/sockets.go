package routes

import (
	. "../sockets"
)

func registerSocketRoutes() {
	router.HandleFunc("/ws", CreateWebSocket)
}