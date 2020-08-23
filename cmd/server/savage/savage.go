package savage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type DB struct {
	driver       string
	dbsourceName string
}

var dbsourceName string

func DataConn(d database.DB) {

	dbsourceName = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Pass, d.DB)

}

// CharID .
func CharID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		w.Write(SWgetChar(vars["id"]))
		break
	case "PUT":
		fmt.Println("PUT method")
		break
	case "DELETE":
		fmt.Println("DELETE method")
		break
	}

}

// SWgetChar method for marshal and get data
func SWgetChar(id string) []byte {
	char := SWChar{}
	//

	db, err := sql.Open("postgres", dbsourceName)
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

// AddCharshit .
func AddCharshit(w http.ResponseWriter, r *http.Request) {
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

	res, err := db.Exec("insert into charshit(name,rank)values($1)", v.Name, v.Rank)
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
