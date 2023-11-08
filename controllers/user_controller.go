package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}