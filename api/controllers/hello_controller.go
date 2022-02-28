package controllers

import (

	"github.com/gofiber/fiber/v2"
)

type HelloController struct {}

type HelloControllerManager interface {
	ReturnGreeting(c *fiber.Ctx) error
}

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (ctrl HelloController) ReturnGreeting(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	personName := p.Name

	greeting := "Hello, " + personName + "!"

	return c.JSON(greeting)
}