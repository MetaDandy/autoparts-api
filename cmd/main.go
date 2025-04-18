package main

import (
	"log"

	"github.com/MetaDandy/autoparts-api/cmd/api"
	"github.com/MetaDandy/autoparts-api/config"
	"github.com/MetaDandy/autoparts-api/middleware"
	"github.com/MetaDandy/autoparts-api/src"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Aloha")
	})

	c := src.SetUpContainer()
	api.SetupApi(app, c)

	log.Println("Server started on port" + config.Port)
	app.Listen(":" + config.Port)
}
