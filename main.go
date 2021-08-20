package main

import (
	"github.com/codethecats/go-todo-api/app"
	"github.com/codethecats/go-todo-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
