package savage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"
)

type DbConn struct {
	dbsourceName string
}

func SWstart() {

}

func SWConnDB(d database.DB) {
	//dbsourceName = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Pass, d.DB)
	fmt.Printf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Pass, d.DB)
}

/*/ SWgetChar method for marshal and get data
func SWgetChar(id string) []byte {
	fmt.Println(dbsourceName)
	char := SWChar{}
	db, err := sql.Open("postgresql", dbsourceName)
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

		err := rows.Scan(&char.ID, &char.Name, &char.Rank)
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
}*/

// AddCharshit .
func AddCharshit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

	res, err := db.Exec("insert into charshit(name)values($1)", v.Name)
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
