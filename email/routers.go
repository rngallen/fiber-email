package email

import "github.com/gofiber/fiber/v2"

func EmailSetupRouters(router fiber.Router) {

	mail := router.Group("/emails")
	mail.Post("/", SendEmails)

}
