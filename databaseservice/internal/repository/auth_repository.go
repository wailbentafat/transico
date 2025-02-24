package repository

import (
	"databaseservice/internal/models"

	"gorm.io/gorm"
	"databaseservice/pkj/utils"
	
	
)


type Authrepository interface{
	
	
}
type AuthRepositoryimpl struct {
	logger *utils.MyLogger
	db *gorm.DB

}
func NewAuthRepositoryimpl(db *gorm.DB,log *utils.MyLogger) *AuthRepositoryimpl {
	return &AuthRepositoryimpl{
		logger: log,
		db: db,
	}


}
func (r*AuthRepositoryimpl)Regiter(email string,password string )error{
	
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	user.Email = email
	user.Password = password
	result = r.db.Create(&user)
	if result.Error != nil {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil



}
func (r*AuthRepositoryimpl)Login(email string,password string )error{
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	if user.Password != password {
		return result.Error
	}
	return nil
}
func (r*AuthRepositoryimpl)UpdatePassword(email string,password string )error{
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	user.Password = password
	result = r.db.Save(&user)
	if result.Error != nil {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil
}
func (r*AuthRepositoryimpl)DeleteUser(email string )error{
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return result.Error
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	result = r.db.Delete(&user)
	if result.Error != nil {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error
	}
	return nil
}
func (r*AuthRepositoryimpl)GetUser(ID string )(error,*models.User){
	var user models.User
	result := r.db.Where("ID = ?", ID).First(&user)
	if result.Error == nil {
		return result.Error,nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error,nil
	}
	return nil,&user
}
func (r*AuthRepositoryimpl)GetUserByEmail(email string )(error,*models.User){
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error == nil {
		return result.Error,nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		r.logger.Error().Msg(result.Error.Error())
		return result.Error,nil	
	}
	

	return nil,&user
}

