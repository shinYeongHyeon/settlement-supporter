package coreUuid

import (
	"reflect"
	"testing"
)

// TestCreateUuid Success For Create UUID
func TestCreateUuid(t *testing.T) {
	createdUuid := CreateUuid()
	if createdUuid == "" || reflect.TypeOf(createdUuid).String() != "string" {
		t.Fatal("Fail to create uuid")
	}
}
