package main

import (
	"github.com/gofiber/fiber/v2"
	"yuyuid.id/config"
	v1 "yuyuid.id/internal/routes/v1"
	"yuyuid.id/pkg/database"
)

func main() {
	cnf := config.Get()
	app := fiber.New()

	db := database.ConnectGORM(cnf.Database)
	//api
	api := app.Group("/api")
	v1.ApiV1(api, db)
	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
