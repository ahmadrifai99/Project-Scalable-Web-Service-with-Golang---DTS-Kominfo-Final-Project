package routes

import (
	"final-project/configuration"
	"final-project/handler"
	"final-project/infrastructure/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesPhoto(fiber *fiber.App, conf configuration.Config) {
	db := configuration.InitDatabase(conf)

	repo := repository.NewPhotoRepository(db)
	svc := service.NewPhotoService(repo, conf)
	ctrl := handler.NewPhotoController(svc)

	app := fiber.Group("/photos")

	app.Post("/", middleware.JwtMiddleware(), ctrl.CreatePhoto)
	app.Get("/", middleware.JwtMiddleware(), ctrl.GetAllPhoto)
	app.Put("/:photoId", middleware.JwtMiddleware(), ctrl.UpdatePhoto)
	app.Delete("/:photoId", middleware.JwtMiddleware(), ctrl.DeletePhoto)
}
