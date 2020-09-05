package main

import (
	"log"
	"net/http"
<<<<<<< HEAD
	savage "server/savage"
=======
	"server/savage"
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3

	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	SWroutes(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// SWroutes .
func SWroutes(r *mux.Router) {
	prefix := "/sw"
<<<<<<< HEAD
	r.HandleFunc(prefix+"/chars", nil)
	r.HandleFunc(prefix+"/chars/ch{id}", savage.CharID).Methods("GET", "PUT", "DELETE")
	r.HandleFunc(prefix+"/chars/add", savage.AddChar).Methods("POST")
	r.HandleFunc(prefix+"/abilities", nil)
	r.HandleFunc(prefix+"/abilities/{id}", nil)
	r.HandleFunc(prefix+"/traits", nil)
	r.HandleFunc(prefix+"/traits/{id}", nil)
	r.HandleFunc(prefix+"/flaws", nil)
	r.HandleFunc(prefix+"/flaws/{id}", nil)
	r.HandleFunc(prefix+"/races", savage.GetAllRaces)
	r.HandleFunc(prefix+"/items/{id}", nil)
=======
	r.HandleFunc(prefix+"/chars", savage.GetAllChars).Methods("GET")
	r.HandleFunc(prefix+"/chars/ch{id}", savage.CharID).Methods("GET", "PUT", "DELETE")
	r.HandleFunc(prefix+"/chars/add", savage.AddChar).Methods("POST")
	/*
		r.HandleFunc(prefix+"/abilities", nil)
		r.HandleFunc(prefix+"/abilities/{id}", nil)
		r.HandleFunc(prefix+"/traits", nil)
		r.HandleFunc(prefix+"/traits/{id}", nil)
		r.HandleFunc(prefix+"/flaws", nil)
		r.HandleFunc(prefix+"/flaws/{id}", nil)
		r.HandleFunc(prefix+"/races", savage.GetAllRaces)
		r.HandleFunc(prefix+"/items/{id}", nil)
	*/
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
}
