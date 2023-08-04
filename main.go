package main

import (
	"github.com/amaru0601/fluent/controllers"
	"github.com/amaru0601/fluent/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authController := controllers.NewAuthController()
	//TODO HASH PASSWORD
	//curl -X POST -H 'Content-Type: application/json' -d '{"username":"jaoks", "password":"sdtc"}' localhost:1323/register
	e.POST("/register", authController.SignUp)
	// curl -X POST -H 'Content-Type: application/json' -d '{"username":"jaoks", "password":"sdtc"}' localhost:1323/login
	e.POST("/login", authController.SignIn)

	chatController := controllers.NewChatController()
	protected := e.Group("/api")
	protected.Use(middleware.JWT(services.MySigningKey))
	protected.GET("/members", chatController.GetMembers)
	protected.GET("/chats", chatController.GetChats)
	//TODO: Hacer endpoint para jalar todos los mensajes

	sockets := e.Group("/ws")
	sockets.GET("/join", chatController.JoinChat)
	sockets.GET("/create-chat", chatController.CreateChat)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
