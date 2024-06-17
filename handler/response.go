package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EmployeeResponse struct {
    EmpID    int       `json:"emp_id"`
    FirstName string    `json:"first_name"`
    LastName string     `json:"last_name"`
    BirthDay *time.Time `json:"birthday,omitempty"`
    Sex      string     `json:"sex"`
    Salary   int        `json:"salary"`
	Branch   string 	`json:"branch"`
}

func sendSuccess(ctx *gin.Context, op string, data interface{}){
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s sucessfull", op),
		"data": data,
	})
}

func sendError(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"error_code": code,
	})
}