package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/app/dependencies"
	"github.com/urcop/go-fiber-template/internal/app/initializers"
	"github.com/urcop/go-fiber-template/internal/repository"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	repository.NewExampleRepository()

	container := &dependencies.Container{}

	initializers.SetupRoutes(app, container)
}
