package userproduct

type GetUserResponse struct {
	PublicId 	string	`json:"public_id"`
	Name 		string	`json:"name"`
	Email		string 	`json:"email"`
	CreatedAt	string	`json:"created_at"`
	UpdatedAt	string	`json:"updated_at"`
}