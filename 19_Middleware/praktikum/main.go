package main

import (
	"praktikum/config"
	"praktikum/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
