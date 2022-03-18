package groupEntity

import (
	"time"
)

// Group is GroupEntity
type Group struct {
	UUID      string `gorm:"primaryKey"`
	Title     string
	CreatedAt time.Time `sql:"DEFAULT:'current_timestamp'"`
}
