package main

import (
	"fiber-crud/database"
	"fiber-crud/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDB()

	app := fiber.New() // ✅ Yeh v3 ke liye sahi hai
	// 
	app.Route("/routes", func(api fiber.Router) {
		routes.Post("/users", routes.CreateUser)
	      routes.Get("/users", routes.GetUsers)
		api.Put("/users/:id", routes.UpdateUser)
		api.Delete("/users/:id", routes.DeleteUser)
	})
	
	
	app.Listen(":3000")
}
