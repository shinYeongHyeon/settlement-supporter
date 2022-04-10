package groupEntity

import (
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"time"
)

// Group is GroupEntity
type Group struct {
	UUID         string          `gorm:"primaryKey"`
	Title        string          ``
	GroupOwnerId string          ``
	GroupOwner   userEntity.User `gorm:"foreignKey:GroupOwnerId;AssociationForeignKey:UUID"`
	CreatedAt    time.Time       `sql:"DEFAULT:'current_timestamp'"`
}
