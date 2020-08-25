package savage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"

	"github.com/gorilla/mux"
	//
	_ "github.com/lib/pq"
)

func dataConn() (dbsourceName string) {
	d := database.NewDB()
	err := d.ConfigToml()
	if err != nil {
		log.Println(err)
	}
	dbsourceName = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Pass, d.DB)
	return
}

// GetAllChars - get all charshit from database
func GetAllChars(w http.ResponseWriter, r *http.Request) {
	chars := []SWChar{}
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select name,rank from chars")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		char := SWChar{}
		err := rows.Scan(&char.Name, &char.Rank)
		if err != nil {
			log.Println(err)
			continue
		}
		chars = append(chars, char)
	}
	data, err := json.Marshal(chars)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

// CharID - funcs chatshit by id from database
func CharID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		w.Write(getCharShit(vars["id"]))
		break
	case "PUT":
		w.Write(updateCharShit(vars["id"]))
		break
	case "DELETE":
		w.Write(deleteCharShit(vars["id"]))
		break
	}

}

// SWgetChar method for marshal and get data
func getCharShit(id string) []byte {
	char := SWChar{}
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select name,rank from chars where id = $1", id)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&char.Name, &char.Rank)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	data, err := json.Marshal(char)
	if err != nil {
		log.Println(err)
	}
	return data
}

func updateCharShit(id string) []byte {
	resp := `{"ok":200}`
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec("update swcharshit set name=$1 where id = $1", id)
	if err != nil {
		log.Println(err)
		resp = `{"401"}`
	}
	data, err := json.Marshal(resp)
	return data
}

func deleteCharShit(id string) []byte {
	resp := `{"401"}`
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec("delete from swcharshit where id = $1", id)
	if err != nil {
		log.Println(err)
		return nil
	}
	resp = `{"complete":true}`
	data, err := json.Marshal(resp)
	return data
}

// AddChar .
func AddChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	v := SWChar{}
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := db.Exec("insert into swcharshit(name,rank)values($1)", v.Name, v.Rank)
	if err != nil {
		log.Println(err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "Успешно добвален: ", lastID)
}

// GetAllAbilities - get all abilities from database
func GetAllAbilities(w http.ResponseWriter, r *http.Request) {

}

// GetAbility - get ability from database
func GetAbility(w http.ResponseWriter, r *http.Request) {

}

func GetAllTraits(w http.ResponseWriter, r *http.Request) {

}

func GetTrait(w http.ResponseWriter, r *http.Request) {

}

func GetAllFlaws(w http.ResponseWriter, r *http.Request) {

}

func GetFlaw(w http.ResponseWriter, r *http.Request) {

}

func GetAllItems(w http.ResponseWriter, r *http.Request) {

}

func GetItem(w http.ResponseWriter, r *http.Request) {

}
