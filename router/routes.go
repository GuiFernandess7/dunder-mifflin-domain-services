package router

import (
	"github.com/GuiFernandess7/db_with_sqlc/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/health-check", )
		v1.GET("/employees", handler.ListEmployeesHandler)
	}
}