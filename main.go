package main

import (
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	env "github.com/joho/godotenv"

	"postsbackend/routes"
	"postsbackend/db"
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
	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	http.Handle("/", corsOpt.Handler(s.Router))
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.Migrate()
	server := Server{}
	server.Initialize(os.Getenv("APP_URL"))
	server.Run()
}