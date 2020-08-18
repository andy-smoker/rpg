package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"rpg/server/models"
	"rpg/server/models/core"
	"text/template"

	"github.com/gorilla/mux"
	// import posgresql famework
	_ "github.com/lib/pq"
)

// GetAllCharshits .
func GetAllCharshits(w http.ResponseWriter, r *http.Request) {
	chars := []models.CharShit{}
	sqlConf := "user=rest password=rest dbname=rpg sslmode=disable"
	db, err := sql.Open("postgres", sqlConf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from chars")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		arr := []uint8{}
		chsh := models.CharShit{}
		c := core.Char{}
		err := rows.Scan(&chsh.ID, &chsh.Name, &arr, &c.Rank)
		if err != nil {
			log.Fatal(err)
			continue
		}
		for i := range arr {
			c.Skills = append(c.Skills, i)
		}

		chsh.Core = c
		chars = append(chars, chsh)
		fmt.Println(c.Skills)
	}
	data, err := json.Marshal(chars)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// GetCharhit .
func GetCharhit(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	chsh := models.CharShit{}
	c := core.Char{}
	sqlConf := "user=rest password=rest dbname=rpg sslmode=disable"
	db, err := sql.Open("postgres", sqlConf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from chars where id = $1", vars["id"])
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		arr := []uint8{}
		err := rows.Scan(&chsh.ID, &chsh.Name, &arr, &c.Rank)
		if err != nil {
			log.Fatal(err)
			continue
		}
		for i := range arr {
			c.Skills = append(c.Skills, i)
		}
	}
	chsh.Core = c

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(chsh)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	w.Write(data)
}

// HH .
func HH(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("add.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddCharshit .
func AddCharshit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	v := models.CharShit{}
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := db.Exec("insert into charshit(name)values($1)", v.Name)
	if err != nil {
		log.Println(err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "Успешно добвален: ", lastID)
}

// DelCharshit .
func DelCharshit(w http.ResponseWriter, r *http.Request) {

}

// DBHandler .
func DBHandler() []models.CharShit {
	var chars = []models.CharShit{}
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()
	rows, err := db.Query("select * from charshit ")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		chsh := models.CharShit{}
		err := rows.Scan(&chsh.ID, &chsh.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		chars = append(chars, chsh)
	}
	return chars
}
