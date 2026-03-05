package userproduct

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	entity "lunba-e-commerce/internal/domain/entity/user"
	repositoryUser "lunba-e-commerce/internal/domain/repository/user"
	"lunba-e-commerce/internal/infrastructure/httpclient"

	restaaRepo "lunba-e-commerce/pkg/restaa"

	"github.com/spf13/viper"
)

type userImpl struct {
	baseURL string
	client *http.Client
}

func NewUser() repositoryUser.UserRepositoryExt {
	return &userImpl{
		baseURL: viper.GetString("USERPRODUCT_SERVICE_BASE_URL"),
		client: &http.Client{},
	}
}

func (i *userImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.User, error) {
	token, _ := ctx.Value("authorization").(string);

	rest := restaaRepo.New(i.baseURL)

	resp, err := rest.Get(ctx, viper.GetString("USERPRODUCT_SERVICE_USER_EP")+"/"+publicId,
		restaaRepo.WithHeader("Authorization", token), restaaRepo.WithHeader("Content-Type", "application/json"))
	if err != nil {
		return nil, err
	}

	var res httpclient.Response[GetUserResponse]
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	createdAt, _ := time.Parse("2006-01-02 15:04:05", res.Data.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", res.Data.UpdatedAt)

	user := &entity.User{
		PublicId: res.Data.PublicId,
		Name: res.Data.Name,
		Email: res.Data.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return user, err
}