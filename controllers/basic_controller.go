package controllers

import "github.com/gofiber/fiber/v2"

func Main(c *fiber.Ctx) error {
	return c.SendString("BASIC: Hello, World 👋!")
}
