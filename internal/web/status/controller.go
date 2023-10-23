package status

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/web"
	"net/http"
)

var (
	_ web.Controller = (*Controller)(nil)
)

type Controller struct{}

func NewStatusController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	app.Get("/api/v1/status/", ctrl.Status)
}

func (ctrl *Controller) Status(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Code:    http.StatusOK,
		Message: "Success",
	})
}
