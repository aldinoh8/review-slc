package main

import (
	"example/config"
	"example/controller"
	"example/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDb()
	app := gin.Default()

	userRoutes := app.Group("/users")
	userController := controller.NewUserController(db)
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
	}

	authMiddleware := middleware.NewAuthMiddleware(db)

	productRoutes := app.Group("/products")
	productController := controller.NewProductController(db)
	productRoutes.Use(authMiddleware.Authenticate())
	{
		productRoutes.GET("", productController.FindAll)
	}

	app.Run()
}
