package gorm

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID			uuid.UUID   `gorm:"primaryKey;unique;not null"`
	FirstName	string 		`gorm:"type:varchar(255); not null;"`
	LastName	string		`gorm:"type:varchar(255); not null;"`
	Email		string 		`gorm:"type:varchar(255);uniqueIndex;not null"`
	Password	string
	CreatedAt	time.Time
	CreatedBy	int
}

func (UserModel) TableName() string {
	return "users"
}