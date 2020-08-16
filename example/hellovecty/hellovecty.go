package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello Vecty!"),
	)
}
