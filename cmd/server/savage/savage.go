package savage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"

	"github.com/gorilla/mux"
)

// GetAllChars -get all charshit from database
func GetAllChars(w http.ResponseWriter, r *http.Request) {
	char := swChar{}

	rows := database.GetAll(&char, "select * from chars")

	data, err := json.Marshal(rows)
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
	sw := swChar{}
<<<<<<< HEAD
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select charname,username, concepid, raceid, epx, rank, points from chars where id = $1", id)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(sw.CharName, sw.UserName, sw.Concept, sw.Race, sw.Exp, sw.Rank, sw.Points)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	data, err := json.Marshal(sw)
=======
	row := database.GetOnce(&sw, "select id, name, rank from chars where id = $1", id)

	data, err := json.Marshal(row)
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
	if err != nil {
		log.Println(err)
	}
	return data
}

func updateCharShit(id string) []byte {
<<<<<<< HEAD
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
=======
	sw := swChar{}
	resp := `{"401"}`
	_, err := database.ExecOnce(&sw, "update swcharshit set name=$1 where id = $1", id)
	if err != nil {
		log.Println(err)
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
	return data
}

func deleteCharShit(id string) []byte {
<<<<<<< HEAD
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
=======
	sw := swChar{}
	resp := `{"401"}`
	_, err := database.ExecOnce(&sw, "delete from swcharshit where id = $1", id)
	if err != nil {
		log.Println(err)
	}
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
	resp = `{"complete":true}`
	data, err := json.Marshal(resp)
	return data
}

// AddChar .
func AddChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
<<<<<<< HEAD
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	sw := swChar{}
	err = json.NewDecoder(r.Body).Decode(&sw)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := db.Exec(`insert into swcharshit(charname,username, concepid, raceid, epx, rank, points)
	values($1,2$,3$,4$,5$,6$)`, sw.CharName, sw.UserName, sw.Concept, sw.Race, sw.Exp, sw.Rank, sw.Points)
	if err != nil {
		log.Println(err)
		return
	}
=======
	sw := swChar{}
	res, err := database.ExecOnce(&sw, `insert into swcharshit(charname,username, concepid, raceid, epx, rank, points)
	values($1,2$,3$,4$,5$,6$)`, sw.CharName, sw.UserName, sw.Concept, sw.Race, sw.Exp, sw.Rank, sw.Points)

>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintln(w, "Успешно добвален: ", lastID)
}

// GetAllRaces - get all race name from database
func GetAllRaces(w http.ResponseWriter, r *http.Request) {
	race := stRace{ID: 0}
<<<<<<< HEAD
	arr := getAll(&race, "select race_id,race_name from sw_racelist")
=======
	arr := database.GetAll(&race, "select race_id,race_name from sw_racelist")
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3

	data, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

<<<<<<< HEAD
=======
/*
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
// GetAbility - get ability from database
func GetRace(w http.ResponseWriter, r *http.Request) {

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
<<<<<<< HEAD
=======
*/
>>>>>>> 67aa5527866ea73618c020c7a1b05606363d4ca3
