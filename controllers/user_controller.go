package controllers

import (
	"adjmedina/fiberapiex/database"
	"adjmedina/fiberapiex/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	log.Printf("Solicitud de Obtención de Usuarios")
	// Conecta con la base de datos
	db := database.DBConn
	// Inicia un array vacio de albums
	var users []models.User
	// Encuentra todos los albums y los guarda en el array
	db.Find(&users)
	// Envia el array como JSON
	return c.JSON(users)

	///fmt.Println("Todos los albums")
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	// Validar que el campo username no esté vacío
	if user.Username == "" || user.Email == "" {
		return c.Status(503).JSON(fiber.Map{"message": "Username is required"})
	}

	//Añade el Registro a la base de datos con GORM
	log.Println("Se creo el usuario ", &user.ID)
	db.Create(&user)
	return c.Status(201).SendString("Created")

}
