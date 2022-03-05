package userDomain

import "testing"

// TestIdCreate Testing for Id Create Successful
func TestIdCreate(t *testing.T) {
	idString := "id"
	id, err := IdCreate(idString)

	if err != nil || id.Value() != idString {
		t.Fatal("Fail to Create Id")
	}
}

// TestIdCreateFailWhenIdStringIsEmpty Testing for idString is not empty
func TestIdCreateFailWhenIdStringIsEmpty(t *testing.T) {
	idString := ""
	_, err := IdCreate(idString)

	if err == nil {
		t.Fatal("idString is not empty string")
	}
}
