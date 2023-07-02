package main

import (
	"adjmedina/fiberapiex/config"
	"adjmedina/fiberapiex/database"
	"adjmedina/fiberapiex/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	//Abrimos nuestrologfile
	config.CreateLogFile()

	//Genramos o Enlazamos nuestro archivo sqlite

	database.InitDatabase()
	defer database.CloseDatabase()
	// Creamos una instancia de nuestro framework
	app := fiber.New()

	// Usamos las rutas definidas en routes.go
	routes.Setup(app)
	app.Listen(":3000")
}
