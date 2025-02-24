package router

import "github.com/gin-gonic/gin"





type ComplainRapportroute struct {
	complainHandler *handler.CmplainHandlerImpl
}

func NewComplainRapportRoute(complainHandler *handler.CmplainHandlerImpl) *ComplainRapportroute {
	return &ComplainRapportroute{complainHandler: complainHandler}
}

func (c *ComplainRapportroute) Route(route *gin.Engine) {
	complainRoute := route.Group("/complain")
	{
		complainRoute.POST("", c.complainHandler.Createcomplain)
		complainRoute.GET("", c.complainHandler.GetAllcomplain)
		complainRoute.GET("/:id", c.complainHandler.GetOneComplain)
		complainRoute.PUT("/:id", c.complainHandler.Updatecomplain)
		complainRoute.DELETE("/:id", c.complainHandler.Deletecomplain)
		complainRoute.GET("/report", c.complainHandler.GetComplainRapport)
	}
}
