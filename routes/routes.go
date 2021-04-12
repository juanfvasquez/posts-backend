package routes

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

func RegisterRoutes() *mux.Router {
	router = mux.NewRouter()
	registerAuthRoutes()
	registerUsersRoutes()
	registerPostsRoutes()
	registerCommentsRoutes()
	registerSocketRoutes()
	return router
}