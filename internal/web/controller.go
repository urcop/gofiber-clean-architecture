package web

import "github.com/gofiber/fiber/v2"

type Controller interface {
	DefineRouter(app *fiber.App)
}
