package userproduct

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	entity "lunba-e-commerce/internal/domain/entity/user"
	repositoryUser "lunba-e-commerce/internal/domain/repository/user"
	httpclient "lunba-e-commerce/internal/infrastructure/httpclient"

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
	url := fmt.Sprintf(
		"%s%s/%s",
		i.baseURL,
		viper.GetString("USERPRODUCT_SERVICE_USER_EP"),
		publicId,
	)

	// TODO: Create http helper
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil, // no body
	)
	if err != nil {
		return nil, err
	}
	if token, ok := ctx.Value("authorization").(string); ok && token != "" {
		req.Header.Set("Authorization", token)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get user detail")
	}

	var res httpclient.Response[GetUserResponse]
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
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