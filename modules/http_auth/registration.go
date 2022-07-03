package http_auth

import (
	"fmt"
	"gofiber-demo/modules/http_auth/controllers"
	"gofiber-demo/plugins/http_server_plugin"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

var CWD string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	CWD = filepath.Dir(filename)
}

func Register(ENV map[string]string, server *fiber.App, session_engine http_server_plugin.SessionEngine, providers []string) {
	for _, provider := range providers {
		switch provider {
		case "google":
			goth.UseProviders(
				google.New(ENV["GOOGLE_CLIENT_ID"], ENV["GOOGLE_CLIENT_SECRET"], ENV["GOOGLE_CALLBACK_URL"]),
			)
		}
	}

	goth_fiber.SessionStore = session_engine.GetSessionStore()

	view_engine := RegisterViewEngine()
	auth := &controllers.AuthController{ENV: ENV, ViewEngine: view_engine, SessionEngine: session_engine}

	server.Static("/auth/static/", fmt.Sprintf("%s/static", CWD))
	server.Get("/auth/login", auth.Login)
	server.Get("/auth/logout", auth.Logout)
	server.Get("/auth/me", auth.GetMe)
	server.Get("/auth/:provider", auth.AuthStart)
	server.Get("/auth/:provider/callback", auth.AuthCallback)
}

func RegisterViewEngine() http_server_plugin.ViewEngine {
	views_directory := fmt.Sprintf("%s/views", CWD)
	view_engine := http_server_plugin.NewBaseViewEngine(views_directory, ".html")
	return view_engine
}
