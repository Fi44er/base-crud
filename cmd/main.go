package main

import (
	"log"
	"root/db"
	"root/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectDB()
	app := fiber.New(fiber.Config{})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
