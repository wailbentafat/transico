package models

import "gorm.io/gorm"

type CmpInfo struct {
	gorm.Model

	CompanyID *uint   `gorm:"uniqueIndex;default:null"`
	Company   *CmpData `gorm:"foreignKey:CompanyID;constraint:OnDelete:SET NULL"`

	WorkSpeciality    string `gorm:"not null"`
	Address          string `gorm:"not null"`
	Email            string `gorm:"not null"`
	Phone            string `gorm:"not null"`
	Description      string `gorm:"not null"`
	GeneralInformation string `gorm:"not null"`
}

