package server

import (
	"log"
	"net/http"
	"rpg/server/auth"
	"rpg/server/handlers"
	"time"

	"github.com/gorilla/mux"
)

// RunServ .
func RunServ() {
	router := mux.NewRouter()
	Router(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Router .
func Router(r *mux.Router) {
	r.HandleFunc("/", HH)
	r.HandleFunc("/characters", GetAllCharshits).Methods("GET")
	r.HandleFunc("/auth", auth.AuthHandler).Methods("POST")
	r.HandleFunc("/sw/chars", handlers.GetAllSWChars)
	r.HandleFunc("/characters/ch{id}", GetCharhit)
	r.HandleFunc("/characters/ch{id}/del", nil)
	r.HandleFunc("/characters/add", AddCharshit)

}
