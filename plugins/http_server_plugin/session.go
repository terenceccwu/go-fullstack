package http_server_plugin

import (
	"fmt"
	"gofiber-demo/plugins/env_plugin"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/storage/sqlite3"
)

type SessionEngine interface {
	GetSession(c *fiber.Ctx) *session.Session
	GetSessionStore() *session.Store
}

type BaseSessionEngine struct {
	SessionStore *session.Store
}

func (e *BaseSessionEngine) GetSession(c *fiber.Ctx) *session.Session {
	session, err := e.SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	return session
}

func (e *BaseSessionEngine) GetSessionStore() *session.Store {
	return e.SessionStore
}

func RegisterSessionEngine(backend string) SessionEngine {
	var storage fiber.Storage
	switch backend {
	case "sqlite3":
		storage = sqlite3.New()
	case "redis":
		storage = redis.New(redis.Config{
			URL:   fmt.Sprintf("redis://%s:%d/%d", env_plugin.ENV["SESSION_REDIS_HOST"], env_plugin.GetEnvInt("SESSION_REDIS_PORT", "6379"), env_plugin.GetEnvInt("SESSION_REDIS_DB", "0")),
			Reset: false,
		})
	case "memory":
		storage = memory.New()
	default:
		storage = memory.New()
	}

	engine := &BaseSessionEngine{
		SessionStore: session.New(session.Config{
			Storage: storage,
		}),
	}

	return engine
}
