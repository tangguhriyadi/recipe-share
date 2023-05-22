package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/recipe-share/controllers/usercontroller"
	"github.com/tangguhriyadi/recipe-share/infrastructure"
	"github.com/tangguhriyadi/recipe-share/repository"
)

func main() {
	infrastructure.ConnectDb()
	userRepo := repository.NewUserRepository(infrastructure.DB)
	userController := usercontroller.NewUserController(&userRepo)

	app := fiber.New()

	v1 := app.Group("/v1")

	v1.Get("/users", userController.GetAll)
	// v1.Get("/users/:id", controllers.FindById)
	// v1.Patch("/users/:id", controllers.Patch)
	// v1.Delete("/users/:id", controllers.Delete)

	app.Listen(":8000")
}
