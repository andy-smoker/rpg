package server

import (
	"rpg/server/auth"

	"rpg/server/handlers"

	"github.com/gorilla/mux"
)

// Router .
func Router(r *mux.Router) {

	r.HandleFunc("/characters", handlers.GetAllCharshits).Methods("GET")
	r.HandleFunc("/auth", auth.AuthHandler).Methods("POST", "GET")
	r.HandleFunc("/auth/reg", auth.Register).Methods("POST")
	r.HandleFunc("/{core}/chars", handlers.SWgetAllChars)
	r.HandleFunc("/{core}/chars/ch{id}", handlers.GetChar)
	r.HandleFunc("/{core}/chars/ch{id}/del", nil)
	r.HandleFunc("/{core}/chars/add", handlers.AddCharshit)

}
