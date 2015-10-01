package impl

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/prop"
)

func (b *FilterButtonImpl) onClick(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Props().Filter(),
	})
}

func (b *FilterButtonImpl) Render() dom.Spec {
	return elem.ListItem(
		elem.Anchor(
			dom.If(store.Filter == b.Props().Filter(), prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			dom.Text(b.Props().Label()),
		),
	)
}
