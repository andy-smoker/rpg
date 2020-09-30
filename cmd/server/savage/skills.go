package savage

type stSkill struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Stats []stat
}
