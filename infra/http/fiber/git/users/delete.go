package users

import (
	entity "belobetty-queue-manager/domain/git"
	fiber2 "belobetty-queue-manager/infra/http/fiber"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func deleteUser(c *fiber.Ctx) error {
	var user = new(entity.UserGit)
	user.UserName = c.Params("user")

	ct, err := fiber2.NewGitController()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = ct.Exec(user, "delete", "users", c.Get("company"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado deleção do usuario %s.", user.UserName))
}
