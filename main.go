package main

import "github.com/ophum/humtodo/pkg/routes"

func main() {

	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
