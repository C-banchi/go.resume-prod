package routes

import (
	"github.com/gofiber/fiber/v2"
)

// pass in app to the setup routes function
func Routes(app *fiber.App) {
	//GET  -------------------------------------------------
	//app.Get("/api/user", controllers.User)
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	//POST -------------------------------------------------
	//app.Post("/api/register", controllers.Register) // create the path in the controllers folder

	// PUT  -------------------------------------------------

	// DELETE -----------------------------------------------
}
