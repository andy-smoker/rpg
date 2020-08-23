package main

import (
	"log"
	"net/http"
	"server/database"
	"server/handlers"

	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	db := database.NewDB()
	db.ConfigToml()
	log.Fatal(srv.ListenAndServe())
}

// Router .
func router(r *mux.Router) {

	//r.HandleFunc("/characters", handlers.GetAllCharshits).Methods("GET")
	//r.HandleFunc("/auth", auth.AuthHandler).Methods("POST", "GET")
	//r.HandleFunc("/auth/reg", auth.Register).Methods("POST")
	//r.HandleFunc("/{core}/chars", handlers.SWgetAllChars)
	//r.HandleFunc("/{core}/chars/ch{id}", handlers.GetChar)
	//r.HandleFunc("/{core}/chars/ch{id}/del", nil)
	//r.HandleFunc("/{core}/chars/add", handlers.AddCharshit)
	r.HandleFunc("/{core}/chars/ch{id}", handlers.Midle)

}
