package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zanovru/gin-rest-api/internal/api/resources"
	"github.com/zanovru/gin-rest-api/internal/models"
	"net/http"
	"regexp"
	"time"
)

const ctxKeyElapsed = "elapsed"

type LoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Routing) auth(c *gin.Context) {
	start := time.Now()

	var newLoginRequest LoginRequest

	if err := c.ShouldBindJSON(&newLoginRequest); err != nil {
		errorResponse(c, http.StatusBadRequest, ErrInvalidReqBody.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(newLoginRequest); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userResource := resources.NewUserResource(1)
	c.JSON(http.StatusOK, userResource)

	elapsed := time.Since(start).Milliseconds()
	c.Set(ctxKeyElapsed, elapsed)

}

type RegisterRequest struct {
	Login                 string `json:"login" validate:"required,min=3,alphaNum"`
	Password              string `json:"password" validate:"required,min=6"`
	Password_confirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Name                  string `json:"name,omitempty" validate:"omitempty"`
}

func IsAlphaNum(fl validator.FieldLevel) bool {
	alphaNumericRegexString := "^[a-zA-Z0-9]+$"
	alphaNumericRegex := regexp.MustCompile(alphaNumericRegexString)
	return alphaNumericRegex.MatchString(fl.Field().String())
}

func (r *Routing) register(c *gin.Context) {
	start := time.Now()

	var newRegisterRequest RegisterRequest

	if err := c.ShouldBindJSON(&newRegisterRequest); err != nil {
		errorResponse(c, http.StatusBadRequest, ErrInvalidReqBody.Error())
		return
	}

	validate := validator.New()
	validate.RegisterValidation("alphaNum", IsAlphaNum)

	if err := validate.Struct(newRegisterRequest); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var user models.User
	user.Login = newRegisterRequest.Login
	user.Name = newRegisterRequest.Name
	user.PasswordHash = newRegisterRequest.Password

	id, err := r.services.RegisterUser(user)
	if err != nil {
		errorResponse(c, http.StatusConflict, err.Error())
		return
	}

	userResource := resources.NewUserResource(id)

	c.JSON(http.StatusCreated, userResource)

	elapsed := time.Since(start).Milliseconds()
	c.Set(ctxKeyElapsed, elapsed)

}

func errorResponse(ctx *gin.Context, status int, errorMessage string, errorsData ...any) {
	errorResource := resources.NewErrorResource(errorMessage, errorsData)
	ctx.JSON(status, errorResource)
}
