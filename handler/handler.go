package handler

import (
	"encoding/json"
	"fmt"
	"github.com/CH02z/movielibrary/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	movieService	services.MovieService
}

func newHandler(movieService services.MovieService) *Handler {
	return &Handler{movieService: movieService}
}

func (h Handler) homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome home!")
}

func (h Handler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies, err := h.movieService.GetAllMovies()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(movies)
	}
}

func (h Handler) newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)

	router.HandleFunc("/", h.homeLink)
	router.HandleFunc("/movies", h.getAllMovies).Methods("GET", "OPTIONS")

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}