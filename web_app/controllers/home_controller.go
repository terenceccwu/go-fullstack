package controllers

import (
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type HomeController struct {
	Views *http_server_plugin.ViewEngine
}

func (h *HomeController) Index(c *fiber.Ctx) error {
	return h.Views.Render(c, "home", fiber.Map{
		"ENV": env_plugin.ENV,
	})
}

func (h *HomeController) Session(c *fiber.Ctx) error {
	session := c.Locals("session").(*session.Session)
	// session.Save()

	return c.JSON(fiber.Map{
		"test":       1,
		"session_id": session.ID(),
	})
}
