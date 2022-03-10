package createUserUseCaseDto

import userDomain "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"

// CreateUserUseCaseRequest : Request of CreateUserUseCase
type CreateUserUseCaseRequest struct {
	Id       userDomain.Id
	NickName userDomain.NickName
	Password userDomain.Password
}

// CreateUserUseCaseResponse : Response of CreateUserUseCase
type CreateUserUseCaseResponse struct {
	Code string
	User userDomain.User
}
