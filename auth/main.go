package main

import (
	"auth-service/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	seed()

	app := fiber.New()
	app.Use(cors.New())

	routes.Auth(app.Group("/api/v1"))

	log.Fatal(app.Listen(":3000"))
}
