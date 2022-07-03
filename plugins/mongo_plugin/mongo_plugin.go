package mongo_plugin

import (
	"context"
	"fmt"
	"gofiber-demo/plugins/app_plugin"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Register(app *app_plugin.Application, ENV map[string]string) *mongo.Database {
	var MONGODB_PROTOCOL = ENV["MONGODB_PROTOCOL"]
	var MONGODB_HOSTS = ENV["MONGODB_HOSTS"]
	var MONGODB_DATABASE = ENV["MONGODB_DATABASE"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s", MONGODB_PROTOCOL, MONGODB_HOSTS)))
	if err != nil {
		panic(err)
	}
	app.OnShutdown(func() {
		client.Disconnect(ctx)
	})

	db := client.Database(MONGODB_DATABASE)

	return db
}
