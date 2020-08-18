package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"rpg/server/config"

	"rpg/server/models"
)

// GetAllSWChars send all charshit response
func GetAllSWChars(w http.ResponseWriter, r *http.Request) {
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
