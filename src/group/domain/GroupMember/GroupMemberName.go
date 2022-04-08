package groupDomainGroupMember

import "errors"

// GroupMemberName : GroupMemberName ValueObject
type GroupMemberName struct {
	value string
}

// GroupMemberNameCreate : create GroupMemberName with groupMemberNameString
func GroupMemberNameCreate(groupMemberNameString string) (GroupMemberName, error) {
	if groupMemberNameString == "" {
		return GroupMemberName{}, errors.New("groupTitleString is not empty")
	}

	return GroupMemberName{groupMemberNameString}, nil
}

// Value : get value
func (groupMemberName GroupMemberName) Value() string {
	return groupMemberName.value
}
