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

type FilterButtonProps struct {
	Label  string
	Filter model.FilterState
}

type FilterButton struct {
	*vecty.Core
	*FilterButtonProps
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) Render() vecty.Component {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}

func NewFilterButton(props *FilterButtonProps) *FilterButton {
	b := &FilterButton{
		FilterButtonProps: props,
	}
	b.Core = vecty.New(b)
	return b
}
