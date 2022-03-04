package user

import "testing"

// TestPasswordCreate Testing for Password Create Successful
func TestPasswordCreate(t *testing.T) {
	passwordString := "abcd"
	password, err := PasswordCreate(passwordString)

	if err != nil || password.Value() != passwordString {
		t.Fatal("Fail to Create Password")
	}
}

// TestPasswordCreateFailWhenPasswordStringIsEmpty Testing for passwordString is not empty
func TestPasswordCreateFailWhenPasswordStringIsEmpty(t *testing.T) {
	passwordString := ""
	_, err := PasswordCreate(passwordString)

	if err == nil {
		t.Fatal("passwordString is not empty string")
	}
}
