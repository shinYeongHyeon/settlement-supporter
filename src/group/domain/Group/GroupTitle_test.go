package groupDomainGroup

import "testing"

func TestGroupTitleCreate(t *testing.T) {
	groupTitleString := "그룹1"
	groupTitle, err := GroupTitleCreate(groupTitleString)

	if err != nil || groupTitle.Value() != groupTitleString {
		t.Fatal("Fail to Create GroupTitle")
	}
}

func TestGroupTitleWhenGroupTitleStringIsEmpty(t *testing.T) {
	_, err := GroupTitleCreate("")

	if err == nil {
		t.Fatal("Create GroupTitle fail when requestString was empty")
	}
}
