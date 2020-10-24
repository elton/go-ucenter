package routes

import (
	"ucenter/app/controllers/sms"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()
	router.GET("/api/v1/sms/send", sms.Send)

	return router
}
