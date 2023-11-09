package controllers

import (
	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}
