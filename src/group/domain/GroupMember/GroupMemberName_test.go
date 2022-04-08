package groupDomainGroupMember

import "testing"

func TestGroupMemberNameCreate(t *testing.T) {
	groupMemberNameString := "신영현"
	groupMemberName, err := GroupMemberNameCreate(groupMemberNameString)

	if err != nil || groupMemberName.Value() != groupMemberNameString {
		t.Fatal("Fail to Create GroupMemberName")
	}
}

func TestGroupMemberNameWhenGroupTitleStringIsEmpty(t *testing.T) {
	_, err := GroupMemberNameCreate("")

	if err == nil {
		t.Fatal("Create GroupMemberName fail when requestString was empty")
	}
}
