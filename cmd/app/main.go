package main

import (
	"net/http"

	"go_go/internal/database"
	"go_go/internal/handlers"
	"go_go/internal/messagesService"

	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)
	// make newrouter
	router := mux.NewRouter()
	router.HandleFunc("/api/post", handler.CreateMessage).Methods("POST")
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/patch", handler.UpdateMessageByID).Methods("PATCH")
	router.HandleFunc("/api/del", handler.DeleteMessageByID).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
