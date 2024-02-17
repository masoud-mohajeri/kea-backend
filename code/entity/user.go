package entity

import (
	"time"
)

type User struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"-"`
	Mobile    string     `gorm:"<-;type:varchar(11);not null;unique" json:"mobile"`
	FirstName string     `gorm:"<-;type:varchar(255)" json:"first_name"`
	Password  string     `gorm:"<-;type:varchar(255)" json:"Password"`
	Role      string     `gorm:"<-;type:varchar(255)" json:"Role"`
	LastName  string     `gorm:"<-;type:varchar(255)" json:"last_name"`
	Birthdate *time.Time `gorm:"<-;type:timestamp;default:null" json:"birth_date"`
	CreatedAt time.Time  `gorm:"<-;type:timestamp;not null" json:"-"`
	UpdatedAt time.Time  `gorm:"<-;type:timestamp" json:"-"`
}
