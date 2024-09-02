package main

import (
	"go_go/internal/database"
	"go_go/internal/handlers"
	"go_go/internal/messagesService"
	"go_go/internal/userService"
	"go_go/internal/web/messages"
	"go_go/internal/web/users"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// DATABASE
	database.InitDB()
	err := database.DB.AutoMigrate(&messagesService.Message{})
	if err != nil {
		log.Fatalf("Failed to AutoMigrate with err: %v", err)
	}

	//MESSAGES
	messageRepo := messagesService.NewMessageRepository(database.DB)
	messageService := messagesService.NewService(messageRepo)
	messageHandler := handlers.NewHandler(messageService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictMessageHandler := messages.NewStrictHandler(messageHandler, nil)
	messages.RegisterHandlers(e, strictMessageHandler)

	//USER
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	//START SERVER
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
