package migrations

import (
	"gorm.io/gorm"

	gormModel "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/model"
)

type CreateUser struct{}

func (m CreateUser) ID() string {
    return "20251221074801_create_user_table"
}

func (m CreateUser) Up(db *gorm.DB) error {
    return db.Migrator().CreateTable(&gormModel.UserModel{})
}

func (m CreateUser) Down(db *gorm.DB) error {
    var mdl gormModel.UserModel

    return db.Migrator().DropTable(mdl.TableName())
}
