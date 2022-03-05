package userDomain

import "errors"

// NickName is NickName ValueObject
type NickName struct {
	value string
}

// NickNameCreate create NickName with nickNameString
func NickNameCreate(nickNameString string) (NickName, error) {
	if nickNameString == "" {
		return NickName{}, errors.New("nickNameString is not empty")
	}

	return NickName{nickNameString}, nil
}

// Value get value
func (nickName NickName) Value() string {
	return nickName.value
}
