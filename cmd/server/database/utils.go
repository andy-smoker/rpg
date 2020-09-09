package database

import (
	"database/sql"
	"fmt"
	"log"

	// posgres driver for sgl
	_ "github.com/lib/pq"
)

func dataPosgresConn() (driver string, datasourceName string) {
	d := NewDB()
	err := d.ConfigToml()
	if err != nil {
		log.Println(err)
	}
	driver = "postgres"
	datasourceName = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Pass, d.DB)
	return
}

// create dest for sql scan
type newDest interface {
	Args() (interface{}, []interface{})
}

// GetAll return all rows from posgres table
func GetAll(d newDest, query string, args ...interface{}) []interface{} {
	var arr []interface{}
	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj, dest := d.Args()
		err := rows.Scan(dest...)
		if err != nil {
			log.Println(err)
			continue
		}
		arr = append(arr, obj)
	}
	return arr
}

// GetOnce return one row from posgres table
func GetOnce(d newDest, query string, args ...interface{}) interface{} {
	obj, dest := d.Args()
	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(dest...)

		if err != nil {
			log.Println(err)
			continue
		}
		return obj
	}
	return nil
}

// ExecOnce fanc for add/delet/upgrade row form posgres table
func ExecOnce(obj interface{}, query string, args ...interface{}) (sql.Result, error) {
	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
