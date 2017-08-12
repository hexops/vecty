// TODO: Not a real example; this won't be merged.
package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func main() {
	vecty.SetTitle("Markdown Demo")
	root := &MyComponent{}
	go root.flip()
	vecty.RenderBody(root)
}

type StatefulComponent struct {
	vecty.Core
	ParentToggle bool `vecty:"prop"` // property set by parent
	toggle       bool // private state
	id           int
}

func (c *StatefulComponent) Render() *vecty.HTML {
	return elem.Div(
		vecty.If(c.ParentToggle, vecty.Text("parentToggle active")),
		vecty.If(!c.ParentToggle, vecty.Text("parentToggle inactive")),
		elem.Break(),
		elem.Button(
			vecty.If(!c.toggle, vecty.Text("Inactive")),
			vecty.If(c.toggle, vecty.Text("Active")),
			event.Click(c.onClick),
		),
		elem.Break(),
		vecty.Text(fmt.Sprint("StatefulComponent instance:", c.id)),
	)
}

func (c *StatefulComponent) onClick(e *vecty.Event) {
	c.toggle = !c.toggle

	// @pdf note this works now :)
	vecty.Rerender(c)
}

type MyComponent struct {
	vecty.Core
	toggle bool
}

func (c *MyComponent) flip() {
	for {
		time.Sleep(20 * time.Millisecond)
		c.toggle = !c.toggle
		vecty.Rerender(c)
	}
}

var id int

func (c *MyComponent) Render() *vecty.HTML {
	id++
	return elem.Body(
		&StatefulComponent{ParentToggle: c.toggle, id: id},
	)
}
