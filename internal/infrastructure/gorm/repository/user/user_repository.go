package user

import (
	"context"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/user"
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/user"
	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

var mUserModel gormModel.UserModel

func (r *userRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := gorm.G[gormModel.UserModel](r.db).Where("email = ?", email).First(ctx); 
	if err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

func (r *userRepository) Create(ctx context.Context, u *entity.User) error {
	data := mUserModel.FromEntity(u)

	if err := gorm.G[gormModel.UserModel](r.db).Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, u *entity.User) error {
	return nil
}