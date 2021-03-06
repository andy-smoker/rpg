package main

import (
	"log"
	"net/http"
	"server/auth"
	"server/savage"

	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//r.Use(auth.Middleware)
	SWroutes(r)
	mainRoute(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func mainRoute(r *mux.Router) {

	r.HandleFunc("/auth", auth.Login).Methods("POST")
	r.HandleFunc("/reg", auth.Register).Methods("POST")
}

// SWroutes .
func SWroutes(r *mux.Router) {
	prefix := "/sw"
	r.HandleFunc(prefix+"/chars", savage.GetAllChars).Methods("GET")
	r.HandleFunc(prefix+"/chars/{id}", savage.CharID).Methods("GET", "PUT", "DELETE")
	r.HandleFunc(prefix+"/chars/add", savage.AddChar).Methods("POST")
	r.HandleFunc(prefix+"/items/{id}", savage.ItemMidleware)
	/*
		r.HandleFunc(prefix+"/abilities", nil)
			r.HandleFunc(prefix+"/abilities/{id}", nil)
			r.HandleFunc(prefix+"/traits", nil)
			r.HandleFunc(prefix+"/traits/{id}", nil)
			r.HandleFunc(prefix+"/flaws", nil)
			r.HandleFunc(prefix+"/flaws/{id}", nil)
			r.HandleFunc(prefix+"/races", savage.GetAllRaces)

	*/
}
