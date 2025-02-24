package models

import "gorm.io/gorm"

type CmpInfo struct {
	gorm.Model
	CarbonWhereToExpected string `gorm:"not null"`
	CarbonWhereToOther string
	CarbonWhereTo string `gorm:"not null"`
	
}

