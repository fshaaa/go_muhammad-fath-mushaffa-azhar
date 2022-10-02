package main

import (
	"problem-2/config"
	"problem-2/routes"
)

func main() {
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
