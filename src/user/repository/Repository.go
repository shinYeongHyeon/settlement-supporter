package userRepository

import (
	userDomain "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"gorm.io/gorm"
)

var userDb *gorm.DB

func Init(db *gorm.DB) {
	userDb = db
}

func Create(user userDomain.User) {
	userDb.Create(&userEntity.User{
		UUID:     user.GetUuid(),
		Id:       user.GetId().Value(),
		Password: user.GetPassword().Value(),
		NickName: user.GetNickName().Value(),
	})
}
