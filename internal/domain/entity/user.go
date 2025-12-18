package entity

import "github.com/google/uuid"

type User struct {
	ID	uuid.UUID
	FirstName	string
	LastName	string
	Email	string
	Password string
	CreatedAt	time.Time
}

// DTO (Data Transfer Object)
type UserInput struct {
	FirstName	string
	LastName	string
	Email		string
	Password	string
}

func NewUser(id uuid.UUID, firstName, lastName, email, password string) *User {
	return &User{
		ID:	id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		CreatedAt: time.Now()
	}
}