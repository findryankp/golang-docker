package main

import (
	"be15/mvc/config"
	"be15/mvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()
	e.Logger.Fatal(e.Start(":80"))
}
