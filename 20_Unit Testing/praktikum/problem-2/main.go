package main

import (
	"praktikum-unit-testing/config"
	"praktikum-unit-testing/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
