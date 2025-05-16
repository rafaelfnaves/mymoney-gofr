package main

import (
	"github.com/rafaelfnaves/mymoney-gofr/internal/handlers"
	"github.com/rafaelfnaves/mymoney-gofr/migrations"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialize gofr
	app := gofr.New()

	// Run migrations
	app.Migrate(migrations.All())

	// register routes
	app.POST("/expenses", handlers.AddExpense)
	app.GET("/expenses", handlers.ListExpenses)
	app.GET("/expenses/{id}", handlers.GetExpense)
	app.DELETE("/expenses/{id}", handlers.DeleteExpense)

	// Runs the server
	app.Run()
}
