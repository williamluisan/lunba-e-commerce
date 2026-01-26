package mysql

import (
	"context"

	"lunba-e-commerce/internal/domain/entity"
	"lunba-e-commerce/internal/domain/repository"
)

type mysqlUserRepositoryImpl struct{
	// add any necessary dependencies or storage logic here
}

func NewMysqlUserRepository() repository.UserRepository {
	return &mysqlUserRepositoryImpl{}
}

func (r *mysqlUserRepositoryImpl) GetById(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}

func (r *mysqlUserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

func (r *mysqlUserRepositoryImpl) Create(ctx context.Context, u *entity.UserInput) error {
	return nil
}

func (r *mysqlUserRepositoryImpl) Update(ctx context.Context, u *entity.UserInput) error {
	return nil
}
