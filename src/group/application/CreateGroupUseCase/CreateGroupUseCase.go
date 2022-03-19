package createGroupUseCase

import (
	"fmt"
	createGroupUseCaseDto "github.com/shinYeongHyeon/settlement-supporter/src/group/application/CreateGroupUseCase/dto"
)

// Exec : Run CreateUserUseCase
func Exec(request createGroupUseCaseDto.CreateGroupUseCaseRequest) {
	fmt.Println(request)
}
