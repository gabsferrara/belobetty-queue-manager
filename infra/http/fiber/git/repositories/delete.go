package repositories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func deleteRepository(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("deleteUser")
}
