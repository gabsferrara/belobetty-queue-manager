package repositories

import (
	entity "belobetty-queue-manager/domain/git"
	fiber2 "belobetty-queue-manager/infra/http/fiber"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func deleteRepository(c *fiber.Ctx) error {
	var repo = new(entity.RepositoryGit)
	repo.Name = c.Params("id")

	ct, err := fiber2.NewGitController()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = ct.Exec(repo, "delete", "repository", c.Get("company"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado deleção do repositorio %s.", repo.Name))
}
