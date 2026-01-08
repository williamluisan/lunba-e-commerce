package migrations

import (
	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type CreateProduct struct{}

func (m CreateProduct) ID() string {
	return "20260108162100_create_product_table"
}

func (m CreateProduct) Up(db *gorm.DB) error {
	return db.Migrator().CreateTable(&gormModel.ProductModel{})
}

func (m CreateProduct) Down(db *gorm.DB) error {
	var mdl gormModel.ProductModel

	return db.Migrator().DropTable(mdl.TableName())
}