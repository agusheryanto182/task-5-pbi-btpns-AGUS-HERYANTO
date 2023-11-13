package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/helpers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/gin-gonic/gin"
)

type PhotoControllerImpl struct {
	photoService services.PhotoService
	authService  helpers.AuthService
}

func (h *PhotoControllerImpl) Create(c *gin.Context) {
	var input app.PhotoInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Upload foto is failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.APIResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(models.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/avatars/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.APIResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.PhotoURL = path

	checkPhoto, _ := h.photoService.GetByUserID(currentUser.ID)
	if checkPhoto.ID != 0 {
		err = h.photoService.Delete(checkPhoto.ID)
		if err != nil {
			response := helpers.APIResponse("Delete photo is failed", http.StatusBadRequest, "failed", err)

			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	_, err = h.photoService.Create(userID, input)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.APIResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{
		"is_uploaded": true,
	}
	response := helpers.APIResponse("Successfully to upload avatar", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *PhotoControllerImpl) GetByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)

	photoDetail, err := h.photoService.GetByUserID(currentUser.ID)
	if err != nil {
		response := helpers.APIResponse("Failed to get detail photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Detail photo", http.StatusOK, "success", photoDetail)
	c.JSON(http.StatusOK, response)
}

func (h *PhotoControllerImpl) Edit(c *gin.Context) {
	var inputData app.PhotoUpdate
	photoID, _ := strconv.Atoi(c.Param("photoId"))

	photoDetail, err := h.photoService.GetByID(photoID)
	if err != nil {
		response := helpers.APIResponse("Failed to get detail photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(models.User)

	if photoDetail.UserID != currentUser.ID {
		response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.photoService.Update(photoID, inputData)
	if err != nil {
		response := helpers.APIResponse("error on update photo", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("Successfully updated user", http.StatusOK, "Sukses", updatedUser)
	c.JSON(http.StatusOK, response)
}

func (h *PhotoControllerImpl) Delete(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)
	ID := currentUser.ID
	photoID, _ := strconv.Atoi(c.Param("id"))

	checkPhoto, err := h.photoService.GetByID(photoID)
	if err != nil {
		response := helpers.APIResponse("Failed get detail photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if checkPhoto.UserID != ID {
		response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = h.photoService.Delete(photoID)
	if err != nil {
		response := helpers.APIResponse("Failed to delete photo", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Successfully, photo is deleted", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func NewPhotoController(photoService services.PhotoService, authService helpers.AuthService) PhotoController {
	return &PhotoControllerImpl{photoService: photoService, authService: authService}
}
