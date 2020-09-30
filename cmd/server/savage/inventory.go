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

type stItem struct {
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
	i := stItem{}
	itemsArr := []stItem{}
	id := vars["id"]
	if id == "all" && r.Method == "GET" {
		rows, err := i.AllItems()

		if err != nil {
			log.Println(err)
			return
		}
		for _, row := range rows {
			i, ok := row.(stItem)
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
		data, ok := row.(*stItem)
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

func (*stItem) GetItem(id string) (interface{}, error) {
	i := stItem{}
	var arr []interface{}
	arr = append(arr, &i.Name, &i.Type)
	return database.GetOnce(func() (interface{}, []interface{}) {
		return &i, arr
	}, "select item_name, item_type from items where item_id = $1", id)
}

func (i *stItem) AddItem() error {

	_, err := database.ExecOnce(
		`insert into items(item_name,item_type, item_price, item_weight, item_stats, item_about)
	values($1,2$,3$,4$,5$,6$)`, i.Name, i.Type, i.Price, i.Weight, i.Stats, i.Note)
	if err != nil {
		return err
	}
	return nil
}

func (i *stItem) DeleteItem() {
	fmt.Print(i)
}

// AllItems .
func (*stItem) AllItems() ([]interface{}, error) {

	return database.GetAll(func() (interface{}, []interface{}) {
		i := stItem{}
		var arr []interface{}
		arr = append(arr, &i.Name, &i.Type)

		return i, arr
	}, "select item_name, item_type from items")
}

// AllItemInventory .
func AllItemInventory(charID int) ([]interface{}, error) {
	inventory, err := database.GetAll(func() (interface{}, []interface{}) {
		var (
			i   int64
			arr []interface{}
		)
		return i, arr
	}, "select inventory from chars where id = $1", charID)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (i *stItem) AddItemInInventory(charID int) error {

	inventory, err := AllItemInventory(charID)
	inventory = append(inventory, i.ID)
	_, err = database.ExecOnce("update chars set inventory = $1  where id = $2", inventory, i.ID)
	if err != nil {
		return err
	}
	return nil

}

func (i *stItem) DropItemFromInventory(itemID int, charID int) error {
	inventory, err := AllItemInventory(charID)
	for index, id := range inventory {
		if id == itemID {
			inventory = append(inventory[:index], inventory[index+1:]...)
		}
	}
	_, err = database.ExecOnce("update chars set inventory = $1  where id = $2", inventory, i.ID)
	if err != nil {
		return err
	}
	return nil
}
