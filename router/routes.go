package router

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/controllers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/middlewares"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	UserController  controllers.UserController
	PhotoController controllers.PhotoController
}

type AuthMiddleware struct {
	AuthService middlewares.AuthService
	UserService services.UserService
}

func NewRouter(c *Controllers, a *AuthMiddleware) *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users/register", c.UserController.Register)
	api.POST("/users/login", c.UserController.Login)
	api.PUT("/users/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Update)
	api.DELETE("/users/:userId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.UserController.Delete)

	api.POST("/photos", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Create)
	api.GET("/photos", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.GetByUserID)
	api.PUT("/photos/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Edit)
	api.DELETE("/photos/:photoId", middlewares.AuthMiddleware(a.AuthService, a.UserService), c.PhotoController.Delete)

	return router
}
