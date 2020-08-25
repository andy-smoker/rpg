package savage

// SWChar - struct charshit
type swChar struct {
	ID       int
	UserName string `json:"username"`
	CharName string `json:"name"`

	Concept string `json:"concept"`
	Race    string `json:"race"`

	Exp    int    `json:"exp"`
	Rank   string `json:"rank"`
	Points int    `json:"points"`

	Stats      []stat    `json:"stats"`
	Skills     []skill   `json:"skills"`
	Traits     []trait   `json:"trait"`
	Flaws      []flaw    `json:"flaws"`
	Abilities  []ability `json:"abilities"`
	PowerPoint int       `json:"power_points"`
	Inventory  []item    `json:"inventory"`
	Look       string    `json:"look"`
	About      string    `json:"about"`
}

// Ability -
type ability struct {
	ID       int64
	Name     string   `json:"name"`
	Rank     string   `json:"rank"`
	Cost     int      `json:"cost"`
	Range    string   `json:"range"`
	Damage   []string `json:"damage"`
	Duration string   `json:"duration"`
	Aspect   string   `json:"aspect"`
	About    string   `json:"about"`
}

type trait struct {
	ID        int64
	Name      string      `json:"name"`
	Rank      string      `json:"rank"`
	Influence string      `json:"influence"`
	Bonus     interface{} `json:"bonus"`
	About     string      `json:"about"`
}

type stTMP struct {
	V []interface{} `json:"b"`
}

type flaw struct {
	ID        int64
	Name      string      `json:"name"`
	Influence string      `json:"influence"`
	Debuff    interface{} `json:"debuff"`
	About     string      `json:"about"`
}

type item struct {
	ID     int64
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Weight int    `json:"weight"`
	Stats  []stat `json:"stats"`
	Note   string `json:"note"`
}

type stRace struct {
	ID        int64     `json:"race_id"`
	Name      string    `json:"race_name"`
	RaceBonus raceBonus `json:"race_bonus"`
}

// func for makeDest interface
func (*stRace) args() (r interface{}, arr []interface{}) {
	race := stRace{}
	arr = append(arr, &race.ID, &race.Name)
	r = &race
	return
}

func (*stRace) makeArr() (arr []interface{}) {
	r := stRace{}
	arr = append(arr, r)
	return
}

func (r *stRace) ArrayOfStruct() []stRace {
	return []stRace{}
}

type raceBonus struct {
	Stats     []stat    `json:"stats"`
	Skills    []skill   `json:"skills"`
	Abilities []ability `json:"abiliteis"`
}

type stat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type skill struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
