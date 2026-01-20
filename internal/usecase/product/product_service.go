package product

import (
	"context"
	"errors"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/product"
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/product"
)

type ProductService interface {
	Create(ctx context.Context, input *entity.ProductInput) error
	Delete(ctx context.Context, id int64) error
	GetByCode(ctx context.Context, code string) (*entity.Product, error)
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) ProductService {
	if r == nil {
		panic("product repository cannot be nil")
	}
	return &ProductServiceImpl{
		repo: r,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, input *entity.ProductInput) error {
	// check if product code is already exists
	code := input.Code
	product, _ := s.repo.GetByCode(ctx, code)
	if product != nil {
		return errors.New("the product with the corresponding code is already exists.")
	}

	data := &entity.ProductInput{
		Name: input.Name,
		Code: input.Code, 
		Price: input.Price,
	}
	newProduct := entity.NewProduct(data)

	if err := s.repo.Create(ctx, newProduct); err != nil {
		return err
	}

	return nil
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id int64) error {
	return nil
}

func (s *ProductServiceImpl) GetByCode(ctx context.Context, code string) (*entity.Product, error) {
	return nil, nil
}