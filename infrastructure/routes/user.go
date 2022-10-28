package routes

import (
	"final-project/configuration"
	"final-project/handler"
	"final-project/infrastructure/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesUser(fiber *fiber.App, conf configuration.Config) {
	db := configuration.InitDatabase(conf)

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(conf, repo)
	ctrl := handler.NewUserController(srv)

	app := fiber.Group("/users")

	app.Post("/register", ctrl.Register)
	app.Post("/login", ctrl.Login)
	app.Put("/:userId", middleware.JwtMiddleware(), ctrl.UpdateUser)
	app.Delete("/:userId", middleware.JwtMiddleware(), ctrl.DeleteUser)
}
