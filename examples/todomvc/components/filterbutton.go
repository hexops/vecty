package components

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/examples/todomvc/store/model"
	"github.com/neelance/dom/prop"
)

type FilterButton struct {
	dom.Composite

	Label  string
	Filter model.FilterState
}

func (b *FilterButton) Apply(element *dom.Element) {
	element.AddChild(b)
}

func (b *FilterButton) Reconcile(oldComp dom.Component) {
	if oldComp, ok := oldComp.(*FilterButton); ok {
		b.Body = oldComp.Body
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func (b *FilterButton) onClick(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) render() dom.Component {
	return elem.ListItem(
		elem.Anchor(
			dom.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			dom.Text(b.Label),
		),
	)
}
