package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinYeongHyeon/settlement-supporter/src"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Evening is Food API v0.0.1",
	})

	app.Mount("/", mainModule.CreateModule())

	log.Fatal(app.Listen(":9999"))
}
