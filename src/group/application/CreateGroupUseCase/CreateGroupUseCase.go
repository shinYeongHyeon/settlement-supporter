package createGroupUseCase

import (
	createGroupUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase/dto"
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"
)

// Exec : Run CreateUserUseCase
func Exec(request createGroupUseCaseDto.CreateGroupUseCaseRequest) createGroupUseCaseDto.CreateGroupUseCaseResponse {
	group, err := groupDomain.NewGroupCreate(groupDomain.NewGroupProps{
		Title: request.Title,
	})

	if err != nil {
		return createGroupUseCaseDto.CreateGroupUseCaseResponse{
			Code:  "FAIL",
			Group: groupDomain.Group{},
		}
	}

	request.GroupRepository.Create(group)

	return createGroupUseCaseDto.CreateGroupUseCaseResponse{
		Code:  "SUCCESS",
		Group: group,
	}
}
