package savage

type flaw struct {
	ID        int64
	Name      string      `json:"name"`
	Influence string      `json:"influence"`
	Debuff    interface{} `json:"debuff"`
	About     string      `json:"about"`
}

type bonus struct {
	Stats     []stat
	Skills    []int // pool of skillsID for get skills from DB
	Abilities []int // pool of abilitiesID for get abilities from DB
	Traits    []int // pool of traitsID for get trsits from DB
}

type stat struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func newStat(n string, v int) stat {
	return stat{
		Name:  n,
		Value: v,
	}
}

func newStats(pool map[string]int) (stats []stat) {
	for n, v := range pool {
		stats = append(stats, newStat(n, v))
	}
	return
}
