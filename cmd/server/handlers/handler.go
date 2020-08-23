package handlers

import (
	"fmt"
	"net/http"
	savage "server/sevage"

	//
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Midle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//GetChar
	switch vars["core"] {
	case "sw":
		if r.Method == "GET" {
			fmt.Println(string(savage.SWgetChar(vars["id"])))
		}
		break
	}
}
