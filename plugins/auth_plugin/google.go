package auth_plugin

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

func RegisterGoogleAuth(ENV map[string]string, server *fiber.App, Session func(c *fiber.Ctx) error) {
	goth.UseProviders(
		google.New(ENV["GOOGLE_CLIENT_ID"], ENV["GOOGLE_CLIENT_SECRET"], ENV["GOOGLE_CALLBACK_URL"]),
	)

	server.Get("/auth/me", Session, func(c *fiber.Ctx) error {
		session := c.Locals("session").(*session.Session)
		user := session.Get("user")
		return c.JSON(fiber.Map{
			"user": user,
		})
	})

	server.Get("/auth/:provider", goth_fiber.BeginAuthHandler)
	server.Get("/auth/:provider/callback", Session, func(c *fiber.Ctx) error {
		user, err := goth_fiber.CompleteUserAuth(c)
		if err != nil {
			log.Fatal(err)
		}
		session := c.Locals("session").(*session.Session)
		session.Set("user", user)
		session.Save()

		return c.Redirect("/auth/me")
	})
}
