package routes

import (
	"final-project/configuration"
	"final-project/handler"
	"final-project/infrastructure/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gofiber/fiber/v2"
)

func RoutesComment(fiber *fiber.App, conf configuration.Config) {
	db := configuration.InitDatabase(conf)

	repo := repository.NewCommentRepository(db)
	svc := service.NewCommentService(repo, conf)
	ctrl := handler.NewCommentController(svc)

	app := fiber.Group("/comments")

	app.Post("/", middleware.JwtMiddleware(), ctrl.CreateComment)
	app.Get("/", middleware.JwtMiddleware(), ctrl.GetComment)
	app.Put("/:commentId", middleware.JwtMiddleware(), ctrl.UpdateComment)
	app.Delete("/:commentId", middleware.JwtMiddleware(), ctrl.DeleteComment)
}
