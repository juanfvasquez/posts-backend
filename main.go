package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"./routes"
	"./db"
)

type Server struct {
	Router 	*mux.Router
	Addr 		string
}

func (s *Server) Initialize(addr string) {
	s.Router = routes.RegisterRoutes()
	s.Addr = addr
}

func (s *Server) Run() {
	log.Println("Server running on", s.Addr)
	http.Handle("/", s.Router)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func main() {
	db.Migrate()
	server := Server{}
	server.Initialize(":5050")
	server.Run()
}