package main

import (
	"be15/mvc/config"
	"be15/mvc/middlewares"
	"be15/mvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	middlewares.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":80"))
}
