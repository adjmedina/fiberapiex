package routes

import (
	"adjmedina/fiberapiex/controllers"

	"github.com/gofiber/fiber/v2"
)

/*
Esta Funcion es la encargada de inyectar a la app las diferentes rutas que el sistema habilitará
*/
func Setup(app *fiber.App) {

	//Ruta Main
	app.Get("/", controllers.Index)
	app.Get("/main", controllers.Main)
	// Rutas para usuarios
	app.Post("/users", controllers.CreateUser)
	app.Get("/users", controllers.GetUser)
	//app.Get("/users/:id", controllers.GetUserById)
	//app.Put("/users/:id", controllers.UpdateUser)
	//app.Delete("/users/:id", controllers.DeleteUser)
}
