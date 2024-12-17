package routes

import (
	"dukebward/search/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// export function so capital letter start
func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return render(c, views.Login())
	})
}

// non exported function
func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}