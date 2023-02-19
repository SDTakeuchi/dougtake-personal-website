package handler

import (
	"blog_app/adapter/config"
	"blog_app/adapter/constants"
	"blog_app/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// returns 200 response with provided parameters
func createJSONResponse(c *gin.Context, body any) {
	c.JSON(
		http.StatusOK,
		gin.H{"data": body},
	)
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
	case model.ErrInvalidParams:
		statusCode = http.StatusBadRequest
		msg = constants.InvalidParams
	default:
		statusCode = http.StatusInternalServerError
		msg = constants.DefaultErrorMessage
	}

	isDebug := config.Get().Debug
	if isDebug {
		// logger.Debug(err.Error())
		res = gin.H{"message": msg.String() + ": === DEBUG === :" + err.Error()}
	} else {
		res = gin.H{"message": msg.String()}
	}

	c.JSON(statusCode, res)
}
