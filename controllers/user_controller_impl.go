package controllers

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/app"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/middlewares"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userService services.UserService
	authService middlewares.AuthService
}

func (h *UserControllerImpl) Register(c *gin.Context) {
	var input app.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := app.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := app.APIResponse("Register account failed 1", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input.Email)
	if err != nil {
		response := app.APIResponse("Register account failed 2", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isEmailAvailable {
		response := app.APIResponse("Email has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isUsernameAvailable, err := h.userService.IsUsernameAvailable(input.Username)
	if err != nil {
		response := app.APIResponse("Register account failed 3", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isUsernameAvailable {
		response := app.APIResponse("Username has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := app.APIResponse("Register account failed 4", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := app.APIResponse("Account registration faile 5", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := app.RegisterFormatUser(newUser, token)

	response := app.APIResponse("Account successfully registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Login(c *gin.Context) {
	var input app.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := app.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := app.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := app.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := app.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := app.RegisterFormatUser(loggedinUser, token)

	response := app.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Update(c *gin.Context) {
	var inputData app.FormUpdateUserInput
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response := app.APIResponse("Invalid get user ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(models.User)

	inputData.User = currentUser

	if userID != currentUser.ID {
		response := app.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := app.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := app.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.IsUsernameAvailable(inputData.Username)
	if err != nil {
		response := app.APIResponse("username has been  used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.IsEmailAvailable(inputData.Email)
	if err != nil {
		response := app.APIResponse("email has been used", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUser(userID, inputData)
	if err != nil {
		response := app.APIResponse("error on update user", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := app.APIResponse("Successfully updated user", http.StatusOK, "Sukses", app.GetFormatUser(updatedUser))
	c.JSON(http.StatusOK, response)
}

func (h *UserControllerImpl) Delete(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)
	ID := currentUser.ID

	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response := app.APIResponse("Invalid get photo ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if currentUser.ID != userID {
		response := app.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err = h.userService.DeleteUser(ID)
	if err != nil {
		response := app.APIResponse("Failed to delete user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := app.APIResponse("Successfully deleted user", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func NewUserController(userService services.UserService, authService middlewares.AuthService) UserController {
	return &UserControllerImpl{userService: userService, authService: authService}
}
