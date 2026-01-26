package order

import "time"

type Order struct {
	ID				uint64
	PublicId		string
	UserPublicId	string
	ProductPublicId	string
	IsPaid			bool
	CreatedAt		time.Time
	CreatedBy		int
	UpdatedAt		time.Time
	UpdatedBy		int
}

type OrderInput struct {
	PublicId		string
	UserPublicId	string
	ProductPublicId	string
	CreatedAt		time.Time
	CreatedBy		int
}