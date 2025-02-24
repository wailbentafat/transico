package models

import "gorm.io/gorm"



type Complain struct {
	gorm.Model

\

	Type         string `gorm:"not null"`
	ComplainDesc string `gorm:"not null"`
	RouteDist    string `gorm:"not null"`
	Price        int    `gorm:"not null"`
}
