package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voideus/mini-ecommerce/handler"
	"github.com/voideus/mini-ecommerce/middleware"
)

func RunAPI(address string) error {
	userHandler := handler.NewUserHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to mini ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/sign-in", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:id", userHandler.GetUser)
		userProtectedRoutes.GET("/:id/products", userHandler.GetProductOrdered)
	}

	return r.Run(address)
}
