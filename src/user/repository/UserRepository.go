package userRepository

import (
	userDomain "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"gorm.io/gorm"
)

// Create : create User Row
func Create(db *gorm.DB, user userDomain.User) error {
	db.Create(&userEntity.User{
		UUID:     user.GetUuid(),
		Id:       user.GetId().Value(),
		Password: user.GetPassword().Value(),
		NickName: user.GetNickName().Value(),
	})

	return nil
}
