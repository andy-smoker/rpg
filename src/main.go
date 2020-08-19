package main

import (
	"log"
	"net/http"
	"rpg/client"
	"rpg/server"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	server.Router(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go client.RunClient()
	log.Fatal(srv.ListenAndServe())

}
