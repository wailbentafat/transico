package repository

import (
	"databaseservice/internal/models"
	"databaseservice/pkj/utils"

	"gorm.io/gorm"
)




type CmpdataRepository interface {}
type CmpdataRepositoryImpl struct {
	logger	*utils.MyLogger
	db *gorm.DB
}
func NewCmpdataRepositoryImpl(db *gorm.DB,log	*utils.MyLogger) *CmpdataRepositoryImpl {
	return &CmpdataRepositoryImpl{
		logger: log,
		db: db,
	}
}

 func (r*CmpdataRepositoryImpl)creaeCmpdata(name string, creatorID uint )(error,*models.CmpData){
	 
 
	
   cmpdata:= &models.CmpData{
	Name: name,
	CreatorID: creatorID,
	CmpInfoID: nil,
	}
	if err := r.db.Create(cmpdata).Error; err != nil {
		r.logger.Error().Msg(err.Error())
		return err,nil
	}
	return nil,cmpdata
   }



func (r*CmpdataRepositoryImpl)GetCmpdataByID(ID string )(error,*models.CmpData){
	var cmpdata models.CmpData
	result := r.db.Where("ID = ?", ID).First(&cmpdata)
	if result.Error == nil {
		return result.Error,nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error,nil
	}
	return nil,&cmpdata
}

func (r*AuthRepositoryimpl)DeleteCmpdata(ID string )error{
	var cmpdata models.CmpData
	result := r.db.Where("ID = ?", ID).First(&cmpdata)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	result = r.db.Delete(&cmpdata)
	if result.Error != nil {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil
}

func (r*CmpdataRepositoryImpl)UpdateCmpdataName(ID string,name string )error{
	var cmpdata models.CmpData
	result := r.db.Where("ID = ?", ID).First(&cmpdata)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	cmpdata.Name = name
	result = r.db.Save(&cmpdata)
	if result.Error != nil {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil
}
func (r*CmpdataRepositoryImpl)GetCmpdatabyowner( creatorID uint )error{
	var cmpdata models.CmpData
	result := r.db.Where("CreatorID = ?", creatorID).First(&cmpdata)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil
	}

	func (r*CmpdataRepositoryImpl)GetComplaintsByCreator(creatorID uint) ([]models.Complain, error) {
		var complaints []models.Complain
	
		
		err := r.db.
			Joins("JOIN cmp_data ON cmp_data.id = complains.company_id").
			Where("cmp_data.creator_id = ?", creatorID).
			Find(&complaints).Error
	
		if err != nil {
			return nil, err
		}
	
		return complaints, nil
	}
	

	





   
