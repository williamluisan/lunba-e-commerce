package migrations

import (
	"gorm.io/gorm"

	gormInfra "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm"
)

type CreateProducts struct{}

func (m CreateProducts) ID() string {
	return "20251221074802_create_product_table"
}

func (m CreateProducts) Up(db *gorm.DB) error {
    return db.Migrator().CreateTable(&gormInfra.ProductModel{})
}

func (m CreateProducts) Down(db *gorm.DB) error {
    return db.Migrator().DropTable("users")
}

