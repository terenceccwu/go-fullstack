package controllers

import (
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"

	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
	Views *http_server_plugin.ViewEngine
}

func (h *HomeController) Index(c *fiber.Ctx) error {
	return h.Views.Render(c, "home", fiber.Map{
		"AppName": env_plugin.GetEnv("APP_NAME"),
		"ENV":     env_plugin.ENV,
	})
}
