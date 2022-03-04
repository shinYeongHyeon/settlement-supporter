package user

import (
	"fmt"
	user "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
	user2 "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"gorm.io/gorm"
)

var userDb *gorm.DB

func Init(db *gorm.DB) {
	userDb = db
}

func Create(user user.User) {
	fmt.Println("user")
	fmt.Println(user)

	userDb.Create(&user2.Entity{
		UUID:     user.GetUuid(),
		Id:       user.GetId().Value(),
		Password: user.GetPassword().Value(),
		NickName: user.GetNickName().Value(),
	})
}
