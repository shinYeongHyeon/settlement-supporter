package groupControllerCommand

import (
	"github.com/gofiber/fiber/v2"
	createGroupUseCase "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase"
	createGroupUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase/dto"
	groupControllerCommandDto "github.com/shinYeongHyeon/settlement-supporter/src/group/controller/command/dto"
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain/Group"
	groupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/postgres"
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

	if groupTitleOrError != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(groupControllerCommandDto.CreateResponse{
			Code: "FAILED",
		})
	}

	createGroupUseCaseResponse := createGroupUseCase.Exec(
		createGroupUseCaseDto.CreateGroupUseCaseRequest{
			Title: groupTitle,
		},
		groupRepository.GetRepository(),
	)

	return context.Status(fiber.StatusCreated).JSON(groupControllerCommandDto.CreateResponse{
		Code: createGroupUseCaseResponse.Code,
		Group: groupControllerCommandDto.Group{
			Uuid:  createGroupUseCaseResponse.Group.GetUuid(),
			Title: createGroupUseCaseResponse.Group.GetGroupTitle().Value(),
		},
	})
}
