package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var DB map[uint]Human
var id uint

type Human struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

// type HumanForJson struct {
// 	Fname string `json:"fname"`
// 	Lname string `json:"lname"`
// 	Age   int    `json:"age"`
// 	Id    uint   `json:"id"`
// }

func main() {
	DB = make(map[uint]Human, 0)
	http.HandleFunc("/read", readHuman)
	http.HandleFunc("/update", updateHuman)
	http.HandleFunc("/delete", deleteHuman)
	http.HandleFunc("/create", createNewHuman)
	http.ListenAndServe(":80", nil)
}

func createNewHuman(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("err")
	}
	defer r.Body.Close()
	human := Human{}
	err = json.Unmarshal(b, &human)
	if err != nil {
		panic("err")
	}
	id++
	DB[id] = human
	fmt.Fprint(w, "Human was created with ID: ", id)
}
func updateHuman(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("err")
	}
	updatingID, err := strconv.Atoi(r.URL.Query().Get("Id"))
	if err != nil {
		panic("err")
	}
	defer r.Body.Close()
	human := Human{}
	err = json.Unmarshal(b, &human)
	if err != nil {
		panic("err")
	}
	if updatingID > 0 {
		_, ok := DB[uint(updatingID)]
		if ok {
			DB[uint(updatingID)] = human
			fmt.Fprint(w, "Human was upated")
		} else {
			fmt.Fprint(w, "There is no such human to be updated")
		}
	}
}
func deleteHuman(w http.ResponseWriter, r *http.Request) {
	idOfDeleting, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		panic("aaa")
	}
	if idOfDeleting > 0 {
		if _, ok := DB[uint(idOfDeleting)]; ok {
			delete(DB, uint(idOfDeleting))
			fmt.Fprint(w, "Human was deleted")
		} else {
			fmt.Fprint(w, "There is no such human to be deleted")
		}
	}
}
func readHuman(w http.ResponseWriter, r *http.Request) {
	idOfReading, err := strconv.Atoi(r.URL.Query().Get("ID"))
	if err != nil {
		panic("aaa")
	}
	if idOfReading > 0 {
		if _, ok := DB[uint(idOfReading)]; ok {
			fmt.Fprint(w, DB[uint(idOfReading)])
		} else {
			fmt.Fprint(w, "There is no such human")
		}
	}
}
