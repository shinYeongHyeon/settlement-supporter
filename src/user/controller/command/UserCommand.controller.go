package userControllerCommand

import (
	"github.com/gofiber/fiber/v2"
	createUserUseCase "github.com/shinYeongHyeon/settlement-supporter/src/user/application/CreateUserUseCase"
	createUserUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/user/application/CreateUserUseCase/dto"
	userControllerCommandDto "github.com/shinYeongHyeon/settlement-supporter/src/user/controller/command/dto"
	userDomain "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
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

	id, idOrError := userDomain.IdCreate(request.Id)
	nickName, nickNameOrError := userDomain.NickNameCreate(request.NickName)
	password, passwordOrError := userDomain.PasswordCreate(request.Password)

	// TODO: 에러 표준화
	if idOrError != nil {
	}
	if nickNameOrError != nil {
	}
	if passwordOrError != nil {
	}

	createUserUseCase.Exec(createUserUseCaseDto.CreateUserUseCaseRequest{
		Id:       id,
		NickName: nickName,
		Password: password,
	})

	response := userControllerCommandDto.CreateResponse{
		Code: "SUCCESS",
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
