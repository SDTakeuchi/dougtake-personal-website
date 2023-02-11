package handler

import (
	"blog_app/adapter/config"
	"blog_app/adapter/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createJSONResponse(c *gin.Context, status int, body any) {
	if http.StatusText(status) == "" {
		panic(fmt.Errorf("unknown status code: %d", status))
	}
	c.JSON(status, gin.H{"data": body})
}

func createErrResponse(c *gin.Context, err error) {
	var (
		statusCode int
		msg        constants.ResponseMessage
		res        gin.H
	)

	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = http.StatusNotFound
		msg = constants.RecordNotFound
	case errFailedToBindQuery:
		statusCode = http.StatusBadRequest
		msg = constants.FailedToBindQuery
	default:
		statusCode = http.StatusInternalServerError
		msg = constants.DefaultErrorMessage
	}

	isDebug := config.Get().Debug
	if isDebug {
		// logger.Debug(err.Error())
		res = gin.H{"message": msg.String() + ":" + err.Error()}
	} else {
		res = gin.H{"message": msg.String()}
	}

	c.JSON(statusCode, res)
}
