package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenSqLiteDatabase(dbPath string, hardLogging bool) (*gorm.DB, error) {
	// Configure GORM
	config := &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// Open SQLite database
	db, err := gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		return nil, err
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if hardLogging {

		// Create a log file
		if err := os.MkdirAll("logs", 0755); err != nil {
			panic(err)
		}

		logFile, err := os.OpenFile(
			fmt.Sprintf("logs/gorm_%s.log", time.Now().Format("2006-01-02_15-04-05")),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		)
		if err != nil {
			panic(err)
		}

		// Configure GORM to use the custom logger
		db.Logger = logger.New(
			log.New(logFile, "\r\n", log.LstdFlags), // Use the file for logging
			logger.Config{
				SlowThreshold:             time.Second, // Log queries slower than this
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Don't log record not found errors
				Colorful:                  false,       // Disable colors for file logging
			},
		)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
