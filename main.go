package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string //global var

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// hello handler
	fmt.Fprintf(w, "Hello, %s\n", message)

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// post handler
	body := json.NewDecoder(r.Body)
	// new struct for message
	var rB requestBody
	// decode body to new struct
	err := body.Decode(&rB)
	if err != nil {
		panic(err)
	}
	// pass string(json) from body to var message
	message = rB.Message
	//fmt.Fprintf(w, "Hello %s", rB.Message)
}

type requestBody struct {
	// new struct
	Message string `json:"message"`
}

func main() {

	// make newrouter
	router := mux.NewRouter()
	// make new handlers GET an POST
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/", PostHandler).Methods("POST")
	// run listener of port
	http.ListenAndServe(":8080", router)
}
