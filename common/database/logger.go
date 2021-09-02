package database

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var ormLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,       // Disable color
	},
)

