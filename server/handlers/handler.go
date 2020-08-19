package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rpg/server/auth"
	"rpg/server/config"
	"rpg/server/models"
	"rpg/server/models/swmodels"

	"github.com/gorilla/mux"
	//
	_ "github.com/lib/pq"
)

var sessions = []auth.Session{
	auth.Session{User: "Igor", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTc3MTkwODIsInVzZXJfaWQiOjF9.YIzbGH6IJp6u8BI3hdT4U1PJeMcQW--FCdvkQcy_TX4"},
}

// CheckToken method for compare tokens
func CheckToken(r *http.Request) error {
	token := r.Header.Get("Token")
	user := r.Header.Get("User")
	for _, s := range sessions {
		if s.User == user && s.Token == token {
			return nil
		}
	}
	return fmt.Errorf("auth failed")
}

// AddToDB .
func AddToDB(data interface{}) {

	db, err := sql.Open(config.DBConnect())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// ....
}

// GetAllCharshits send all charshit response
func GetAllCharshits(w http.ResponseWriter, r *http.Request) {
	chars := []swmodels.SWChar{}
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
		sw := swmodels.SWChar{}
		err := rows.Scan(&sw.ID, &sw.Name, &arr, &sw.Rank)
		if err != nil {
			log.Fatal(err)
			continue
		}
		for i := range arr {
			sw.Skills = append(sw.Skills, i)
		}

		chars = append(chars, sw)
		fmt.Println(sw.Skills)
	}
	data, err := json.Marshal(chars)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
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

// GetChar - get by id
func GetChar(w http.ResponseWriter, r *http.Request) {
	var data []byte
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	switch vars["core"] {
	case "sw":
		data = SWgetChar(vars["id"])
		break
	case "DnD":
		fmt.Println("DnD")
	}

	w.Write(data)

}
