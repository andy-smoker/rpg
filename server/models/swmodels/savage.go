package swmodels

// Char .
type SWChar struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Concept    string    `json:"concept"`
	Look       string    `json:"look"`
	Exp        int       `json:"exp"`
	Rank       string    `json:"rank"`
	Points     int       `json:"points"`
	Race       Race      `json:"race"`
	Stats      Stats     `json:"stats"`
	Skills     []int     `json:"skills"`
	Traits     []Trait   `json:"trait"`
	Flaws      []Flaw    `json:"flaws"`
	Abilities  []Ability `json:"abilities"`
	PowerPoint int       `json:"power_points"`
	Inventory  []Item    `json:"inventory"`
	About      string    `json:"about"`
}

// Ability .
type Ability struct {
	Name     string      `json:"name"`
	Rank     string      `json:"rank"`
	Cost     int         `json:"cost"`
	Range    string      `json:"range"`
	Damage   []string    `json:"damage"`
	Duration string      `json:"duration"`
	Aspect   interface{} `json:"aspect"`
	About    string      `json:"about"`
}

// Trait .
type Trait struct {
	Name      string      `json:"name"`
	Rank      string      `json:"rank"`
	Influence string      `json:"influence"`
	Bonus     interface{} `json:"bonus"`
	About     string      `json:"about"`
}

// Flaw .
type Flaw struct {
	Name      string      `json:"name"`
	Influence string      `json:"influence"`
	Debuff    interface{} `json:"debuff"`
	About     string      `json:"about"`
}

// Item .
type Item struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Price  int         `json:"price"`
	Weight int         `json:"weight"`
	Stats  interface{} `json:"stats"`
	Note   string      `json:"note"`
}

// Race .
type Race struct {
	Name string `json:"name"`
}

// Stats .
type Stats struct {
	Step       int `json:"step"`
	Defanse    int `json:"defanse"`
	Durability int `json:"durability"`
	Charm      int `json:"charm"`
	Dex        int `json:"dex"`
	Savviy     int `json:"savvy"`
	Character  int `json:"character"`
	Force      int `json:"force"`
	Stamina    int `json:"satamina"`
}
