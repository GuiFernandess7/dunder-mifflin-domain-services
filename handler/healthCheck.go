package handler

import "github.com/gin-gonic/gin"

type checkResponse struct {
	Status	string `json:"status"`
}

func HealthCheck(ctx *gin.Context) {
	responseData := checkResponse{
		Status: 	"Working!",
	}
	sendSuccess(ctx, "Sucess", responseData)
}