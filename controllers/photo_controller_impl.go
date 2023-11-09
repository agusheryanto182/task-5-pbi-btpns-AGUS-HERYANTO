package controllers

import (
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/middlewares"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/gin-gonic/gin"
)

type PhotoControllerImpl struct {
	photoService services.PhotoService
	authService  middlewares.AuthService
}

func (h *PhotoControllerImpl) Create(c *gin.Context) {

}

func (h *PhotoControllerImpl) GetByID(c *gin.Context) {

}

func (h *PhotoControllerImpl) Edit(c *gin.Context) {

}

func (h *PhotoControllerImpl) Delete(c *gin.Context) {

}
