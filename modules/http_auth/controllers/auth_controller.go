package controllers

import (
	"gofiber-demo/plugins/http_server_plugin"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

type AuthController struct {
	ENV           map[string]string
	ViewEngine    http_server_plugin.ViewEngine
	SessionEngine http_server_plugin.SessionEngine
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	return a.ViewEngine.Render(c, "login", fiber.Map{
		"ENV": a.ENV,
	})
}

func (a *AuthController) Logout(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		return err
	}

	return c.Redirect("/auth/me")
}

func (a *AuthController) GetMe(c *fiber.Ctx) error {
	session := a.SessionEngine.GetSession(c)
	user := session.Get("user")

	return c.JSON(fiber.Map{
		"user": user,
	})
}

func (a *AuthController) AuthStart(c *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(c)
}

func (a *AuthController) AuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
	}

	// session should be get AFTER auth, as new session id is generated after auth
	// due to session fixation https://medium.com/passportjs/fixing-session-fixation-b2b68619c51d
	session := a.SessionEngine.GetSession(c)
	session.Set("user", user)
	session.Save()

	return c.Redirect("/auth/me")
}
