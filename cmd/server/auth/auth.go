package auth

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"server/database"
	"time"
)

var tempDB = []user{
	user{ID: 1, Username: "Igor", Password: "111"},
	user{ID: 2, Username: "Jija", Password: "555"},
}

//
// User is user of programm
type user struct {
	ID                uint64 `json:"id"`
	Login             string `json:"login"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncriptedPassword string `json:"encpass"`
}

// Session .
type Session struct {
	User    user   `json:"user"`
	Token   string `json:"token"`
	Key     int
	Timeout time.Time
}

func (*user) Args() func() (interface{}, []interface{}) {
	var arr []interface{}
	u := user{}

	return func() (interface{}, []interface{}) {
		arr = append(arr, &u.Login, &u.Password)
		return &u, arr
	}
}

/*
func (u *user) refreshToken(token string) {
	db, err := sql.Open(database.DBConnect())
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec(`update users set token = $1
	where login = $2`, token, u.Login)
	if err != nil {
		log.Println(err)
		return
	}

}
*/

// Register new user
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	u := user{}
	defer r.Body.Close()
	if r.Body == nil {
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	_, err = database.GetOnce(u.Args(), "select login, username, password where login=$1", u.Login)
	if err != nil {
		log.Println(err)
		return
	} else if err == sql.ErrNoRows {
		_, err = database.ExecOnce(`insert into users (login, password, username)
	values($1, $2, $3);`, u.Login, u.Password, u.Username)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}
