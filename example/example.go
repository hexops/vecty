package main

import (
	"github.com/neelance/dom"
	_ "github.com/neelance/dom/example/components/impl"
	"github.com/neelance/dom/example/components/spec"
)

func main() {
	dom.RenderAsBody(
		&spec.Greeter{
			Greeting: "Hello",
		},
	)
}
