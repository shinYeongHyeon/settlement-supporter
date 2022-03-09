package userController

import (
	"github.com/gofiber/fiber/v2"
	userControllerDto "github.com/shinYeongHyeon/settlement-supporter/src/user/controller/dto"
)

// Create : User Create Controller
func Create(context *fiber.Ctx) error {
	request := new(userControllerDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	return context.Send(context.Body())
}
