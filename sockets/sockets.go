package sockets

import (
	"net/http"
	"log"

	"github.com/gorilla/websocket"

	"../models"
	"../io/response"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}

func CreateWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	clients[ws] = true
	log.Println("New Socket connection", ws)
	response.Json("Connected", http.StatusOK, w)
}

func Broadcast(msg models.Message) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}