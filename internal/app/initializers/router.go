package initializers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/web"
	"github.com/urcop/go-fiber-template/internal/web/status"
	"github.com/urcop/go-fiber-template/internal/web/swagger"
)

func SetupRoutes(app *fiber.App) {
	ctrl := buildRouters()

	for i := range ctrl {
		ctrl[i].DefineRouter(app)
	}
}

func buildRouters() []web.Controller {
	return []web.Controller{
		status.NewStatusController(),
		swagger.NewSwaggerController(),
	}
}
