package Group

import "errors"

// GroupTitle : GroupTitle ValueObject
type GroupTitle struct {
	value string
}

// GroupTitleCreate : create GroupTitle with groupTitleString
func GroupTitleCreate(groupTitleString string) (GroupTitle, error) {
	if groupTitleString == "" {
		return GroupTitle{}, errors.New("groupTitleString is not empty")
	}

	return GroupTitle{groupTitleString}, nil
}

// Value : get value
func (groupTitle GroupTitle) Value() string {
	return groupTitle.value
}
