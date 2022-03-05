package userEntity

import "time"

// User is UserEntity
type User struct {
	UUID      string `gorm:"primaryKey"`
	Id        string
	Password  string
	NickName  string
	CreatedAt time.Time `sql:"DEFAULT:'current_timestamp'"`
}
