package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID
	PublicId	string
	Name		string
	FirstName	string
	LastName	string
	Email		string
	Password 	string
	CreatedAt	time.Time
	CreatedBy	int
	UpdatedAt	time.Time
	UpdatedBy	int
}

// DTO (Data Transfer Object)
type UserInput struct {
	FirstName	string
	LastName	string
	Email		string
	Password	string
}

func NewUser(input *UserInput) *User {
	return &User{
		ID:	uuid.New(),
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
		Password: input.Password,
		CreatedAt: time.Now(),
	}
}