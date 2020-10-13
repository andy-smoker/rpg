package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	t := CreatToken(u)
	if _, ok := t.(error); ok {
		log.Println(t)
		return
	}
	tokenString := t.(string)
	err = newSession(u, tokenString)
	if err != nil {
		log.Println(err)
		return
	}
	isLogginedIn[tokenString] = u.Login

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, t)))
	return
}

// CreatToken .
func CreatToken(u User) interface{} {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": false,
		"login": u.Login,
		"exp":   time.Now().Unix(),
	})
	var key interface{}
	key, err := base64.StdEncoding.DecodeString(u.Login)
	if err != nil {
		return err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Println(err)
		return nil
	}
	return tokenString
}

// ValidToken func for chek valid token
func ValidToken(tokenString string, key interface{}) bool {
	/*
		короче, нужно создавать клейм, в него записывать логин и дебаг уровень(типа адин или нет)
		и в него же вытягивать время токена из бд и сравнивать с поступающим токеном
		так что еби могз додела и будет тебе коммит и пуш
	*/
	return true
}
