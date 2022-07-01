package main

import (
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/auth_plugin"
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"
	"gofiber-demo/web_app"
)

func main() {
	app := app_plugin.Register()

	ENV := env_plugin.Register()

	server := http_server_plugin.Register(app)
	http_server_plugin.RegisterSessionStore()
	auth_plugin.RegisterGoogleAuth(ENV, server, http_server_plugin.Session)

	web_app.Register(server)

	app.Start()
	app.WaitForShutdown()
}
