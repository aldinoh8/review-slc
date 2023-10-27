package controller

import (
	"example/dto"
	"example/entity"
	"example/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) User {
	return User{DB: db}
}

func (controller User) Register(c *gin.Context) {
	reqBody := dto.CustomerRegisterBody{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	newCustomer := entity.Customer{
		FirstName:   reqBody.FirstName,
		LastName:    reqBody.LastName,
		Email:       reqBody.Email,
		Password:    reqBody.Password,
		PhoneNumber: reqBody.PhoneNumber,
		Address:     reqBody.Address,
	}

	if err := controller.DB.Create(&newCustomer).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success register",
		"account": newCustomer,
	})
}

func (controller User) Login(c *gin.Context) {
	reqBody := dto.CustomerLoginBody{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
		return
	}

	customer := entity.Customer{}
	err := controller.DB.Where("email = ?", reqBody.Email).First(&customer).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid email/password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(reqBody.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Message: "Invalid email/password"})
		return
	}

	token, _ := helper.GenerateToken(jwt.MapClaims{
		"id": customer.ID,
	})

	// generate token jwt
	c.JSON(http.StatusOK, gin.H{
		"message": "success login",
		"token":   token,
	})
}
