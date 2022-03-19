package groupRepository

import (
	core "github.com/shinYeongHyeon/settlement-supporter/src/core/postgres"
	groupDomain "github.com/shinYeongHyeon/settlement-supporter/src/group/domain"
	groupEntity "github.com/shinYeongHyeon/settlement-supporter/src/group/entity"
	"gorm.io/gorm"
	"log"
)

type GroupRepository struct {
	Repository *gorm.DB
}

var groupRepository *GroupRepository

func init() {
	log.Println("GroupRepository init")
	groupRepository = &GroupRepository{core.GetManager().Db.Table("groups")}
}

// GetRepository : Get groupRepository
func GetRepository() *GroupRepository {
	return groupRepository
}

// Create : create Group Row
func (groupRepository *GroupRepository) Create(group groupDomain.Group) error {
	groupRepository.Repository.Create(&groupEntity.Group{
		UUID:  group.GetUuid(),
		Title: group.GetGroupTitle().Value(),
	})

	return nil
}
