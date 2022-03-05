package userDomain

import core "github.com/shinYeongHyeon/settlement-supporter/src/core/uuid"

// User domain type
type User struct {
	Uuid     string
	Id       Id
	Password Password
	NickName NickName
}

// NewUserProps domain type for new one
type NewUserProps struct {
	Id       Id
	Password Password
	NickName NickName
}

// NewUserCreate create new user
func NewUserCreate(props NewUserProps) (User, error) {
	return User{
		Uuid:     core.CreateUuid(),
		Id:       props.Id,
		Password: props.Password,
		NickName: props.NickName,
	}, nil
}

// GetUuid get uuid
func (user User) GetUuid() string {
	return user.Uuid
}

// GetId get id
func (user User) GetId() Id {
	return user.Id
}

// GetPassword get password
func (user User) GetPassword() Password {
	return user.Password
}

// GetNickName get nickName
func (user User) GetNickName() NickName {
	return user.NickName
}
