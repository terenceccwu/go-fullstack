package http_server_plugin

import (
	"fmt"
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Register(app *app_plugin.Application) *fiber.App {
	// timeout settings rationale
	// if timeout is not set, server might not able to gracefully shutdown
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	// https://blog.cloudflare.com/exposing-go-on-the-internet/
	// https://nodejs.org/api/http.html#serverkeepalivetimeout
	server := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
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
