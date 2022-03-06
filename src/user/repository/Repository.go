package userRepository

import (
	"errors"
	userDomain "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"gorm.io/gorm"
)

var userDb *gorm.DB

// Init : userDb instance create
func Init(db *gorm.DB) {
	userDb = db
}

// Create : create User Row
func Create(user userDomain.User) error {
	if userDb == nil {
		return errors.New("database not initialized")
	}

	userDb.Create(&userEntity.User{
		UUID:     user.GetUuid(),
		Id:       user.GetId().Value(),
		Password: user.GetPassword().Value(),
		NickName: user.GetNickName().Value(),
	})

	return nil
}
