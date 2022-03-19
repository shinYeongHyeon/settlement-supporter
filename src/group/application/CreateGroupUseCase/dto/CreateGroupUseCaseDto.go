package createGroupUseCaseDto

import groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"

// CreateGroupUseCaseRequest : Request of CreateGroupUseCase
type CreateGroupUseCaseRequest struct {
	Title groupDomain.GroupTitle
}

// CreateGroupUseCaseResponse : Response of CreateGroupUseCase
type CreateGroupUseCaseResponse struct {
	Code  string
	Group groupDomain.Group
}
