package createUserUseCase

import (
	"fmt"
	"github.com/shinYeongHyeon/settlement-supporter/src/user/application/CreateUserUseCase/dto"
)

// Exec : Run CreateUserUseCase
func Exec(request createUserUseCaseDto.CreateUserUseCaseRequest) {
	fmt.Println(request)
}
