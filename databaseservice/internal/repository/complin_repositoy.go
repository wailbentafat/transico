package repository

import (
	"databaseservice/internal/models"
	"databaseservice/pkj/utils"

	"gorm.io/gorm"
)



type CmplinRepo interface {
}
type CmplinReporImpl struct {
	logger	*utils.MyLogger
	db *gorm.DB
}

func NewCmplinReporImpl(db *gorm.DB,log	*utils.MyLogger) *CmplinReporImpl {
	return &CmplinReporImpl{
		logger: log,
		db: db,
	}
}


func (r *CmpinfoRepoImpl)Createcomplain(Companyid *uint, types	string , complaindesc string)(error,*models.Complain){
	var complain models.Complain
	complain.CompanyID=Companyid
	complain.Type=types
	complain.ComplainDesc=complaindesc

	if err:=r.db.Create(&complain).Error;err!=nil{
		r.logger.Error().Msg(err.Error())
		return err,nil



	}
	return nil,&complain
}
func (r *CmpinfoRepoImpl)Gtbyid(ID uint )(*models.Complain,error){
	var complain 	models.Complain
	if err:=r.db.Where("CompanyID=?").First(&complain);err!=nil{
		r.logger.Log().Msg(err.Error.Error())
		return nil,err.Error
	}

	return &complain,nil
}









