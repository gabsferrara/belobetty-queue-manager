package users

import (
	entity "belobetty-queue-manager/domain/git"
	"belobetty-queue-manager/infra/http/fiber/git"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func createUser(c *fiber.Ctx) error {
	var user = new(entity.UserGit)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	ct, err := git.NewGitController()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = ct.Exec(user, "create", "users", c.Get("user"))

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado criação do usuario %s.", user.UserName))
}
