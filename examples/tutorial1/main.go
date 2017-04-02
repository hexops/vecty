package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.SetTitle("Vecty Tutorial")
	c := &MyComponent{}
	vecty.RenderBody(c)
}

type MyComponent struct {
	vecty.Core
}

func (c *MyComponent) Render() *vecty.HTML {
	return elem.Body(
		elem.Heading1(vecty.Text("Hello, World")),
	)
}
