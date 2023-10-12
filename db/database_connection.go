package db

import (
	"fmt"
	"raedmajeed/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func ConnnectToDB() *gorm.DB {

	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil
	}
	db.AutoMigrate(&model.Entity{})
	fmt.Println("Connected to DB")
	return db
	
}