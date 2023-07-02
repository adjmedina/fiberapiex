package main

import (
	"adjmedina/fiberapiex/config"
	"adjmedina/fiberapiex/database"
	"adjmedina/fiberapiex/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	//Abrimos nuestro archivo logfile
	config.CreateLogFile()

	//Generamos o Enlazamos nuestro archivo SQLite

	database.InitDatabase()
	defer database.CloseDatabase()
	// Creamos una instancia de nuestro framework
	app := fiber.New()

	// Middleware para servir archivos estáticos
	app.Static("/", "./public")

	// Inicializamos el enrutador a partir de  rutas definidas en routes.go
	routes.Setup(app)

	// El server se inicia
	app.Listen(":3000")
}
