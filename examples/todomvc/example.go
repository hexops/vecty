package main

import (
	"github.com/neelance/dom"
	_ "github.com/neelance/dom/examples/todomvc/components/impl"
	"github.com/neelance/dom/examples/todomvc/components/spec"
)

func main() {
	dom.SetTitle("GopherJS • TodoMVC")
	dom.AddStylesheet("bower_components/todomvc-common/base.css")
	dom.RenderAsBody(&spec.PageView{})
}
