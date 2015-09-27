package main

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	_ "github.com/neelance/dom/example/components/impl"
	"github.com/neelance/dom/example/components/spec"
)

func main() {
	dom.RenderAsBody(
		elem.Div(
			dom.Text("This is a very simple component:"),
			&spec.Greeter{
				Greeting: "Hello",
			},
		),
	)
}
