package order

import (
	repository "github.com/williamluisan/lunba-e-commerce/internal/domain/repository/order"
	"gorm.io/gorm"
)

type MysqlGormOrder struct {
	gorm *gorm.DB
}

func New(r *repository.OrderRepository) MysqlGormOrder {

}

