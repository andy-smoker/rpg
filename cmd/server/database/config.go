package database

import (
	"github.com/BurntSushi/toml"
)

// DB exported type
type DB struct {
	Drive string `toml:"drive"`
	User  string `toml:"user"`
	Pass  string `toml:"pass"`
	DB    string `toml:"db"`
}

// NewDB new dbconnect config
func NewDB() DB {
	return DB{
		User: "user",
		Pass: "password",
	}
}

// ConfigToml enterin config for connect from .toml
func (config *DB) ConfigToml() error {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return err
	}
	return nil

}
