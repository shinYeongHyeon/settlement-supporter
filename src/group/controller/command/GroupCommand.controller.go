package groupControllerCommand

import (
	"github.com/gofiber/fiber/v2"
	createGroupUseCase "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase"
	createGroupUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase/dto"
	groupControllerCommandDto "github.com/shinYeongHyeon/settlement-supporter/src/group/controller/command/dto"
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"
)

// Create : Create group
// @Summary      Create a group
// @Description  Create a group
// @Tags         Group
// @Accept       json
// @Produce      json
// @Param        group  body      groupControllerCommandDto.CreateRequest        true  "CreateGroupRequest"
// @Success      201    {object}  groupControllerCommandDto.CreateResponse       "CreateGroupResponse"
// @Failure      400    {object}  groupControllerCommandDto.CreateResponseError  "CreateGroupErrorResponse"
// @Router       /groups [post]
func Create(context *fiber.Ctx) error {
	request := new(groupControllerCommandDto.CreateRequest)

	if err := context.BodyParser(request); err != nil {
		return err
	}

	if request.Title == "" {
		return context.Status(fiber.StatusBadRequest).JSON(groupControllerCommandDto.CreateResponseErrorForBadRequest)
	}

	groupTitle, groupTitleOrError := groupDomain.GroupTitleCreate(request.Title)

	// TODO: 에러 표준화
	if groupTitleOrError != nil {
	}

	createGroupUseCase.Exec(createGroupUseCaseDto.CreateGroupUseCaseRequest{
		Title: groupTitle,
	})

	response := groupControllerCommandDto.CreateResponse{
		Code: "SUCCESS",
	}

	return context.Status(fiber.StatusCreated).JSON(response)
}
