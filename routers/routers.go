package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rngallen/fiber-email/email"
)

func SetupRouters(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	email.EmailSetupRouters(v1)

}
