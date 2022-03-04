package user

import "errors"

// Id is id ValueObject
type Id struct {
	value string
}

// IdCreate create Id with idString
func IdCreate(idString string) (Id, error) {
	if idString == "" {
		return Id{}, errors.New("idString is not empty")
	}

	return Id{idString}, nil
}

// Value get value
func (id Id) Value() string {
	return id.value
}
