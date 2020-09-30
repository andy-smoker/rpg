package savage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"

	"github.com/gorilla/mux"
)

// SWChar - struct charshit
type swChar struct {
	ID       int
	Owner    string `json:"owner"`
	CharName string `json:"name"`

	Concept string `json:"concept"`
	Race    string `json:"race"`

	Exp    int    `json:"exp"`
	Rank   string `json:"rank"`
	Points int    `json:"points"`

	Stats      []stat   `json:"stats"`
	Skills     []string `json:"skills"`
	Traits     []string `json:"trait"`
	Flaws      []string `json:"flaws"`
	Abilities  []string `json:"abilities"`
	PowerPoint int      `json:"power_points"`
	Inventory  []stItem `json:"inventory"`
	Look       string   `json:"look"`
	About      string   `json:"about"`
}

func (*swChar) Args(q interface{}) func() (interface{}, []interface{}) {

	if q.(string) == "*" {

		return func() (interface{}, []interface{}) {
			st := swChar{}
			var arr []interface{}
			arr = append(arr, &st.ID, &st.CharName, &st.Rank, &st.Inventory)
			return &st, arr
		}

	}

	return nil
}

// GetAllChars -get all charshit from database
func GetAllChars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println(r.Header.Get("Authorization"))
	log.Println(r.Body)
	sw := swChar{}
	rows, err := database.GetAll(sw.Args("*"), "select id, name, rank from chars")
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(string(resp))
		w.Write(resp)
		break
	case "PUT":
		err = updateCharShit(vars["id"])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		break
	case "DELETE":
		err = deleteCharShit(vars["id"])
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		break
	}

	w.WriteHeader(http.StatusOK)
	return

}

func getCharShit(id string) (interface{}, error) {
	sw := swChar{}
	row, err := database.GetOnce(sw.Args("*"), "select id, name, rank from chars where id = $1", id)
	if err != nil {
		return nil, err
	}
	return row, err
}

func updateCharShit(id string) error {
	_, err := database.ExecOnce("update chars set name=$1 where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func deleteCharShit(id string) error {
	_, err := database.ExecOnce("delete from chars where id = $1", id)
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
	res, err := database.ExecOnce(`insert into chars(charname,owner, concepid, raceid, epx, rank, points)
	values($1,2$,3$,4$,5$,6$)`, sw.CharName, sw.Owner, sw.Concept, sw.Race, sw.Exp, sw.Rank, sw.Points)
	if err != nil {
		log.Println(err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Успешно добвален: ", lastID)
}
