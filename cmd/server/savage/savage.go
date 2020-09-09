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
	char := swChar{}

	rows := database.GetAll(&char, "select id, name, rank from chars")

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
		status int
		i      interface{}
	)

	switch r.Method {
	case "GET":
		status, i = getCharShit(vars["id"])
		break
	case "PUT":
		status, i = updateCharShit(vars["id"])
		break
	case "DELETE":
		status, i = deleteCharShit(vars["id"])
		break
	}
	resp, err := json.Marshal(i)
	if err != nil {
		log.Println(err)
	}

	w.Header().Add("Status-code", fmt.Sprint(status))
	w.Write(resp)
}

func getCharShit(id string) (status int, row interface{}) {
	sw := swChar{}
	row = database.GetOnce(&sw, "select id, name, rank from chars where id = $1", id)

	if sw.ID == 0 {
		status = http.StatusNotFound
		row = nil
		return

	}
	status = http.StatusOK
	return
}

func updateCharShit(id string) (status int, row interface{}) {
	sw := swChar{}
	_, err := database.ExecOnce(&sw, "update swcharshit set name=$1 where id = $1", id)
	if err != nil {
		log.Println(err)
		status = http.StatusBadRequest
	}
	status = http.StatusOK
	return
}

func deleteCharShit(id string) (status int, row interface{}) {
	sw := swChar{}
	_, err := database.ExecOnce(&sw, "delete from swcharshit where id = $1", id)
	if err != nil {
		log.Println(err)
		status = http.StatusBadRequest
	}
	status = http.StatusOK
	return
}

// AddChar .
func AddChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
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
	race := stRace{}
	arr := database.GetAll(&race, "select race_id,race_name from sw_racelist")

	data, err := json.Marshal(arr)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

/*
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
*/
