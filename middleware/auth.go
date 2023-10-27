package middleware

import (
	"example/dto"
	"example/entity"
	"example/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) Auth {
	return Auth{DB: db}
}

func (a Auth) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// cek request header nya bawa token atau tidak?
		// error: 401 please provide valid token
		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Please provide valid Auth Token",
			})
			return
		}

		// decode token pake secret -> claims (ada cust id)
		// error: 401
		claims, err := helper.ParseToken(authToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Please provide valid Auth Token",
			})
			return
		}

		// db find cust by id
		// kalo gaada: error 401
		loggedInCustomer := entity.Customer{}
		err = a.DB.First(&loggedInCustomer, claims["id"]).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Please provide valid Auth Token",
			})
			return
		}

		ctx.Set("loggedInCustomer", loggedInCustomer)
		ctx.Next()
	}
}
