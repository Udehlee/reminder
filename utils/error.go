package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Status     string
	Message    string
	StatusCode int
}

func BadRequestErr(c *gin.Context, msg, status string, statusCode int) {

	Err := ErrResponse{
		Status:     status,
		Message:    msg,
		StatusCode: statusCode,
	}
	c.JSON(statusCode, Err)
}
