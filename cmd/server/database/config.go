package database

import (
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

//
func (config *DB) ConfigToml() error {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return err
	}
	return nil

}
