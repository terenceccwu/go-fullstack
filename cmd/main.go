package main

import (
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/http_server_plugin"
	"gofiber-demo/web_app"
)

func main() {
	app := app_plugin.Register()

	env_plugin.Register()

	server := http_server_plugin.Register(app)
	web_app.Register(server)

	app.Start()
	app.WaitForShutdown()
}
