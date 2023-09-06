package usermodule

import "github.com/gofiber/fiber/v2"

// apply routes in app fiber, with controllers and services defined
func UserModuleDecorator(app *fiber.App) {
	s := newService()
	c := newController(s)
	newRoutes(c, app)
}
