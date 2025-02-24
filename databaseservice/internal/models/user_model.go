package models
import (
	"gorm.io/gorm"
)



type User struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsAdmin  bool   `gorm:"default:false"`

	// A User can own multiple companies (Nullable)
	Companies []*CmpData `gorm:"foreignKey:CreatorID"`
}
