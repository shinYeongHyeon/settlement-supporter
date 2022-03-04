package user

import "time"

// Entity is UserEntity
type Entity struct {
	UUID      string `gorm:"primaryKey"`
	Id        string
	Password  string
	NickName  string
	CreatedAt time.Time `sql:"DEFAULT:'current_timestamp'"`
}
