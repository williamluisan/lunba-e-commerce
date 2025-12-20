package mysql

import (
	gormMysqlMigrations "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/mysql/migrations"
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
		gormMysqlMigrations.CreateUsers{},
		gormMysqlMigrations.CreateProducts{},
	}
}
