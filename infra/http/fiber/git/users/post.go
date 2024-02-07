package users

import (
	"belobetty-queue-manager/application/usecases"
	"belobetty-queue-manager/domain/git"
	"belobetty-queue-manager/infra/queue"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func createUser(c *fiber.Ctx) error {
	var user = new(git.UserGit)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	q, err := queue.NewRabbitMQ("git")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	uc := usecases.NewSenderUseCase(q)

	err = uc.Exec(user, "create", "users", c.Get("user"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("Solicitado criação do usuario %s.", user.UserName))
}
