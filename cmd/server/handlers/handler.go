package handlers

<<<<<<< HEAD
import (
	"fmt"
	"net/http"
	savage "server/sevage"

	//
	"github.com/gorilla/mux"
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
=======
//
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
