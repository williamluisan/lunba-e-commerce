package main

import (
	"log"

	gormInfra "project/internal/infrastructure/gorm"
	"project/internal/infrastructure/mysql"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/app?parseTime=true"

	db, err := mysql.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := gormInfra.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("migration completed successfully")
}