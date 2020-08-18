package config

import (
	"flag"
	"fmt"
)

var (
	dbUser = flag.String("user", "user", "Enterg you postgres username")
	dbPass = flag.String("password", "pass", "")
	dbName = flag.String("database", "", "databeses name")
)

type DBconfig struct {
	DriverName string
	Config     string
}

func DBConnect() (string, string) {
	driverName := "postgres"
	config := fmt.Sprint("user=$1 password=$2 dbname=$3 sslmode=disable", dbUser, dbPass, dbName)
	return driverName, config
}
