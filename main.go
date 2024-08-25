package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
	msg.Text = rB.Message
	// create new entry to DB
	DB.Create(&msg)
}

type requestBody struct {
	// new struct
	Message string `json:"text"`
}

func main() {

	InitDB()
	DB.AutoMigrate(&Message{})
	// make newrouter
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessages).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
