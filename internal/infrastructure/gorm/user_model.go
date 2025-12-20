package gorm

import (
	"time"

	"github.com/google/uuid"
)

// type UserModel struct {
// 	ID	uuid.UUID   `gorm:"type:uuid;primaryKey;unique;not null"`
// 	FirstName	string `gorm:"type:varchar(255); not null;"`
// 	LastName	string	`gorm:"type:varchar(255); not null;"`
// 	Email	string `gorm:"uniqueIndex"`
// 	Password	string
// 	CreatedAt	time.Time
// 	CreatedBy	int
// }

type UserModel struct {
	ID	uuid.UUID   
	FirstName	string 
	LastName	string	
	Email	string 
	Password	string
	CreatedAt	time.Time
	CreatedBy	int
}

func (UserModel) TableName() string {
	return "users"
}