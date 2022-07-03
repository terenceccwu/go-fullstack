package controllers

import (
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	ViewEngine    http_server_plugin.ViewEngine
	SessionEngine http_server_plugin.SessionEngine
}

func (h *HomeController) Index(c *fiber.Ctx) error {
	return h.ViewEngine.Render(c, "home", fiber.Map{
		"ENV": env_plugin.ENV,
	})
}

func (h *HomeController) Session(c *fiber.Ctx) error {
	session := h.SessionEngine.GetSession(c)

	return c.JSON(fiber.Map{
		"test":       1,
		"session_id": session.ID(),
	})
}
