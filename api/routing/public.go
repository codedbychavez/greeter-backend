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

	// A status route to check that the router is OK
	router.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	router.Post("/hello", controllerManager.HelloController.ReturnGreeting)
}