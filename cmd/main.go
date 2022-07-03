package main

import (
	"gofiber-demo/modules/http_auth"
	"gofiber-demo/modules/web_app"
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"
)

func main() {
	app := app_plugin.Register()

	ENV := env_plugin.Register()

	server := http_server_plugin.Register(app)
	session_engine := http_server_plugin.RegisterSessionEngine("sqlite3")

	http_auth.Register(ENV, server, session_engine, []string{"google"})
	web_app.Register(server, session_engine)

	app.Start()
	app.WaitForShutdown()
}
