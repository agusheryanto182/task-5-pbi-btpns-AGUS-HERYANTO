package main

import (
	"fmt"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/controllers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/database/connection"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/helpers"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/repositories"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/router"
	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/services"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := connection.NewDB()
	validate := validator.New()

	authService := helpers.NewAuthService()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, validate)

	photoRepository := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository, validate)

	controller := &router.Controllers{
		UserController:  controllers.NewUserController(userService, authService),
		PhotoController: controllers.NewPhotoController(photoService, authService),
	}

	middleware := &router.AuthMiddleware{
		AuthService: authService,
		UserService: userService,
	}

	r := router.NewRouter(controller, middleware)

	err := r.Run()
	if err != nil {
		fmt.Println("Error on the route run")
	}
}
