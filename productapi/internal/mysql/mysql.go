package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(url string, dbName string, adminUser string, adminPassword string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", adminUser, adminPassword, url, dbName)
	fmt.Printf("connecting to mysql %s\n", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("could not connect to mysql")
		return nil, err
	}

	fmt.Println("connected to mysql")

	return db, nil
}
