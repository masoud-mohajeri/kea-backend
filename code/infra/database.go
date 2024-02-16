package infra

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masoud-mohajeri/kea-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DatabaseConfig.Host,
		config.DatabaseConfig.User,
		config.DatabaseConfig.Pass,
		config.DatabaseConfig.Name,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.SSLMode,
		config.DatabaseConfig.TimeZone,
	)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic("Can not connect to database.")
	}

	DB = db
}
