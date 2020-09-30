package savage

type stRace struct {
	ID        int64   `json:"race_id"`
	Name      string  `json:"race_name"`
	RaceBonus []bonus `json:"race_bonus"`
}

// func for makeDest interface
func (*stRace) Args() (r interface{}, arr []interface{}) {
	race := stRace{}
	arr = append(arr, &race.ID, &race.Name)
	r = &race
	return
}

/*
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
*/
