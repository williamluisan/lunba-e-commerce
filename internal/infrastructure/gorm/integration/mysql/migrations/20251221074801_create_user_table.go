package migrations

import (
	"gorm.io/gorm"

	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
)

type CreateUsers struct{}

func (m CreateUsers) ID() string {
    return "20251221074801_create_user_table"
}

func (m CreateUsers) Up(db *gorm.DB) error {
    return db.Migrator().CreateTable(&gormModel.UserModel{})
}

func (m CreateUsers) Down(db *gorm.DB) error {
    return db.Migrator().DropTable("users")
}
