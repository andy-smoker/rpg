package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type DBConfig struct {
	User string `toml:"user"`
	Pass string `toml:"pass"`
	DB   string `toml:"db"`
}

func DBConnect() (string, string) {
	config := DBConfig{}
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	fmt.Println(config)
	driverName := "postgres"
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Pass, config.DB)
	return driverName, conn
}
