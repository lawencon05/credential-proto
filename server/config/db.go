package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"lawencon.com/credential/model"
)

var tables = []interface{}{
	&model.UsersDb{},
}

const (
	host    = "103.30.180.34"
	user    = "postgres"
	pass    = "bootlawen123"
	dbname  = "credential"
	port    = 9595
	sslmode = "disable"
)

// Conn for return gorm.DB
func Conn() (*gorm.DB, error) {
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta", host, user, pass, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

// MigrateSchema for auto create tables
func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
