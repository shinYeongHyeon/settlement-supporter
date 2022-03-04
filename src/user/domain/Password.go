package user

import "errors"

// Password is password ValueObject
type Password struct {
	value string
}

// PasswordCreate create Password with passwordString
func PasswordCreate(passwordString string) (Password, error) {
	if passwordString == "" {
		return Password{}, errors.New("passwordString is not empty")
	}

	return Password{passwordString}, nil
}

// Value get value
func (password Password) Value() string {
	return password.value
}
