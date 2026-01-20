package repository

import (
	"context"

	"github.com/williamluisan/lunba-e-commerce/internal/domain/entity"
)

type PaymenRepository interface {
	Process(ctx context.Context, e *entity.Payment) error
}