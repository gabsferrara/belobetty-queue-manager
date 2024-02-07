package repositories

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func createRepository(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("createUser")
}
