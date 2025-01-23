package handler

import (
	"github.com/gin-gonic/gin"
	"newExampleServer/internal/app"
)

type handler struct {
	di *app.DI
}

func InitRoutes(router *gin.Engine, di *app.DI) {
	h := handler{
		di: di,
	}
	inner := router.Group("/api/v1")
	{
		inner.POST("/register", h.postRegister)
	}
}
