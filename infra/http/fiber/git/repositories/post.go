package repositories

import (
	entity "belobetty-queue-manager/domain/git"
	fiber2 "belobetty-queue-manager/infra/http/fiber"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func createRepository(c *fiber.Ctx) error {
	var repo = new(entity.RepositoryGit)
	err := c.BodyParser(repo)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	ct, err := fiber2.NewGitController()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = ct.Exec(repo, "create", "repository", c.Get("company"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado criação do usuario %s.", repo.Name))
}
