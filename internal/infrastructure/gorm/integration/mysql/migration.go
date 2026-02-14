package mysql

import (
	gormMysqlMigrations "lunba-e-commerce/internal/infrastructure/gorm/integration/mysql/migrations"

	"gorm.io/gorm"
)	

type Migration interface {
    ID() string
    Up(db *gorm.DB) error
    Down(db *gorm.DB) error
}

// list all models
func Models() []Migration{
	return []Migration{
		// gormMysqlMigrations.CreateUser{},
		// gormMysqlMigrations.CreateProduct{},
		// gormMysqlMigrations.UpdateProductCodeColumn{},
		// gormMysqlMigrations.UpdateProductPriceColumnType{},
		gormMysqlMigrations.CreateOrderTable{},
		gormMysqlMigrations.UpdateOrderTable{},
		gormMysqlMigrations.UpdateOrderTableAddQtyColumn{},
	}
}
