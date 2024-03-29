package handler

import (
	"blog_app/adapter/config"
	"blog_app/adapter/constants"
	"blog_app/domain/model"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/password"
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
	)

	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = http.StatusNotFound
		msg = constants.RecordNotFound
	case errFailedToBindQuery:
		statusCode = http.StatusBadRequest
		msg = constants.FailedToBindQuery

	case password.ErrIncorrectPassword:
		statusCode = http.StatusBadRequest
		msg = constants.IncorrectPassword
	case password.ErrPasswordTooShort:
		statusCode = http.StatusBadRequest
		msg = constants.PasswordTooShort
	case auth.ErrInvalidToken:
		statusCode = http.StatusUnauthorized
		msg = constants.InvalidToken
	case auth.ErrExpiredToken:
		statusCode = http.StatusUnauthorized
		msg = constants.ExpiredToken

	case model.ErrConstructor:
		statusCode = http.StatusBadRequest
		msg = constants.InvalidParams

	default:
		statusCode = http.StatusInternalServerError
		msg = constants.DefaultErrorMessage
	}

	isDebug := config.Get().Debug
	var res gin.H
	if isDebug {
		// logger.Debug(err.Error())
		res = gin.H{"message": msg.String() + ": === DEBUG === :" + err.Error()}
	} else {
		res = gin.H{"message": msg.String()}
	}

	c.JSON(statusCode, res)
}
