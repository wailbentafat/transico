package repository

import (
	"databaseservice/internal/models"
	"databaseservice/pkj/utils"

	"gorm.io/gorm"
)



type CmpinfoRepo interface {}
type CmpinfoRepoImpl struct {
	logger*utils.MyLogger
	db *gorm.DB
}


func NewCmpinfoRepoImpl(db *gorm.DB,log	*utils.MyLogger) *CmpinfoRepoImpl {
	return &CmpinfoRepoImpl{
		logger: log,
		db: db,
	}
}
func (r*CmpinfoRepoImpl)CreateCmpinfo(Cmpinfo *models.CmpInfo		)(error,*models.CmpInfo){
 
	if err:= r.db.Create(Cmpinfo).Error; err != nil {
		r.logger.Error().Msg(err.Error())
		return err,nil
	}
	return nil,Cmpinfo
}



func (r*CmpinfoRepoImpl)GetCmpinfo(ID string )(error,*models.CmpInfo){
	var cmpinfo models.CmpInfo
	result := r.db.Where("ID = ?", ID).First(&cmpinfo)
	if result.Error == nil {
		return result.Error,nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error,nil
	}
	return nil,&cmpinfo
}
func (r *CmpinfoRepoImpl)GetCmpinfobycompanyID(ID uint )(error,*models.CmpInfo){
	var cmpinfo models.CmpInfo
	result := r.db.Where("CompanyID = ?", ID).First(&cmpinfo)
	if result.Error == nil {
		return result.Error,nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error,nil
	}
	return nil,&cmpinfo
}

