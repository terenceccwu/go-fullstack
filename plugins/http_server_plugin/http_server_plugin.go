package http_server_plugin

import (
	"fmt"
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Register(app *app_plugin.Application) *fiber.App {
	server := fiber.New()

	server.Use(logger.New())

	server.Get("/health_check", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.OnStart(func() {
		go server.Listen(fmt.Sprintf(":%d", env_plugin.GetEnvInt("PORT", "3000")))
	})

	app.OnShutdown(func() {
		server.Shutdown()
	})

	return server
}
