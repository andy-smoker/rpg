package auth

import (
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

	u := user{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login or password is invalid"))
		return
	}

	_, err = database.GetOnce(u.Args(), "select login, password from users where login = $1 and password = $2", u.Login, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login or password is invalid"))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": false,
		"login": u.Login,
		"exp":   time.Now().Unix(),
	})

	t, err := token.SignedString([]byte("huita"))
	if err != nil {
		log.Println(err)
		return
	}

	isLogginedIn[t] = u.Login

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, t)))
	return
}

// ValidToken func for chek valid token
func ValidToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("")
		}
		return "huita", nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"], claims["time"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
