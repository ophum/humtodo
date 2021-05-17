package main

import (
	"fmt"
	"sort"

	"github.com/ophum/humtodo/pkg/routes"
)

func main() {

	e := routes.Init()

	rr := e.Routes()
	sort.Slice(rr, func(i, j int) bool {
		return rr[i].Path < rr[j].Path
	})

	for _, r := range rr {
		if r.Name == "github.com/labstack/echo.(*Group).Use.func1" {
			continue
		}
		fmt.Printf("%s\t%s\t%s\n", r.Method, r.Path, r.Name)
	}

	e.Logger.Fatal(e.Start(":8080"))
}
