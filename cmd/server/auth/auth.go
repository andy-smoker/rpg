package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/database"
	"time"
)

//
// User is user of programm
type User struct {
	Login             string `json:"login"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncriptedPassword string `json:"encpass"`
}

// getUser use filtre values for spesiol return
func getUser(u *User) (*User, error) {
	row, err := database.GetOnce(func() (interface{}, []interface{}) {
		var arr []interface{}
		u := User{}
		arr = append(arr, u.Login, u.Password, u.Email, u.Username)
		return &u, arr
	}, "select * from rpg_users where login=$1 password=$2", u.Login, u.Password)
	if err != nil {
		return nil, err
	}
	return row.(*User), nil
}

// Session .
type Session struct {
	User    string `json:"user"`
	Token   string `json:"token"`
	Key     int
	Timeout time.Time
}

func newSession(u User, token string) error {
	s := Session{
		User:    u.Login,
		Token:   token,
		Timeout: time.Now().Add(2 * time.Minute),
	}
	_, err := database.ExecOnce("insert into sessions(u_login, u_token, timeout) values($1,$2,$3)", s.User, s.Timeout)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) valid() error {
	row, err := database.GetOnce(func() (interface{}, []interface{}) {
		var arr []interface{}
		ses := Session{}
		arr = append(arr, ses.User, ses.Token, ses.Timeout)
		return &ses, arr
	}, "select * from sessions where u_login = $1", s.User)
	if err == sql.ErrNoRows {
		return fmt.Errorf("not exist")
	} else if err != nil {
		return err
	}
	ses, ok := row.(*Session)
	if !ok {
		return fmt.Errorf("Format error")
	}
	if ses.Token != s.Token {
		return fmt.Errorf("Token is invalid")
	}
	if t := time.Now().Sub(ses.Timeout); t < 0 {
		return fmt.Errorf("Token is axpired")
	}

	return nil
}

func (*User) Args() func() (interface{}, []interface{}) {
	var arr []interface{}
	u := User{}

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
	u := User{}
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
