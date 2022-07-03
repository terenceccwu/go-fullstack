package http_server_plugin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type ViewEngine interface {
	Render(c *fiber.Ctx, template string, binding interface{}, layout ...string) error
}

type BaseViewEngine struct {
	Views fiber.Views
}

func NewBaseViewEngine(directory string, extension string) ViewEngine {
	views := html.New(directory, extension)
	views.Load()

	return &BaseViewEngine{
		Views: views,
	}
}

func (e *BaseViewEngine) Render(c *fiber.Ctx, template string, binding interface{}, layout ...string) error {
	c.Set("Content-Type", fiber.MIMETextHTMLCharsetUTF8)
	return e.Views.Render(c, template, binding, layout...)
}
