package migrations

import (
	gormModel "lunba-e-commerce/internal/infrastructure/gorm/model"

	"gorm.io/gorm"
)

type UpdateOrderTable struct {}

func (m UpdateOrderTable) ID() string {
	return "20260126174600_update_order_table"
}

var columns = []string{"PublicId", "CreatedBy", "UpdatedBy"}

func (m UpdateOrderTable) Up(db *gorm.DB) error {
	for _, v := range columns {
		if err := db.Migrator().AlterColumn(&gormModel.OrderModel{}, v); err != nil {
			return err
		}
	}

	return nil
}

func (m UpdateOrderTable) Down(db *gorm.DB) error {
	return nil
}