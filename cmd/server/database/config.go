package database

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type DB struct {
	Drive string `toml:"drive"`
	User  string `toml:"user"`
	Pass  string `toml:"pass"`
	DB    string `toml:"db"`
}

func NewDB() DB {
	return DB{
		User: "user",
		Pass: "password",
	}
}

func (config *DB) ConfigToml() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)

}
