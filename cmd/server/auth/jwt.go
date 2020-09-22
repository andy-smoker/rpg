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

type handler struct{}

var isLogginedIn = map[string]string{}

// Middleware ...
func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := handler{}
		xToken := r.Header.Get("token")
		fmt.Println(xToken)
		if xToken == "" {
			h.Login(w, r)
			return
		} else {
			user, found := isLogginedIn[xToken]
			if found {

				log.Printf("\n Authenticated user %s\n", user)
				// Pass down the request to the next middleware (or final handler)
				next.ServeHTTP(w, r)
			} else {

				// Write an error and stop the handler chain
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
		}
	})
}

// Login ...
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	u := user{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("login or password is invalid"))
		return
	}

	_, err = database.GetOnce(&user{}, "select login, password from users where login = $1 and password = $2", u.Login, u.Password)
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

	t, err := token.SignedString([]byte(u.Login))
	if err != nil {
		log.Println(err)
		return
	}

	isLogginedIn[t] = u.Login

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`"token": %s`, t)))
	return
}
