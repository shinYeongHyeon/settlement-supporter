package testGroupRepository

import (
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"
	"log"
)

// TestDB : [test] DB struct
type TestDB struct{}

// TestGroupRepository : [test] GroupRepository
type TestGroupRepository struct {
	Repository TestDB
}

var testGroupRepository *TestGroupRepository

func init() {
	log.Println("TestGroupRepository init")
	testGroupRepository = &TestGroupRepository{}
}

// GetRepository : [test] Get groupRepository
func GetRepository() *TestGroupRepository {
	return testGroupRepository
}

// Create : [test] create Group Row
func (groupRepository *TestGroupRepository) Create(_ groupDomain.Group) (bool, error) {
	return true, nil
}
