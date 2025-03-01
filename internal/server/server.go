package server

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	RegisterRoutes(router)

	
	return router
}