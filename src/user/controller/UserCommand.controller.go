package userController

import (
	"github.com/gofiber/fiber/v2"
	userControllerDto "github.com/shinYeongHyeon/settlement-supporter/src/user/controller/dto"
)

// Create : Create user
// @Summary      Create an account
// @Description  Create an account
// @Tags                User
// @Accept       json
// @Produce      json
// @Param        user  body      userControllerDto.CreateRequest   true  "CreateUserRequest"
// @Success      201   {object}  userControllerDto.CreateResponse  tru   "CreateUserResponse"
// @Router       /users [post]
func Create(context *fiber.Ctx) error {
	request := new(userControllerDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	response := userControllerDto.CreateResponse{
		Code: "SUCCESS",
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
