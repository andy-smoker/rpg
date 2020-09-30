package savage

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
