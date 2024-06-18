package router

import (
	"github.com/GuiFernandess7/db_with_sqlc/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/health-check", handler.HealthCheck)
		v1.GET("/clients/employee", handler.GetClientByEmployeeHandler)
		v1.GET("/clients/search", handler.GetClientByBranchHandler)
		v1.GET("/employees", handler.ListEmployeesHandler)
		v1.GET("/employees/filter", handler.ListEmployeesFilterHandler)
		v1.GET("/employees/search", handler.SearchEmployeeHandler)
		v1.GET("/employees/sales/search", handler.GetEmployeeBySalesQtd)
		v1.GET("/suppliers/search", handler.GetSupplierBySupplyType)
	}
}