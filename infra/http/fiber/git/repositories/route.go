package repositories

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {

	repo := r.Group("/repository")

	repo.Post("/", createRepository)
	repo.Put("/", updateRepository)
	repo.Delete("/:id", deleteRepository)

}
