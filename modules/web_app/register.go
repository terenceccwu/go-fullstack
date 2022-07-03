package web_app

import (
	"fmt"
	"gofiber-demo/modules/web_app/controllers"
	"gofiber-demo/plugins/http_server_plugin"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

var CWD string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	CWD = filepath.Dir(filename)
}

func Register(server *fiber.App, session_engine http_server_plugin.SessionEngine) {
	view_engine := RegisterViewEngine()

	home := &controllers.HomeController{ViewEngine: view_engine, SessionEngine: session_engine}

	server.Static("/web_app/static/", fmt.Sprintf("%s/static", CWD))
	server.Get("/", home.Index)
	server.Get("/session", home.Session)
}

func RegisterViewEngine() http_server_plugin.ViewEngine {
	views_directory := fmt.Sprintf("%s/views", CWD)
	view_engine := http_server_plugin.NewBaseViewEngine(views_directory, ".html")
	return view_engine
}
