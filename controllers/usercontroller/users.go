package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/recipe-share/repository"
)

type UserController struct {
	UserRepo repository.UserRepository
}

func NewUserController(UserRepo *repository.UserRepository) UserController {
	return UserController{
		UserRepo: *UserRepo,
	}
}

func (ur *UserController) GetAll(ctx *fiber.Ctx) error {
	c := ctx.Context()

	result, err := ur.UserRepo.GetAll(c)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
