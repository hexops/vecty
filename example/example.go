package main

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func main() {
	scope := bind.NewScope()
	var name string

	elem.Div(
		dom.Text("Your name: "),
		elem.Input(
			prop.Type("text"),
			bind.InputValue(&name, scope),
		),
		elem.H1(
			style.Color("blue"),
			dom.Text("Hello "),
			bind.Text(&name, scope),
			dom.Text("!"),
		),
	).Apply(dom.Body())
}
