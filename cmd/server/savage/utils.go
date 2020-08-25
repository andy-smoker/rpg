package savage

import (
	"database/sql"
	"fmt"
	"log"
)

type stGetter struct {
	obj interface{}
	arr []interface{}
}

func makeArr(i interface{}) (arr []interface{}) {
	arr = append(arr, i)
	return
}

type newDest interface {
	args() (interface{}, []interface{})
	makeArr() []interface{}
}

func newGetter(obj interface{}) stGetter {
	array := makeArr(obj)
	return stGetter{
		obj: obj,
		arr: array,
	}
}

func getAll(d newDest, query string) []interface{} {
	arr := d.makeArr()
	db, err := sql.Open("postgres", dataConn())
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		obj, dest := d.args()
		fmt.Println(dest)
		fmt.Println(obj)
		err := rows.Scan(dest...)
		if err != nil {
			log.Println(err)
			continue
		}
		arr = append(arr, obj)
	}
	return arr
}
