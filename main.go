package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateMessages(w http.ResponseWriter, r *http.Request) {
	// proc request PATCH
	body := json.NewDecoder(r.Body)
	var rB requestBody
	var msg Message
	err := body.Decode(&rB)
	if err != nil {
		panic(err)
	}

	// req to DB for update by ID
	msg.Text, msg.ID = rB.Message, rB.ID
	//fmt.Fprintf(w, "ID:_%d\nText:_%s", msg.ID, msg.Text)
	DB.Save(&msg)
}

func DeleteMessages(w http.ResponseWriter, r *http.Request) {
	// proc request PATCH
	body := json.NewDecoder(r.Body)
	var rB requestBody
	var msg Message
	err := body.Decode(&rB)
	if err != nil {
		panic(err)
	}
	// req to DB for update by ID
	msg.ID = rB.ID
	DB.Delete(&msg)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	// hello handler
	var msg []Message
	DB.Find(&msg)
	// print all slice
	for _, v := range msg {
		fmt.Fprintf(w, "%s\n", v.Text)
	}
}

func CreateMessages(w http.ResponseWriter, r *http.Request) {
	// post handler
	body := json.NewDecoder(r.Body)
	// new struct for message
	var rB requestBody
	var msg Message
	// decode body to new struct
	err := body.Decode(&rB)
	if err != nil {
		panic(err)
	}
	// pass string(json) from body to var message
	msg.ID = rB.ID
	// create new entry to DB
	DB.Create(&msg)
}

type requestBody struct {
	// new struct
	ID      uint   `json: "id"`
	Message string `json:"text"`
}

func main() {

	InitDB()
	DB.AutoMigrate(&Message{})
	// make newrouter
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessages).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", UpdateMessages).Methods("PATCH")
	router.HandleFunc("/api/messages", DeleteMessages).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
