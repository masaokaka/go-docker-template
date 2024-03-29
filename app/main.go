package main

import (
	"sample/db"
	"sample/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// DB接続
	db.Init()

	server := echo.New()
	routes.RegisterRoutes(server)

	server.Logger.Fatal(server.Start(":5050"))
}
