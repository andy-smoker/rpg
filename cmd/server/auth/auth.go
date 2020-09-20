package auth

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/database"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func (*user) Args() (r interface{}, arr []interface{}) {
	u := user{}
	arr = append(arr, &u.Login, &u.Password)
	r = &u
	return
}

// AuthHandler .
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	u := user{}
	defer r.Body.Close()

	if r.Body == nil {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		if err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("null"))
		}
		return
	}

	token, err := u.authMethod()
	//u.refreshToken(token)
	if err != nil {
		log.Print(err)
		w.Header().Set("Status-code", fmt.Sprint(http.StatusBadRequest))
		w.Write([]byte("login or pass not exist"))
		return
	}

	resp := Session{
		User:  user{Login: u.Login},
		Token: token,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	}
	w.WriteHeader(http.StatusFound)
	w.Write(data)

}

func (u *user) authMethod() (string, error) {
	dbreq, err := database.GetOnce(u, "select login, password from users where login = $1 and password = $2", u.Login, u.Password)
	if err != nil {
		return "", err
	}
	usr, _ := dbreq.(*user)

	log.Println(usr, dbreq)
	token, err := creatToken(usr)
	if err != nil {
		return "", err
	}
	return token, nil

}

// creatToken code from
//https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func creatToken(u *user) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"auth":    true,
		"user_id": u.ID,
		"exp":     time.Now().Unix(),
	})
	key := sha256.Sum256([]byte(u.Login))
	tokenString, err := token.SignedString(key[:])
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("")
		}
		return "huica", nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"])
	} else {
		fmt.Println(err)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	_, err = database.GetOnce(&u, "select login, username, password where login=$1", u.Login)
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
