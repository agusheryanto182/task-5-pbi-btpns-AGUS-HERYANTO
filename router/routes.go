package router

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/controllers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/helpers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/middlewares"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController  controllers.UserController
	PhotoController controllers.PhotoController
}

type AuthMiddleware struct {
	AuthService helpers.AuthService
	UserService services.UserService
}

func NewRouter(c *Controllers, a *AuthMiddleware) *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", c.UserController.Register)
		userRoute.POST("/login", c.UserController.Login)
		userRoute.PUT("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Update)
		userRoute.DELETE("/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Delete)
	}

	photoRoute := api.Group("/photos")
	{
		photoRoute.POST("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Create)
		photoRoute.GET("/", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.GetByUserID)
		photoRoute.PUT("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Edit)
		photoRoute.DELETE("/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Delete)
	}

	return router
}
