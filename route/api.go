package route

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/voideus/mini-ecommerce/handler"
	"github.com/voideus/mini-ecommerce/middleware"
)

func RunAPI(address string) error {
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()
	orderHandler := handler.NewOrderHandler()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	app := gin.Default()

	app.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         time.Minute,
	}))

	app.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeWeNeedIt")
		}
		ctx.Next()
	})

	app.GET("/", func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
				hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
			})
		}

		ctx.String(http.StatusOK, "Welcome to mini ecommerce")
	})

	apiRoutes := app.Group("/api")

	app.GET("/foo", func(ctx *gin.Context) {
		// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
		// in the middleware before, it will be sent through as well
		panic("y tho again")
	})

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

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)
		productRoutes.GET("/:id", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
	{
		orderRoutes.POST("/product/:product/quantity/:quantity", orderHandler.OrderProduct)
	}

	return app.Run(address)
}
