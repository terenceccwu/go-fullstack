package http_server_plugin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type ViewEngine struct {
	Views fiber.Views
}

func NewViewEngine(directory string, extension string) *ViewEngine {
	views := html.New(directory, extension)
	views.Load()

	return &ViewEngine{
		Views: views,
	}
}

func (e *ViewEngine) Render(c *fiber.Ctx, template string, binding interface{}, layout ...string) error {
	c.Set("Content-Type", fiber.MIMETextHTMLCharsetUTF8)
	return e.Views.Render(c, template, binding, layout...)
}
