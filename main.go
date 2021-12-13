package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Player!")
	})

	app.Get("/:coord", doKnightPath)

	app.Listen(":80")

}
