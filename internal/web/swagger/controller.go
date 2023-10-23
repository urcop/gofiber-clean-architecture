package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/urcop/go-fiber-template/api"
	"github.com/urcop/go-fiber-template/internal/web"
)

var _ web.Controller = (*Controller)(nil)

type Controller struct{}

func NewSwaggerController() *Controller {
	return &Controller{}
}

func (c *Controller) DefineRouter(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
