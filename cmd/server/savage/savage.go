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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := database.GetAll(&swChar{}, "select id, name, rank from chars")
	if err != nil {
		log.Println(err)
	}
	data, err := json.Marshal(rows)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

// CharID - funcs chatshit by id from database
func CharID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	var (
		err error
	)

	switch r.Method {
	case "GET":
		row, err := getCharShit(vars["id"])
		swCh, _ := row.(*swChar)
		resp, err := json.Marshal(swCh)
		if err != nil {
			log.Println(err)
		}
		w.Write(resp)
		break
	case "PUT":
		err = updateCharShit(vars["id"])
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	case "DELETE":
		err = deleteCharShit(vars["id"])
		if err != nil {
			fmt.Println(err)
		}
		break
	}

	w.WriteHeader(http.StatusOK)

}

func getCharShit(id string) (interface{}, error) {
	row, err := database.GetOnce(&swChar{}, "select id, name, rank from chars where id = $1", id)
	if err != nil {
		return nil, err
	}
	return row, err
}

func updateCharShit(id string) error {
	_, err := database.ExecOnce(&swChar{}, "update swcharshit set name=$1 where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func deleteCharShit(id string) error {
	_, err := database.ExecOnce(&swChar{}, "delete from swcharshit where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// AddChar .
func AddChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sw := swChar{}
	err := json.NewDecoder(r.Body).Decode(&sw)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := database.ExecOnce(&sw, `insert into swcharshit(charname,username, concepid, raceid, epx, rank, points)
	values($1,2$,3$,4$,5$,6$)`, sw.CharName, sw.UserName, sw.Concept, sw.Race, sw.Exp, sw.Rank, sw.Points)
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

// GetAllRaces - get all race name from database
func GetAllRaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	race := stRace{}
	arr, err := database.GetAll(&race, "select race_id,race_name from sw_racelist")
	if err != nil {
		log.Println(err)
	}
	data, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}
