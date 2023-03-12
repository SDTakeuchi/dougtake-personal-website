package handler

import (
	"blog_app/usecase"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	AuthHandler interface {
		Signup(c *gin.Context)
		Login(c *gin.Context)
		RenewToken(c *gin.Context)
	}
	authHandler struct {
		loginUsecase      usecase.Login
		signupUsecase     usecase.Signup
		renewTokenUsecase usecase.RenewToken
	}

	loginRequest struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	loginResponse struct {
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiresAt  time.Time `json:"access_token_exp"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_exp"`
		UserID                string    `json:"user_id"`
	}

	signupRequest struct {
		Name     string `form:"name"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	signupResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	renewTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
	renewTokenResponse struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_exp"`
	}
)

func NewAuthHandler(
	loginUsecase usecase.Login,
	signupUsecase usecase.Signup,
	renewTokenUsecase usecase.RenewToken,
) AuthHandler {
	return &authHandler{
		loginUsecase:      loginUsecase,
		signupUsecase:     signupUsecase,
		renewTokenUsecase: renewTokenUsecase,
	}
}

func (h *authHandler) Login(c *gin.Context) {
	params := loginRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.loginUsecase.Execute(
		c,
		usecase.LoginInput{
			Email:       params.Email,
			RawPassword: params.Password,
			ClientIP:    c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	resp := loginResponse{
		AccessToken:           output.AccessToken,
		AccessTokenExpiresAt:  output.AccessTokenExpiresAt,
		RefreshToken:          output.RefreshToken,
		RefreshTokenExpiresAt: output.RefreshTokenExpiresAt,
		UserID:                output.UserID.String(),
	}
	createJSONResponse(c, resp)
}

func (h *authHandler) Signup(c *gin.Context) {
	params := signupRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	output, err := h.signupUsecase.Execute(
		c,
		usecase.SignupInput{
			Name:        params.Name,
			Email:       params.Email,
			RawPassword: params.Password,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	resp := signupResponse{
		output.ID.String(),
		output.Name,
		output.Email,
	}
	createJSONResponse(c, resp)
}

func (h *authHandler) RenewToken(c *gin.Context) {
	params := renewTokenRequest{}
	if err := c.Bind(&params); err != nil {
		createErrResponse(c, errFailedToBindQuery)
		return
	}
	fmt.Println("params.RefreshToken")
	fmt.Println(params.RefreshToken)
	output, err := h.renewTokenUsecase.Execute(
		c,
		usecase.RenewTokenInput{
			Token: params.RefreshToken,
		},
	)
	if err != nil {
		createErrResponse(c, err)
		return
	}
	resp := renewTokenResponse{
		output.AccessToken,
		output.AccessTokenExpiresAt,
	}
	createJSONResponse(c, resp)
}
