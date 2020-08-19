package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rpg/server/config"
	"rpg/server/models/swmodels"

	"rpg/server/models"
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

	rows, err := db.Query(`select id from savage_world_chars`)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		char := models.CharShit{}
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

	char := swmodels.SWChar{}

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
		err := rows.Scan(&char.ID, &char.Name, &arr, &char.Rank)
		if err != nil {
			log.Println(err)
			continue
		}
		for i := range arr {
			char.Skills = append(char.Skills, i)
		}
	}

	data, err := json.Marshal(char)
	if err != nil {
		log.Println(err)
	}
	return data
}
