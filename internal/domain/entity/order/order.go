package order

import (
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type Order struct {
	ID				uint64
	PublicId		string
	UserPublicId	string
	ProductPublicId	string
	IsPaid			bool
	CreatedAt		time.Time
	CreatedBy		string
	UpdatedAt		time.Time
	UpdatedBy		string
}

type OrderInput struct {
	PublicId		string
	UserPublicId	string
	ProductPublicId	string
	// CreatedAt		time.Time // should be automatically added by gorm
	CreatedBy		string // user public id
}

func NewOrder(input *OrderInput) *Order {
	return &Order{
		PublicId: ulid.Make().String(),
		UserPublicId: input.UserPublicId,
		ProductPublicId: input.ProductPublicId,
		CreatedBy: ulid.Make().String(), // TODO: get from token/session
	}
}