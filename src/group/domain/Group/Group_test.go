package Group

import "testing"

// TestNewGroupCreate : Testing for NewGroupCreate Create Successful
func TestNewGroupCreate(t *testing.T) {
	titleString := "타이틀1"

	groupProps := NewGroupProps{
		Title: GroupTitle{titleString},
	}
	group, err := NewGroupCreate(groupProps)

	if err != nil || group.GetUuid() == "" || group.GetGroupTitle().Value() != titleString {
		t.Fatal("Fail to Create Group")
	}
}
