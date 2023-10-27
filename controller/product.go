package controller

import (
	"example/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductController(db *gorm.DB) Product {
	return Product{DB: db}
}

func (controller Product) FindAll(c *gin.Context) {
	// contoh untuk create order
	// product_id, customer_id (token)

	custVal, _ := c.Get("loggedInCustomer")
	loggedInCustomer := custVal.(entity.Customer)
	fmt.Println(loggedInCustomer, "<--- ini data cust login dari controller")

	c.JSON(http.StatusOK, gin.H{
		"products": "ceritanya ini list product",
	})
}
