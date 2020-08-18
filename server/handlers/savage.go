package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rpg/server/config"
	"rpg/server/models/swmodels"

	"rpg/server/models"

	//pq - postgres framework
	_ "github.com/lib/pq"
)

// SWgetAllChars send all charshit response
func SWgetAllChars(w http.ResponseWriter, r *http.Request) {
	err := CheckToken(r)
	if err != nil {
		log.Println(err)
		return
	}

	chars := []models.CharShit{}
	db, err := sql.Open(config.DBConnect())
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`select id from savage_world_chars, sw_char_stats`)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		char := models.CharShit{}
		//core := core.Char{}
		err := rows.Scan(&char.ID)
		if err != nil {
			log.Fatal(err)
			continue
		}
		chars = append(chars, char)
	}
	data, err := json.Marshal(chars)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// SWgetChar method for marshal and get data
func SWgetChar(id string) []byte {

	chsh := models.CharShit{}
	sw := swmodels.Char{}
	db, err := sql.Open(config.DBConnect())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from chars where id = $1", id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		arr := []uint8{}
		err := rows.Scan(&chsh.ID, &chsh.Name, &arr, &sw.Rank)
		if err != nil {
			log.Println(err)
			continue
		}
		for i := range arr {
			sw.Skills = append(sw.Skills, i)
		}
	}
	chsh.Core = sw
	data, err := json.Marshal(chsh)
	if err != nil {
		log.Println(err)
	}
	return data
}
