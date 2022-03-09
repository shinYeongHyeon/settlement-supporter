package userModule

import (
	"github.com/gofiber/fiber/v2"
	coreError "github.com/shinYeongHyeon/settlement-supporter/src/core/error"
	userController "github.com/shinYeongHyeon/settlement-supporter/src/user/controller"
)

// CreateModule : returned fiber.App for mounting userModule
func CreateModule() *fiber.App {
	userModule := fiber.New()

	userModule.Post("/", userController.Create)

	userModule.Get("/*", coreError.NotFoundError)

	return userModule
}
