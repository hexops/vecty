package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func listHeader(m *model.ItemList, l *PageListeners) dom.Aspect {
	return elem.Header(
		prop.Id("header"),

		elem.Header1(
			dom.Text("todos"),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			dom.PreventDefault(event.Submit(l.AddItem)),

			elem.Input(
				prop.Id("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(),
				bind.Value(&m.AddItemTitle, m.Scope),
			),
		),
	)
}
