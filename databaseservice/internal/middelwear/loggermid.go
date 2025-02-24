package middelwear

import (
	"databaseservice/pkj/utils"

	"github.com/gin-gonic/gin"
)



type LoggerMiddleware struct {
	logger *utils.MyLogger
}


func Newloggermid(logger  *utils.MyLogger)*LoggerMiddleware{
	return &LoggerMiddleware{
		logger: logger,
	}


}
func (m *LoggerMiddleware) LoggerMiddleware()gin.HandlerFunc  {
	return func(c *gin.Context) {
		m.logger.LogInfo().Msg("Request received")
		m.logger.LogInfo().Msg("Request method: " + c.Request.Method,)
		m.logger.LogInfo().Msg("Request URL: " + c.Request.URL.String())
		m.logger.LogInfo().Msg("Request headers: " + c.Request.Header.Get("Content-Type"))
		c.Next()
		m.logger.LogInfo().Msgf("Response status: %d", c.Writer.Status())

	}
}