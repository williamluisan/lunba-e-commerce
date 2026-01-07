package user

type CreateUserRequest struct {
	FirstName	string	`form:"firstName" json:"firstName" binding:"required"`
	LastName	string	`form:"lastName" json:"lastName" binding:"required"`
	Email		string	`form:"email" json:"email" binding:"required,email"`
	Password	string	`form:"password" json:"password" binding:"required"`
}