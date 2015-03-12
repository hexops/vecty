package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/prop"
)

type PageListeners struct {
	AddItem        dom.Listener
	DestroyItem    func(*model.Item) dom.Listener
	StartEdit      func(*model.Item) dom.Listener
	StopEdit       dom.Listener
	ClearCompleted dom.Listener
	ToggleAll      dom.Listener
}

func Page(m *model.ItemList, l *PageListeners) dom.Aspect {
	return dom.Group(
		elem.Section(
			prop.Id("todoapp"),

			listHeader(m, l),
			bind.IfFunc(func() bool { return len(m.Items) != 0 }, m.Scope,
				itemList(m, l),
				listFooter(m, l),
			),
		),

		info(),

		elem.Script(
			prop.Src("bower_components/todomvc-common/base.js"),
		),
	)
}

func itemList(m *model.ItemList, l *PageListeners) dom.Aspect {
	return elem.Section(
		prop.Id("main"),

		elem.Input(
			prop.Id("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			bind.IfFunc(func() bool { return m.CompletedItemCount() == len(m.Items) }, m.Scope,
				prop.Checked(),
			),
			event.Change(l.ToggleAll),
		),
		elem.Label(
			prop.For("toggle-all"),
			dom.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			prop.Id("todo-list"),

			bind.Dynamic(m.Scope, func(aspects *bind.Aspects) {
				for _, item := range m.Items {
					if !(m.Filter == model.All || (m.Filter == model.Active && !item.Completed) || (m.Filter == model.Completed && item.Completed)) {
						continue
					}
					if !aspects.Reuse(item) {
						theItem := item
						editing := func() bool { return theItem == m.EditItem }
						aspects.Add(item, itemElem(item, editing, l))
					}
				}
			}),
		),
	)
}
