package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var isLogginedIn = map[string]string{}

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/reg" || r.URL.Path == "/auth" {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println("token: " + r.Header.Get("Authorization"))
		xToken := r.Header.Get("Access-Control-Request-Headers")
		//valid := ValidToken(xToken)

		user, found := isLogginedIn[xToken]
		if found {

			log.Printf("\n Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {

			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

	})
}

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	defer r.Body.Close()

	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login or password is invalid"))
		return
	}
	_, err = getUser(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login or password is invalid"))
		return
	}

	// создаём токен
	t, key, err := CreatToken(u)
	if err != nil {
		log.Println(err)
		return
	}
	err = NewSession(u, t.(string), key)
	if err != nil {
		log.Println(err)
		return
	}
	isLogginedIn[t.(string)] = u.Login

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, t)))
	return
}

// CreatToken .
func CreatToken(u User) (interface{}, interface{}, error) {
	// создаём токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": false,
		"login": u.Login,
		"exp":   time.Now().Unix(),
	})
	// создаём ключ
	var key interface{}
	key, err := base64.StdEncoding.DecodeString(u.Login)
	if err != nil {
		return nil, nil, err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	return tokenString, key, nil
}

// ValidToken func for chek valid token
func ValidToken(tokenString string, u User) bool {
	// достаём из БД данные по сесии
	row, err := database.GetOnce(func() (interface{}, []interface{}) {
		var arr []interface{}
		s := Session{}
		arr = append(arr, s.User, s.Token, s.Timeout)
		return &s, arr
	}, "select * from sessions whith s_user = $1", u.Login)
	if err != nil {
		return false
	}
	session := row.(*Session) // парсим строку из БД в структуру
	// сравниваем время таймаута с настоящим времинем
	if time.Now().Unix() < session.Timeout {
		return false
	}
	// создаём токен для верификации
	vToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": false,
		"login": u.Login,
		"exp":   session.Timeout,
	})
	vString, err := vToken.SignedString(session.Key)
	if err != nil {
		log.Println(err)
		return false
	}

	return vString == tokenString
}
