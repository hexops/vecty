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

func itemElem(item *model.Item, m *model.Model) dom.Aspect {
	return elem.ListItem(
		bind.IfPtr(&item.Completed, m.Scope,
			prop.Class("completed"),
		),
		bind.IfFunc(func() bool { return item == m.EditItem }, m.Scope,
			prop.Class("editing"),
		),

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				bind.Checked(&item.Completed, m.Scope),
			),
			elem.Label(
				bind.TextPtr(&item.Title, m.Scope),
				event.DblClick(func(c *dom.EventContext) { m.EditItem = item; m.Scope.Digest() }),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(m.DestroyItem(item)),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			dom.PreventDefault(event.Submit(func(c *dom.EventContext) { m.EditItem = nil; m.Scope.Digest() })),
			elem.Input(
				prop.Class("edit"),
				bind.Value(&item.Title, m.Scope),
			),
		),
	)
}
