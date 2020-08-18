package server

import (
	"fmt"
	"log"
	"net/http"
	"rpg/server/auth"
	"rpg/server/config"

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
	fmt.Println(config.DBConnect())
	log.Fatal(srv.ListenAndServe())

}

// Router .
func Router(r *mux.Router) {

	r.HandleFunc("/characters", handlers.GetAllCharshits).Methods("GET")
	r.HandleFunc("/auth", auth.AuthHandler).Methods("POST")
	r.HandleFunc("/{core}/chars", handlers.SWgetAllChars)
	r.HandleFunc("/{core}/chars/ch{id}", handlers.GetChar)
	r.HandleFunc("/{core}/chars/ch{id}/del", nil)
	r.HandleFunc("/{core}/chars/add", handlers.AddCharshit)

}
