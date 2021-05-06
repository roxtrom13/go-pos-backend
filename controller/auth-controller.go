package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roier/POS/dto"
	"github.com/roier/POS/helper"
	"github.com/roier/POS/models"
	"github.com/roier/POS/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.Login
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Ok!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Invalid credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDto dto.Register
	errDTO := ctx.ShouldBind(&registerDto)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to proces request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicatedEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicated email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		user := c.authService.CreateUser(registerDto)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(user.ID), 10))
		user.Token = token
		response := helper.BuildResponse(true, "Ok!", user)
		ctx.JSON(http.StatusCreated, response)
	}
}
