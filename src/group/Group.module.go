package groupModule

import (
	"github.com/gofiber/fiber/v2"
	coreError "github.com/shinYeongHyeon/settlement-supporter/src/core/error"
	groupControllerCommand "github.com/shinYeongHyeon/settlement-supporter/src/group/controller/command"
)

// CreateModule : returned fiber.App for mounting groupModule
func CreateModule() *fiber.App {
	groupModule := fiber.New()

	groupModule.Post("/", groupControllerCommand.Create)
	groupModule.Get("/*", coreError.NotFoundError)

	return groupModule
}
