package models

import "gorm.io/gorm"



type Complain struct {
	gorm.Model

	CompanyID *uint   `gorm:"index;default:null"`
	Company   *CmpData `gorm:"foreignKey:CompanyID;constraint:OnDelete:SET NULL"`

	Type         string `gorm:"not null"`
	ComplainDesc string `gorm:"not null"`
}
