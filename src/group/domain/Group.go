package groupDomain

import core "github.com/shinYeongHyeon/settlement-supporter/src/core/uuid"

// Group : group domain
type Group struct {
	Uuid  string
	Title GroupTitle
}

// NewGroupProps : domain type for new one
type NewGroupProps struct {
	Title GroupTitle
}

// NewGroupCreate : create new group
func NewGroupCreate(props NewGroupProps) (Group, error) {
	return Group{
		Uuid:  core.CreateUuid(),
		Title: props.Title,
	}, nil
}

// GetUuid get uuid
func (group Group) GetUuid() string {
	return group.Uuid
}

// GetGroupTitle : get groupTitle
func (group Group) GetGroupTitle() GroupTitle {
	return group.Title
}
