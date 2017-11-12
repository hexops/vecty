package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() *vecty.HTML {
	return elem.Body(
		// this works:
		//nil,

		// this does not:
		p.returnsNil(),
	)
}

func (p *PageView) returnsNil() vecty.List {
	return nil
}
