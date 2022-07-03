package main

import (
	"context"
	"gofiber-demo/plugins/app_plugin"
	"gofiber-demo/plugins/env_plugin"
	"gofiber-demo/plugins/mongo_plugin"
)

var app = &app_plugin.Application{}
var ENV = env_plugin.Register()
var db = mongo_plugin.Register(app, ENV)
var ctx = context.Background()

func init() {
	app.Start()
}
