package email

import "github.com/gofiber/fiber/v2"

func SendEmails(c *fiber.Ctx) error {

	emails := new([]UserList)

	err := c.BodyParser(emails)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err})
	}

	er := SendNewsLetter(*emails)

	if er != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": er})
	}
	return c.Status(200).JSON(fiber.Map{"msg": "email(s) sent successfully"})

}
