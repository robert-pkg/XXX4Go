package main

import (
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/app"
)

func main() {
	app := app.NewApp()
	app.Init()

	app.Run()
}
