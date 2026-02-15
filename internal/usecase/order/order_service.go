package order

import (
	"context"

	entity "lunba-e-commerce/internal/domain/entity/order"
	repository "lunba-e-commerce/internal/domain/repository/order"
	repositoryUser "lunba-e-commerce/internal/domain/repository/user"

	ulid "github.com/oklog/ulid/v2"
)

type OrderService interface {
	Get(ctx context.Context, orderPublicId string) (*entity.Order, error)
	Create(ctx context.Context, e *entity.OrderInput) error
	Delete(ctx context.Context, orderPublicId string) error
}

type OrderServiceImpl struct {
	repo repository.OrderRepository
	repoUserExt repositoryUser.UserRepositoryExt
}

func New(r repository.OrderRepository, ruExt repositoryUser.UserRepositoryExt) OrderService {
	if r == nil || ruExt == nil {
		panic("Repository not implemented properly.")
	}
	return &OrderServiceImpl{
		repo: r,
		repoUserExt: ruExt,
	}
}

func (i *OrderServiceImpl) Get(ctx context.Context, orderPublicId string) (*entity.Order, error) {
	return nil, nil
}

func (i *OrderServiceImpl) Create(ctx context.Context, input *entity.OrderInput) error {
	userPublicId := ctx.Value("user_public_id").(string)
	
	data := &entity.OrderInput{
		UserPublicId: userPublicId,
		ProductPublicId: input.ProductPublicId,
	}

	// check if valid ULID
	if _, err := ulid.Parse(data.UserPublicId); err != nil {
		return err
	}
	if _, err := ulid.Parse(data.ProductPublicId); err != nil {
		return err
	}

	// check if user public id is exists
	_, err := i.repoUserExt.GetByPublicId(ctx, userPublicId)
	if err != nil {
		return err
	}

	// TODO: check if product is exists via public id (on laravel microservices)
	// ...

	// TODO: check if stock still available
	// ...

	// check if order already exists
	// result, err := i.repo.GetByProductAndUser(ctx, data.ProductPublicId, data.UserPublicId)
	// if  err != nil {
	// 	return err
	// }
	// if result != nil {
	// 	return &messages.OrderAlreadyExistsError{
	// 		ProductName: "dummy",
	// 	}
	// }

	newOrder := entity.NewOrder(data)

	if err := i.repo.Create(ctx, newOrder); err != nil {
		return err
	}

	return nil
}

func (i *OrderServiceImpl) Delete(ctx context.Context, orderPublicId string) error {
	return nil
}

