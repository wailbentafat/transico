package di

import (
	"databaseservice/internal/config"
	"databaseservice/internal/db"
	"databaseservice/internal/handler"
	"databaseservice/internal/middelwear"
	"databaseservice/internal/repository"
	"databaseservice/internal/service"
	"databaseservice/pkj/redis"
	"databaseservice/pkj/utils"
	"fmt"
	
)


type Container struct{
	Services  *Service 
	Repository *Repository
	Handler *Handler
	Config *Config
	Database *Database
	Middeleware *Middleware
	redisconf *redis.Redisconf
	Logger *utils.MyLogger

}
type (
	Service struct{
		Authservice *service.Authserviceimpl
		CmplainService *service.CmplainServiceImpl
		CmpinfoService *service.CmpinfoHandlerImpl
		CmpdataService *service.CmpdataServiceImpl
	}
	Repository struct{
		Authrepository *repository.AuthRepositoryimpl
		Cmplainrepository *repository.CmplinReporImpl
		Cmpinforepository *repository.CmpinfoRepoImpl
		Cmpdatarepository *repository.CmpdataRepositoryImpl
		
	}
	Handler struct{
		Authhandler *handler.AuthHandlerImpl
		Complainhandler *handler.CmplainHandlerImpl
		Cmpinfohandler *handler.CmpinfoHandlerImpl
		Cmpdatahandler *handler.CmpHandler
	}
	Config struct{
		Config *config.Config
	}
	Database struct{
		Database *db.DB
	}
	Middleware struct{
		AdminMiddleware *middelwear.AdminMiddleware
		LoggerMiddleware *middelwear.LoggerMiddleware
		JwtMiddleware *middelwear.JwtMiddleware

		
	}
	redisconf struct{
		redisserver *redisconf
	}
	

)
func Newcontainer(log *utils.MyLogger) (*Container,error) {
	var container Container
	container.Logger=log
	container.Logger.LogInfo().Msg("Newcontainer")
	config,err:= config.LoadConfig()
	if err != nil {
		log.LogError().Msg(err.Error())
		return nil,err
	}
	container.Config = &Config{Config: config}

	 err = container.Initdatabase(config.Db_name)
	if err != nil {
		log.LogError().Msg(err.Error())
		return nil,err
	}
	

container.initredis()
	container.initRepositories()
	container.initServices()
	container.initHandlers()
	container.initMiddlewares()
	
	return &container,nil
}
func (Container *Container) Initdatabase(db_name string) error {
fmt.Println("heyy database is here ")
	dbInstance, err := db.Initdb(db_name)
	if err != nil {
		Container.Logger.LogError().Msg(err.Error())
		return err
	}
	Container.Database=&Database{
		Database: &dbInstance,
	}
	return nil
}
func (container *Container) initRepositories() {
	fmt.Println("heyy repository is here ")
 container.Repository=&Repository{
 	Authrepository: repository.NewAuthRepositoryimpl(
		container.Database.Database.DB,
		container.Logger,
		
	),
	Cmplainrepository: repository.NewCmplinReporImpl( 
	
		container.Database.Database.DB,
		container.Logger,
	),
	Cmpinforepository: repository.NewCmpinfoRepoImpl(
		container.Database.Database.DB,
		container.Logger,

	),
	Cmpdatarepository: repository.NewCmpdataRepositoryImpl(
	
		container.Database.Database.DB,
		container.Logger,
	),
	
   }
	
}

func (container *Container) initServices() {
	fmt.Println("heyy init service is here ")
   container.Services=&Service{
 	Authservice: service.NewAuthservice(
		container.Repository.Authrepository,
	),
	CmplainService: service.NewComplainserviceimpl(
		container.Repository.Cmplainrepository,
	),
	CmpinfoService: service.NewCmpinfoHandlerImpl(
		container.Repository.Cmpinforepository,
	),
	CmpdataService: service.NewCmpdataServiceImpl(
		container.Repository.Cmpdatarepository,
	),

    }
   
}

func (container *Container) initHandlers() {
  container.Handler=&Handler{
	Authhandler: handler.Newauthhandler(container.Services.Authservice),
	Complainhandler: handler.Newcomplainhandler(container.Services.CmplainService),
	Cmpinfohandler: handler.NewcmpinfoHandler(container.Services.CmpinfoService),
	Cmpdatahandler: handler.Newcmpdatahandler(container.Services.CmpdataService),
  }
 
}

	


func (container *Container) initMiddlewares() {
 container.Middeleware=&Middleware{
	AdminMiddleware: middelwear.NewAdminmiddelweare(),
	LoggerMiddleware: middelwear.Newloggermid( container.Logger),
	JwtMiddleware: middelwear.Newjwtmi(),
 }
}
func (container *Container) initredis() {
	config := container.Config.Config


	redisc, err := redis.NewRedisconfig(config.REDIS_ADDR, config.Channelname, container.Logger)
	if err != nil {
		container.Logger.LogError().Msg(err.Error())
		return
	}

	
	container.redisconf = redisc

	container.Logger.LogInfo().Msg("Initialized Redis with address: " + config.REDIS_ADDR)
}



