package users

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {

	user := r.Group("/user")

	user.Post("/", createUser)
	user.Put("/", updateUser)
	user.Delete("/:user", deleteUser)

}
