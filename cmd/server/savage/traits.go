package savage

type trait struct {
	ID        int64
	Name      string      `json:"name"`
	Rank      string      `json:"rank"`
	Influence string      `json:"influence"`
	Bonus     interface{} `json:"bonus"`
	About     string      `json:"about"`
}
