package main

import (
	"github.com/rafaelfnaves/mymoney-gofr/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialize gofr
	app := gofr.New()

	// Run migrations
	app.Migrate(migrations.All())

	// register routes
	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello World!", nil
	})

	// Runs the server
	app.Run()
}
