package groupRepository

import (
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain/Group"
	postgresGroupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/postgres"
	testGroupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/test"
)

// IGroupRepository : DI 용 GroupRepository
type IGroupRepository interface {
	*postgresGroupRepository.GroupRepository | *testGroupRepository.TestGroupRepository
	Create(g groupDomain.Group) (bool, error)
}
