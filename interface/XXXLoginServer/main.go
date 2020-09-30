package main

import (
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/app"
)

func main() {
	app := app.NewApp()
	app.Init()

	app.Run()
}
