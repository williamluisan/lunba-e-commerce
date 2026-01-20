package user

import (
	"context"
	"errors"
	"fmt"

	entity "github.com/williamluisan/lunba-e-commerce/internal/domain/entity/user"
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, u *entity.UserInput) error
	Update(ctx context.Context, u *entity.UserInput) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	if r == nil {
		panic("user repository cannot be nil")
	}
	return &userServiceImpl{
		repo: r,
	}
}

func (s *userServiceImpl) GetById(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) Create(ctx context.Context, input *entity.UserInput) error {
	// check duplicate email
	existingUser, _ := s.repo.GetByEmail(ctx, input.Email)
	if (existingUser != nil) {
		return errors.New("the email is already in use")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if (err != nil) {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// create the domain entity
	data := &entity.UserInput{
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
		Password: string(hashedPassword),
	}
	newUser := entity.NewUser(data)

	if err := s.repo.Create(ctx, newUser); err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Update(ctx context.Context, input *entity.UserInput) error {
	// ...
	return nil
}