package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/shinYeongHyeon/settlement-supporter/docs"
	"github.com/shinYeongHyeon/settlement-supporter/src"
	"log"
)

// @title          Settlement support API
// @version        0.1
// @description    당신의 정산을 쉽게 해줄 API
// @contact.name   Den
// @contact.email  den.shin.dev@gmail.com
// @host           localhost:9999
// @BasePath       /
func main() {
	app := fiber.New(fiber.Config{
		AppName: "Evening is Food API v0.0.1",
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Mount("/", mainModule.CreateModule())

	log.Fatal(app.Listen(":9999"))
}
