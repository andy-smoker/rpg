package savage

// SWChar - struct charshit
type SWChar struct {
	ID       int
	UserName string `json:"username"`
	Name     string `json:"name"`
	Concept  string `json:"concept"`
	Look     string `json:"look"`

	Exp        int       `json:"exp"`
	Rank       string    `json:"rank"`
	Points     int       `json:"points"`
	Race       string    `json:"race"`
	Stats      []Stat    `json:"stats"`
	Skills     []Skill   `json:"skills"`
	Traits     []Trait   `json:"trait"`
	Flaws      []Flaw    `json:"flaws"`
	Abilities  []Ability `json:"abilities"`
	PowerPoint int       `json:"power_points"`
	Inventory  []Item    `json:"inventory"`
	About      string    `json:"about"`
}

// Ability -
type Ability struct {
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

// Trait .
type Trait struct {
	ID        int64
	Name      string      `json:"name"`
	Rank      string      `json:"rank"`
	Influence string      `json:"influence"`
	Bonus     interface{} `json:"bonus"`
	About     string      `json:"about"`
}

// Flaw .
type Flaw struct {
	ID        int64
	Name      string      `json:"name"`
	Influence string      `json:"influence"`
	Debuff    interface{} `json:"debuff"`
	About     string      `json:"about"`
}

// Item .
type Item struct {
	ID     int64
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Weight int    `json:"weight"`
	Stats  []Stat `json:"stats"`
	Note   string `json:"note"`
}

// Race .
type Race struct {
	ID        int64
	Name      string
	Stats     []Stat
	Skills    []Skill
	Abilities []Ability
}

// Stat .
type Stat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Skill .
type Skill struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
