package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/prop"
)

func listFooter(m *model.ItemList, l *PageListeners) dom.Aspect {
	return elem.Footer(
		prop.Id("footer"),

		elem.Span(
			prop.Id("todo-count"),

			elem.Strong(
				bind.TextFunc(bind.Itoa(m.ActiveItemCount), m.Scope),
			),
			bind.IfFunc(func() bool { return m.ActiveItemCount() == 1 }, m.Scope,
				dom.Text(" item left"),
			),
			bind.IfFunc(func() bool { return m.ActiveItemCount() != 1 }, m.Scope,
				dom.Text(" items left"),
			),
		),

		elem.UnorderedList(
			prop.Id("filters"),
			filterButton("All", model.All, m),
			filterButton("Active", model.Active, m),
			filterButton("Completed", model.Completed, m),
		),

		bind.IfFunc(func() bool { return m.CompletedItemCount() != 0 }, m.Scope,
			elem.Button(
				prop.Id("clear-completed"),
				dom.Text("Clear completed ("),
				bind.TextFunc(bind.Itoa(m.CompletedItemCount), m.Scope),
				dom.Text(")"),
				event.Click(l.ClearCompleted),
			),
		),
	)
}

func filterButton(label string, state model.FilterState, m *model.ItemList) dom.Aspect {
	return elem.ListItem(
		elem.Anchor(
			bind.IfFunc(func() bool { return m.Filter == state }, m.Scope,
				prop.Class("selected"),
			),
			prop.Href("#"),
			dom.PreventDefault(event.Click(func(c *dom.EventContext) { m.Filter = state; m.Scope.Digest() })),

			dom.Text(label),
		),
		dom.Text(" "),
	)
}
