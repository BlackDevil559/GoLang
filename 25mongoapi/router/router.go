package router

import (
	"github.com/BlackDevil559/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router:=mux.NewRouter()
	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/add",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/update/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}",controller.DeleteOneRecord).Methods("DELETE")
	router.HandleFunc("/api/deleteall",controller.DeleteManyRecord).Methods("DELETE")
	return router
}