package db

import (
	"databaseservice/internal/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type DB struct {
	DB *gorm.DB
	
}






func Initdb(db_name string)(DB , error){
	conn,err:=gorm.Open(sqlite.Open(db_name), &gorm.Config{})
	if err != nil {
		return DB{},err
	}
	err = conn.AutoMigrate(&models.User{}, &models.CmpData{}, &models.CmpInfo{}, &models.Complain{})
	if err != nil {
		fmt.Println("Error migrating database:", err)
		return DB{},err
	}

	fmt.Println("Database migration completed successfully.")
	return DB{DB: conn},nil
}

	
