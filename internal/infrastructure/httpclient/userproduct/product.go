package userproduct

import (
	"context"
	"encoding/json"
	entity "lunba-e-commerce/internal/domain/entity/product"
	repositoryProduct "lunba-e-commerce/internal/domain/repository/product"
	"lunba-e-commerce/internal/infrastructure/httpclient"
	"strconv"
	"time"

	restaaRepo "lunba-e-commerce/pkg/restaa"

	"github.com/spf13/viper"
)

type productImpl struct {
	baseURL string
}

func NewProduct() repositoryProduct.ProductRepositoryExt {
	return &productImpl{
		baseURL: viper.GetString("USERPRODUCT_SERVICE_BASE_URL"),
	}
}

func (i *productImpl) GetByPublicId(ctx context.Context, publicId string) (*entity.Product, error) {
	token, _ := ctx.Value("authorization").(string);

	rest := restaaRepo.New(i.baseURL)

	resp, err := rest.Get(ctx, viper.GetString("USERPRODUCT_SERVICE_PRODUCT_EP")+"/"+publicId,
		restaaRepo.WithHeader("Authorization", token), restaaRepo.WithHeader("Content-Type", "application/json"))
	if err != nil {
		return nil, err
	}

	var res httpclient.Response[GetProductResponse]
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	price, _ := strconv.ParseFloat(res.Data.Price, 64)
	createdAt, _ := time.Parse("2006-01-02 15:04:05", res.Data.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", res.Data.UpdatedAt)

	product := &entity.Product{
		PublicId: res.Data.PublicId,
		Name: res.Data.Name,
		Code: res.Data.Code,
		Price: price,
		Stock: res.Data.Stock,
		Status: res.Data.Status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return product, err
}