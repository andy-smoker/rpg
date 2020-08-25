package auth

/*
var tempDB = []User{
	User{ID: 1, Username: "Igor", Password: "111"},
	User{ID: 2, Username: "Jija", Password: "555"},
}

// User is user of programm
type User struct {
	ID                uint64 `json:"id"`
	Login             string `json:"login"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string
	EncriptedPassword string `json:"encpass"`
}

// Session .
type Session struct {
	Address string
	Key     int
}

func NewSession(addr string) Session {
	rand.Seed(time.Now().UnixNano())
	key := rand.Intn(100)

	return Session{
		Address: addr,
		Key:     key,
	}
}

// AuthHandler .
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

	token, err := u.authMethod()
	u.refreshToken(token)
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Token", token)
	w.WriteHeader(http.StatusOK)

}


func (u *User) authMethod() (string, error) {

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
		return creatToken(user.ID)
	}

	return "", fmt.Errorf("login or password is fail")

}

// creatToken code from
//https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func creatToken(userid uint64) (string, error) {
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

func (u *User) refreshToken(token string) {
	db, err := sql.Open(config.DBConnect())
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
*/
