package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rpg/server/auth"
	"rpg/server/config"
)

var sessions = []auth.Session{
	auth.Session{User: "Igor", Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTc3MTkwODIsInVzZXJfaWQiOjF9.YIzbGH6IJp6u8BI3hdT4U1PJeMcQW--FCdvkQcy_TX4"},
}

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

	db, err := sql.Open(dbConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// ....
}

func dbConfig() (string, string) {

	dbConf := config.DBconfig{
		User:     "user",
		Password: "password",
		DBname:   "database",
		SSLmode:  "disable",
	}
	return dbConf.Connect()
}
