package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/examples/todomvc/actions"
	"github.com/gopherjs/vecty/examples/todomvc/dispatcher"
	"github.com/gopherjs/vecty/examples/todomvc/store"
	"github.com/gopherjs/vecty/examples/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
)

type FilterButton struct {
	vecty.Core

	Label  string
	Filter model.FilterState
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) Render() *vecty.HTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}
