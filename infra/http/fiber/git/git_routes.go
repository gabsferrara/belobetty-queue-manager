package git

import (
	"belobetty-queue-manager/infra/http/fiber/git/repositories"
	"belobetty-queue-manager/infra/http/fiber/git/users"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router) {

	git := r.Group("/github/")

	users.SetRoutes(git)
	repositories.SetRoutes(git)

}
