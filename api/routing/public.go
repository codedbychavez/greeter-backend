package routing

import (
	"github.com/gofiber/fiber/v2"
	"greeter-backend/api/controllers"
)

func CreatePublicRoutes(router fiber.Router, controllerManager controllers.ControllerManager) {

	router.Use(func(c *fiber.Ctx) error {
		c.Set("Allow", "GET, POST, PUT")
		return c.Next()
	})

	router.Post("/hello", controllerManager.HelloController.ReturnGreeting)
}