package createGroupUseCase

import (
	createGroupUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase/dto"
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain/Group"
	groupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository"
)

// Exec : Run CreateUserUseCase
func Exec[T groupRepository.IGroupRepository](
	request createGroupUseCaseDto.CreateGroupUseCaseRequest,
	repository T,
) createGroupUseCaseDto.CreateGroupUseCaseResponse {
	group, err := groupDomain.NewGroupCreate(groupDomain.NewGroupProps{
		Title: request.Title,
	})

	if err != nil {
		return createGroupUseCaseDto.CreateGroupUseCaseResponse{
			Code:  "FAIL",
			Group: groupDomain.Group{},
		}
	}

	_, createError := repository.Create(group)

	if createError != nil {
		return createGroupUseCaseDto.CreateGroupUseCaseResponse{
			Code:  "FAIL",
			Group: groupDomain.Group{},
		}
	}

	return createGroupUseCaseDto.CreateGroupUseCaseResponse{
		Code:  "SUCCESS",
		Group: group,
	}
}
