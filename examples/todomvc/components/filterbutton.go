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
	vecty.Composite

	Label  string
	Filter model.FilterState
}

// Apply implements the vecty.Markup interface.
func (b *FilterButton) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *FilterButton) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*FilterButton); ok {
		b.Body = oldComp.Body
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) render() vecty.Component {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}
