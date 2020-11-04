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

// GetAll return all rows from posgres table
func GetAll(d func() (interface{}, []interface{}), query string, args ...interface{}) ([]interface{}, error) {
	arr := make([]interface{}, 0)

	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj, dest := d()
		err := rows.Scan(dest...)

		if err != nil {
			log.Println(err)
			continue
		}

		arr = append(arr, obj)

	}
	return arr, nil
}

// GetOnce return one row from posgres table
func GetOnce(d func() (interface{}, []interface{}), query string, args ...interface{}) (interface{}, error) {
	obj, dest := d()
	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(dest...)

		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(obj)
		return obj, nil
	}
	return nil, sql.ErrNoRows
}

// ExecOnce fanc for add/delet/upgrade row form posgres table
func ExecOnce(query string, args ...interface{}) (sql.Result, error) {
	db, err := sql.Open(dataPosgresConn())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

// CeateTable - creating a table for specific tasks
func CeateTable() {

}
