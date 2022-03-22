package createGroupUseCaseDto

import (
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"
	groupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/postgres"
)

// CreateGroupUseCaseRequest : Request of CreateGroupUseCase
type CreateGroupUseCaseRequest struct {
	Title           groupDomain.GroupTitle
	GroupRepository *groupRepository.GroupRepository
}

// CreateGroupUseCaseResponse : Response of CreateGroupUseCase
type CreateGroupUseCaseResponse struct {
	Code  string
	Group groupDomain.Group
}
