package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

// SPAHandler struct
type SPAHandler struct {
	StaticPath string
	IndexPath  string
}

// NewSPAHandler make new SPAHandler struct
func NewSPAHandler(index string) SPAHandler {
	return SPAHandler{
		StaticPath: "client/static",
		IndexPath:  index,
	}
}

// func for SPAHandler realisation as Handler in router
func (spa SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		log.Println(err)
		return
	}

	path = filepath.Join(spa.StaticPath, path)

	// check for files
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		//
		http.ServeFile(w, r, filepath.Join(spa.StaticPath, spa.IndexPath))
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	http.FileServer(http.Dir(spa.StaticPath)).ServeHTTP(w, r)
}

// RunClient public finc for running client
func main() {
	r := mux.NewRouter()
	router(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":1234",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func router(r *mux.Router) {
	r.PathPrefix("/").Handler(NewSPAHandler("index.html"))
}
