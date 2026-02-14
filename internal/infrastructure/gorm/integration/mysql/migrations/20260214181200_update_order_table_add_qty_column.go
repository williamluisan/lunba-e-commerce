package migrations

import (
	gormModel "lunba-e-commerce/internal/infrastructure/gorm/model"

	"gorm.io/gorm"
)

type UpdateOrderTableAddQtyColumn struct {}

func (m UpdateOrderTableAddQtyColumn) ID() string {
	return "20260126174600_update_order_add_qty_column"
}

var columns2 = []string{"Qty"}

func (m UpdateOrderTableAddQtyColumn) Up(db *gorm.DB) error {
	for _, v := range columns2 {
		if err := db.Migrator().AddColumn(&gormModel.OrderModel{}, v); err != nil {
			return err
		}
	}

	return nil
}

func (m UpdateOrderTableAddQtyColumn) Down(db *gorm.DB) error {
	return nil
}