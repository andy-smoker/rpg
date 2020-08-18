package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	User  string
	Token string
	Time  string
}

func NewUser(name string, pass string) User {
	return User{
		Username: name,
		Password: pass,
	}
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
	for _, user := range tempDB {
		if user.Username == u.Username && user.Password == u.Password {
			log.Println("login")
			return CreatToken(user.ID)
		}
	}
	return "", fmt.Errorf("login or password is fail")

}

func CreatToken(userid uint64) (string, error) {
	os.Setenv("ACCESS_SECRET", "wefghjmsdfg") // this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
