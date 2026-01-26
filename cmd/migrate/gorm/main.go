package main

import (
	"fmt"
	"log"
	"os"
	"time"

	gormMysql "lunba-e-commerce/internal/infrastructure/gorm/integration/mysql"

	"gorm.io/gorm"
)

type SchemaMigration struct {
	ID			string	`gorm:"primaryKey"`
	AppliedAt 	time.Time
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: migrate [up|down]")
	}

	command := os.Args[2]

	dsn := "root:toor@tcp(mysql:3306)/ln-e-commerce_order_payment?parseTime=true"

	db, err := gormMysql.NewMysqlDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&SchemaMigration{}); err != nil {
		log.Fatal(err)
	}

	migrations := gormMysql.Models()

	switch command {
	case "up":
		runUp(db, migrations)
	case "down":
		if len(os.Args) < 4 {
			log.Fatal("usage: migrate down {migration_name}")
		}
		
		migration_id := os.Args[3]

		runDown(db, migrations, migration_id)
	}

	log.Println("migration completed successfully")
}

func runUp(db *gorm.DB, migrations []gormMysql.Migration) {
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
}

func runDown(db *gorm.DB, migrations []gormMysql.Migration, m_id string) {
	for i := len(migrations) - 1; i >= 0; i-- {
        m := migrations[i]

		if m_id == m.ID() {
			break
		}

		var count int64
        db.Model(&SchemaMigration{}).
            Where("id = ?", m.ID()).
            Count(&count)

        if count == 0 {
            fmt.Println("Skipping", m.ID())
            continue
        }

        fmt.Println("Reverting", m.ID())
        if err := m.Down(db); err != nil {
            panic(err)
        }

        db.Delete(&SchemaMigration{}, "id = ?", m.ID())
    }
}