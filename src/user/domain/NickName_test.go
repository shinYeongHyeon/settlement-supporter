package user

import "testing"

// TestNickNameCreate Testing for NickName Create Successful
func TestNickNameCreate(t *testing.T) {
	nickNameString := "Guest"
	nickName, err := NickNameCreate(nickNameString)

	if err != nil || nickName.Value() != nickNameString {
		t.Fatal("Fail to Create NickName")
	}
}

// TestNickNameCreateFailWhenNickNameStringIsEmpty Testing for nickNameString is not empty
func TestNickNameCreateFailWhenNickNameStringIsEmpty(t *testing.T) {
	nickNameString := ""
	_, err := NickNameCreate(nickNameString)

	if err == nil {
		t.Fatal("nickNameString is not empty string")
	}
}
