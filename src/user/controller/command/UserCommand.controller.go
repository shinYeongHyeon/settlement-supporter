package command

import (
	"github.com/gofiber/fiber/v2"
	userControllerCommandDto "github.com/shinYeongHyeon/settlement-supporter/src/user/controller/command/dto"
)

// Create : Create user
// @Summary      Create an account
// @Description  Create an account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      userControllerCommandDto.CreateRequest        true  "CreateUserRequest"
// @Success      201   {object}  userControllerCommandDto.CreateResponse       "CreateUserResponse"
// @Failure      400   {object}  userControllerCommandDto.CreateResponseError  "CreateUserErrorResponse"
// @Router       /users [post]
func Create(context *fiber.Ctx) error {
	request := new(userControllerCommandDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	if request.Id == "" || request.NickName == "" || request.Password == "" {
		return context.Status(fiber.StatusBadRequest).JSON(userControllerCommandDto.CreateResponseErrorForBadRequest)
	}

	response := userControllerCommandDto.CreateResponse{
		Code: "SUCCESS",
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
