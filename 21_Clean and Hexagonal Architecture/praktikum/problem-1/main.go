package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	routes.NewRoute(app, db)
	app.Start(":8080")
}
