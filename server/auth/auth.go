package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rpg/server/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var tempDB = []User{
	User{ID: 1, Username: "Igor", Password: "111"},
	User{ID: 2, Username: "Jija", Password: "555"},
}

// Auth struct for authification
type Auth struct {
	Name     string
	Password string
	EncPass  string
	Token    string
}

// User struct
type User struct {
	ID                uint64 `json:"id"`
	Login             string `json:"login"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	EncriptedPassword string `json:"encpass"`
}

type Session struct {
	User  string
	Token string
	Time  string
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	u := User{}
	defer r.Body.Close()
	if r.Body == nil {
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if err == io.EOF {
			fmt.Fprint(w, "пустой запрос")
		}
		return
	}

	token, err := u.AuthMethod()
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusOK)

}

func (u *User) AuthMethod() (string, error) {
	user := User{}
	db, err := sql.Open(config.DBConnect())
	if err != nil {
		return "", err
	}
	defer db.Close()

	row, err := db.Query("select id,login,password from users where login = $1", u.Login)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Println(user)
	fmt.Println(u)
	if user.Login == u.Login && user.Password == u.Password {
		log.Println("login")
		return CreatToken(user.ID)
	}

	return "", fmt.Errorf("login or password is fail")

}

func CreatToken(userid uint64) (string, error) {
	os.Setenv("ACCESS_SECRET", "wefghjmsdfg") // this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Second * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Register new user
func Register(w http.ResponseWriter, r *http.Request) {
	u := User{}
	defer r.Body.Close()
	if r.Body == nil {
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if err == io.EOF {
			fmt.Fprint(w, "пустой запрос")
		}
		return
	}

	db, err := sql.Open(config.DBConnect())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec(`insert into users (login, password, username)
	values($1, $2, $3);`, u.Login, u.Password, u.Username)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
