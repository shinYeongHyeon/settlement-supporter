package userDomain

import "testing"

// TestNewUserCreate Testing for NewUserCreate Create Successful
func TestNewUserCreate(t *testing.T) {
	idString := "id"
	passwordString := "abcd"
	nickNameString := "guest"

	userProps := NewUserProps{
		Id:       Id{idString},
		Password: Password{passwordString},
		NickName: NickName{nickNameString},
	}
	user, err := NewUserCreate(userProps)

	if err != nil || user.GetUuid() == "" || user.GetId().Value() != idString || user.GetNickName().Value() != nickNameString || user.GetPassword().Value() != passwordString {
		t.Fatal("Fail to Create UserEmail")
	}
}
