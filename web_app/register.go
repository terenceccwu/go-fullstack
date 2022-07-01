package web_app

import (
	"fmt"
	"gofiber-demo/plugins/http_server_plugin"
	"gofiber-demo/web_app/controllers"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

var CWD string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	CWD = filepath.Dir(filename)
}

func Register(server *fiber.App) {
	views := RegisterViewEngine()

	home := &controllers.HomeController{Views: views}

	server.Static("/web_app/static/", fmt.Sprintf("%s/static", CWD))
	server.Get("/", home.Index)
}

func RegisterViewEngine() *http_server_plugin.ViewEngine {
	views_directory := fmt.Sprintf("%s/views", CWD)
	views := http_server_plugin.NewViewEngine(views_directory, ".html")
	return views
}
