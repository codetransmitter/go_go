package main

import (
	"go_go/internal/database"
	"go_go/internal/handlers"
	"go_go/internal/messagesService"
	"go_go/internal/web/messages"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.InitDB()
	err := database.DB.AutoMigrate(&messagesService.Message{})
	if err != nil {
		log.Fatalf("Failed to AutoMigrate with err: %v", err)
	}
	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// make newrouter
	strictHandler := messages.NewStrictHandler(handler, nil)
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
