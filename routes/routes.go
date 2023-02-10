package routes

import (
	"todo-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	//Auth
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)

	//Todo
	app.Post("/api/todos/create", controllers.CreateToDo)
	app.Get("/api/todos/", controllers.GetToDos)
}
