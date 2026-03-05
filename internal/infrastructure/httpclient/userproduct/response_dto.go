package userproduct

type GetUserResponse struct {
	PublicId 	string	`json:"public_id"`
	Name 		string	`json:"name"`
	Email		string 	`json:"email"`
	CreatedAt	string	`json:"created_at"`
	UpdatedAt	string	`json:"updated_at"`
}

type GetProductResponse struct {
	PublicId	string 	`json:"public_id"`
	Name		string 	`json:"name"`
	Code		string 	`json:"code"`
	Price		string  `json:"price"`
	Stock		int		`json:"stock"`
	Status		string 	`json:"status"`	
	CreatedAt	string	`json:"created_at"`
	UpdatedAt	string	`json:"updated_at"`
}