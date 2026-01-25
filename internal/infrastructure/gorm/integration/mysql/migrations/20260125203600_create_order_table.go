package migrations

import (
	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
	"gorm.io/gorm"
)
	

type CreateOrderTable struct {}

func (m CreateOrderTable) ID() string {
	return "20260125203600_create_order_table"
}

func (m CreateOrderTable) Up(db *gorm.DB) error {
	return db.Migrator().CreateTable(&gormModel.OrderModel{})
}

func (m CreateOrderTable) Down(db *gorm.DB) error {
	var mdl gormModel.OrderModel

	return db.Migrator().DropTable(mdl.TableName())
}