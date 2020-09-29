package savage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/database"
	"strconv"

	"github.com/gorilla/mux"
)

type item struct {
	ID     int64
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Weight int    `json:"weight"`
	Stats  []stat `json:"stats"`
	Note   string `json:"note"`
}

// ItemMidleware ...
func ItemMidleware(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	i := item{}
	itemsArr := []item{}
	id := vars["id"]
	if id == "all" && r.Method == "GET" {
		rows, err := i.AllItems()

		if err != nil {
			log.Println(err)
			return
		}
		for _, row := range rows {
			i, ok := row.(item)
			if !ok {
				continue
			}
			itemsArr = append(itemsArr, i)
		}
		data, err := json.Marshal(itemsArr)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(data)
		return
	}
	if _, err := strconv.Atoi(id); err == nil && r.Method == "GET" {
		row, err := i.GetItem(vars["id"])
		if err != nil {
			log.Println(err)
			return
		}
		data, ok := row.(*item)
		log.Println(data)
		if !ok {
			return
		}
		resp, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(resp)
	}

	if id == "add" && r.Method == "POST" {
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Println(err)
			return
		}
		err = i.AddItem()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusContinue)
		return
	}

}

func (i *item) DropItemFromInventory() {

}

func (*item) GetItem(id string) (interface{}, error) {
	i := item{}
	var arr []interface{}
	arr = append(arr, &i.Name, &i.Type)
	return database.GetOnce(func() (interface{}, []interface{}) {
		return &i, arr
	}, "select item_name, item_type from items where item_id = $1", id)
}

func (i *item) AddItem() error {

	_, err := database.ExecOnce(
		`insert into items(item_name,item_type, item_price, item_weight, item_stats, item_about)
	values($1,2$,3$,4$,5$,6$)`, i.Name, i.Type, i.Price, i.Weight, i.Stats, i.Note)
	if err != nil {
		return err
	}
	return nil
}

func (i *item) DeleteItem() {
	fmt.Print(i)
}

// AllItems .
func (*item) AllItems() ([]interface{}, error) {

	return database.GetAll(func() (interface{}, []interface{}) {
		i := item{}
		var arr []interface{}
		arr = append(arr, &i.Name, &i.Type)

		return i, arr
	}, "select item_name, item_type from items")
}

// AillItemInventory .
func AillItemInventory() {

}
