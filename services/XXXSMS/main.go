package main

import (
	"github.com/robert-pkg/XXX4Go/services/XXXSMS/app"
)

func main() {

	app := app.NewApp()
	app.Init()

	app.Run()

}
