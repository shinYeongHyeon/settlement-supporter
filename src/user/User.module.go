package userModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// CreateModule : returned fiber.App for mounting userModule
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s, in User Module", c.Params("*"))
		return c.SendString(msg)
	})

	return userModule
}
