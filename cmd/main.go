package main

import (
	"log"

	"github.com/MetaDandy/autoparts-api/config"
	"github.com/MetaDandy/autoparts-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Aloha")
	})

	log.Println("Server started on port" + config.Port)
	app.Listen(":" + config.Port)
}
