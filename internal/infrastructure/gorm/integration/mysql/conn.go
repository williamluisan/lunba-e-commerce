package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(dsn string) (*gorm.DB, error) {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,    // Slow SQL threshold
	// 		LogLevel:                  logger.Silent,  // Log level
	// 		IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,           // Don't include params in the SQL log
	// 		Colorful:                  false,          // Disable color
	// 	},
	// )

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
}