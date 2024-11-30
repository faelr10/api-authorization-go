

package database_pkg

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	InitDatabase(*Config) (*gorm.DB, error)  
	RunMigrations(db *gorm.DB, migrationsPath string) error  
	SetupDatabase(*Config) (*gorm.DB, error)  
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}
