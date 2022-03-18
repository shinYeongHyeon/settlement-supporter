package groupEntity

import (
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"time"
)

// GroupUsers : GroupUsersEntity
type GroupUsers struct {
	UUID      string          `gorm:"primaryKey"`
	GroupId   string          ``
	Group     Group           `gorm:"foreignKey:GroupId;AssociationForeignKey:UUID"`
	UserId    string          ``
	User      userEntity.User `gorm:"foreignKey:UserId;AssociationForeignKey:UUID"`
	CreatedAt time.Time       `sql:"DEFAULT:'current_timestamp'"`
}
