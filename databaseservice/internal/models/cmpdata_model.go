package models

import "gorm.io/gorm"





type CmpData struct {
	gorm.Model
	Name      string `gorm:"not null"`

	CreatorID uint `gorm:"index"`
	Creator   User `gorm:"foreignKey:CreatorID;references:ID"`

    
	CmpInfoID *uint    `gorm:"uniqueIndex;default:null"`
	CmpInfo   *CmpInfo `gorm:"foreignKey:CmpInfoID;constraint:OnDelete:SET NULL"`


	Complains []Complain `gorm:"foreignKey:CompanyID"`
}
