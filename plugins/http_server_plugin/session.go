package http_server_plugin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore *session.Store

func RegisterSessionStore() {
	SessionStore = session.New()
}

func Session(c *fiber.Ctx) error {
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	c.Locals("session", session)

	return c.Next()
}
