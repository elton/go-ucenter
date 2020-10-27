package routes

import (
	"ucenter/app/controllers/sms"
	"ucenter/app/controllers/user"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()
	router.GET("/api/v1/sms/send", sms.Send)
	router.POST("/api/v1/user/reg", user.Reg)

	return router
}
