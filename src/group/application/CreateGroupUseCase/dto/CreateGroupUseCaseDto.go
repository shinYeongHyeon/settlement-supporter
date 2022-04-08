package createGroupUseCaseDto

import (
	"github.com/shinYeongHyeon/settlement-supporter/src/group/domain/Group"
)

// CreateGroupUseCaseRequest : Request of CreateGroupUseCase
type CreateGroupUseCaseRequest struct {
	Title Group.GroupTitle
}

// CreateGroupUseCaseResponse : Response of CreateGroupUseCase
type CreateGroupUseCaseResponse struct {
	Code  string
	Group Group.Group
}
