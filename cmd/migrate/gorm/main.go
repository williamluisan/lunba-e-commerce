package main

import (
	"fmt"
	"log"
	"time"

	gormMysql "github.com/williamluisan/lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"
)

type SchemaMigration struct {
	ID			string	`gorm:"primaryKey"`
	AppliedAt 	time.Time
}

func main() {
	dsn := "lunba:toor@tcp(localhost:3306)/lunba_e_commerce?parseTime=true"

	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&SchemaMigration{}); err != nil {
		log.Fatal(err)
	}

	migrations := gormMysql.Models()

	for _, m := range migrations {
        var count int64
        db.Model(&SchemaMigration{}).
            Where("id = ?", m.ID()).
            Count(&count)

        if count > 0 {
            fmt.Println("Skipping", m.ID())
            continue
        }

        fmt.Println("Applying", m.ID())
        if err := m.Up(db); err != nil {
            panic(err)
        }

        db.Create(&SchemaMigration{
            ID:        m.ID(),
            AppliedAt: time.Now(),
        })
    }


	log.Println("migration completed successfully")
}