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

func itemElem(item *model.Item, editing func() bool, l *PageListeners) dom.Aspect {
	return elem.ListItem(
		bind.IfPtr(&item.Completed, item.Scope,
			prop.Class("completed"),
		),
		bind.IfFunc(editing, item.Scope,
			prop.Class("editing"),
		),

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				bind.Checked(&item.Completed, item.Scope),
			),
			elem.Label(
				bind.TextPtr(&item.Title, item.Scope),
				event.DblClick(l.StartEdit(item)),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(l.DestroyItem(item)),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			dom.PreventDefault(event.Submit(l.StopEdit)),
			elem.Input(
				prop.Class("edit"),
				bind.Value(&item.Title, item.Scope),
			),
		),
	)
}
