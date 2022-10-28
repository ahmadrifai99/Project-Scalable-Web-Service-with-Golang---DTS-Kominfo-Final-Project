package routes

import (
	"final-project/configuration"
	"final-project/handler"
	"final-project/infrastructure/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesSosmed(fiber *fiber.App, conf configuration.Config) {
	db := configuration.InitDatabase(conf)

	repo := repository.NewSosmedRepository(db)
	svc := service.NewSosmedService(repo, conf)
	ctrl := handler.NewSosmedController(svc)

	app := fiber.Group("/socialmedias")

	app.Post("/", middleware.JwtMiddleware(), ctrl.CreateSosmed)
	app.Get("/", middleware.JwtMiddleware(), ctrl.GetSosmed)
	app.Put("/:socialMediaId", middleware.JwtMiddleware(), ctrl.UpdateSosmed)
	app.Delete("/:socialMediaId", middleware.JwtMiddleware(), ctrl.DeleteSosmed)
}
