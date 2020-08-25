package savage

import (
	"database/sql"
	"log"
)

type stGetter struct {
	obj interface{}
	arr []interface{} // it not work
}

func (g *stGetter) getAll(query string, scanValue ...interface{}) {

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
		i := g.obj
		err := rows.Scan(scanValue)
		if err != nil {
			log.Println(err)
			continue
		}
		g.arr = append(g.arr, i)
	}
}
