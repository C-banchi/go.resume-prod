package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.resume-prod/routes"
)

func main() {

	engine := html.New("./view", ".html")

	// pass the engine into the views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// app config
	//app.Static("/", "./public")

	// gets routes
	routes.Routes(app)

	app.Listen(":3000")
}
