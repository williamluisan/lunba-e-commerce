package gorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
)

type UserModel struct {
	ID			uuid.UUID   `gorm:"primaryKey;unique;not null"`
	FirstName	string 		`gorm:"type:varchar(255);not null;"`
	LastName	string		`gorm:"type:varchar(255);not null;"`
	Email		string 		`gorm:"type:varchar(255);uniqueIndex;not null"`
	Password	string
	CreatedAt	time.Time
	CreatedBy	int
}

func (UserModel) TableName() string {
	return "users"
}

// convert entity to model
func (m *UserModel) FromEntity(u *entity.User) *UserModel {
	return &UserModel{
		ID: u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
		Password: u.Password,
		CreatedAt: u.CreatedAt,
		CreatedBy: 1, // get from logged in user
	}
}

// reverse model to entity
func (m *UserModel) ToEntity() *entity.User {
	return &entity.User{
		ID: m.ID,
		FirstName: m.FirstName,
		LastName: m.LastName,
		Email: m.Email,
		Password: m.Password,
		CreatedAt: m.CreatedAt,
	}
}