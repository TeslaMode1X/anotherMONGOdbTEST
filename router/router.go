package router

import (
	"github.com/gorilla/mux"
	"mongoAPI/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movies/delete", controller.DeleteAllRecords).Methods("DELETE")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneRecord).Methods("DELETE")

	return r
}
