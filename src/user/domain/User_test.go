package user

import "testing"

// TestNewUserCreate Testing for NewUserCreate Create Successful
func TestNewUserCreate(t *testing.T) {
	idString := "id"
	passwordString := "abcd"
	nickNameString := "guest"

	userProps := NewUserProps{
		id:       Id{idString},
		password: Password{passwordString},
		nickName: NickName{nickNameString},
	}
	user, err := NewUserCreate(userProps)

	if err != nil || user.Uuid() == "" || user.Id().Value() != idString || user.NickName().Value() != nickNameString || user.Password().Value() != passwordString {
		t.Fatal("Fail to Create UserEmail")
	}
}
