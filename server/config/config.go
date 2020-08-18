package config

import (
	"fmt"
)

type DBconfig struct {
	User     string
	Password string
	DBname   string
	SSLmode  string
}

func NewDBconfig() DBconfig {
	return DBconfig{
		User:     "user",
		Password: "password",
		DBname:   "database",
		SSLmode:  "disable",
	}
}

func (conf *DBconfig) Connect() (string, string) {
	driverName := "postgres"
	sqlConf := fmt.Sprint("user=$1 password=$2 dbname=$3 sslmode=disable", conf.User, conf.Password, conf.DBname)
	return driverName, sqlConf
}
