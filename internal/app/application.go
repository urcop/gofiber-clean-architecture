package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/app/initializers"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	initializers.InitEnv()
	initializers.SetupRoutes(app)
}
