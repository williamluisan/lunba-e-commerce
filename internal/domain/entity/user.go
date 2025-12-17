package entity

type User struct {
	ID	uuid.UUID
	FirstName	string
	LastName	string
	Email	string
	Password string
	CreatedAt	time.Time
}