package controllers

import "github.com/gofiber/fiber/v2"

func Main(c *fiber.Ctx) error {
	return c.SendString("BASIC: Hello, World 👋!")
}

func Index(c *fiber.Ctx) error {
	rutaArchivo := "public/form.html"
	return c.SendFile(rutaArchivo)
}
