package main

import (
	"praktikum-weekly-task-2/config"
	"praktikum-weekly-task-2/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}