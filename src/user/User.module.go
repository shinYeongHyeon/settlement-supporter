package userModule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shinYeongHyeon/settlement-supporter/src/core/error"
	"github.com/shinYeongHyeon/settlement-supporter/src/user/controller/command"
)

// CreateModule : returned fiber.App for mounting userModule
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Post("/", userControllerCommand.Create)

	userModule.Get("/*", coreError.NotFoundError)

	return userModule
}
