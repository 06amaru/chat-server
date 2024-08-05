package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/jaox1/chat-server/controllers"
	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// LOAD VAR IN LOCAL ENVIRONMENT
	_ = godotenv.Load(".env")
	security.MySigningKey = []byte(os.Getenv("SIGNING_KEY"))
}

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", func(ctx echo.Context) error { return ctx.JSON(200, models.HealthCheck{Status: "UP"}) })

	authController := controllers.NewAuthController()

	// curl -X POST -H 'Content-Type: application/json' -d '{"username":"jaoks", "password":"sdtc"}' localhost:8081/login
	e.POST("/login", authController.SignIn)

	chatController := controllers.NewChatController()
	protected := e.Group("/api")

	protected.Use(security.CustomMiddleware)

	// curl localhost:8081/api/chats --cookie "token=<YOUR_TOKEN>"
	protected.GET("/chats", chatController.GetChats)

	// curl "localhost:8081/api/:chatID/messages?limit=5&offset=0" --cookie "token=<YOUR_TOKEN>"
	protected.GET("/:chatID/messages", chatController.GetMessages)

	sockets := e.Group("/socket")
	sockets.Use(security.CustomMiddleware)
	// websocat "ws://localhost:8081/socket/join?id=<CHAT_ID>" -H "Cookie: token=<YOUR_TOKEN>"
	sockets.GET("/join", chatController.JoinChat)

	// websocat "ws://localhost:8081/socket/create-chat?to=<USERNAME>" -H "Cookie: token=<YOUR_TOKEN>"
	sockets.GET("/create-chat", chatController.CreateChat)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
