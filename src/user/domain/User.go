package user

import core "github.com/shinYeongHyeon/settlement-supporter/src/core/uuid"

// User domain type
type User struct {
	uuid     string
	id       Id
	password Password
	nickName NickName
}

// NewUserProps domain type for new one
type NewUserProps struct {
	id       Id
	password Password
	nickName NickName
}

// NewUserCreate create new user
func NewUserCreate(props NewUserProps) (User, error) {
	return User{
		uuid:     core.CreateUuid(),
		id:       props.id,
		password: props.password,
		nickName: props.nickName,
	}, nil
}

func (user User) Uuid() string {
	return user.uuid
}

// Id get id
func (user User) Id() Id {
	return user.id
}

// Password get password
func (user User) Password() Password {
	return user.password
}

// NickName get nickName
func (user User) NickName() NickName {
	return user.nickName
}
