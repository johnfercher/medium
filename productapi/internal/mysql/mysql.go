package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(url string, dbName string, adminUser string, adminPassword string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", adminUser, adminPassword, url, dbName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
