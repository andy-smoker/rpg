package savage_test

import (
	"server/database"
	. "server/sevage"
	"testing"
)

var db = database.DB{
	Drive: "postgres",
	User:  "rest",
	Pass:  "rest",
	DB:    "rpg",
}

func Test_SWgetChar(t *testing.T) {

	SWConnDB(db)
	//SWgetChar("1")
}
