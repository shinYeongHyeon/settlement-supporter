package groupRepository

import (
	postgresGroupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/postgres"
	testGroupRepository "github.com/shinYeongHyeon/settlement-supporter/src/group/repository/test"
)

// DIGroupRepository : DI 용 GroupRepository
type DIGroupRepository interface {
	postgresGroupRepository.GroupRepository | testGroupRepository.TestGroupRepository
}
