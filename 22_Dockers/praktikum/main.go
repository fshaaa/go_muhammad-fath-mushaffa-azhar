package main

import (
	"docker-problem/config"
	"docker-problem/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
