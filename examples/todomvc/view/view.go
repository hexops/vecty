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

func Page(m *model.Model) dom.Aspect {
	return dom.Group(
		elem.Section(
			prop.Id("todoapp"),

			header(m),
			bind.IfFunc(func() bool { return len(m.Items) != 0 }, m.Scope,
				main(m),
				footer(m),
			),
		),

		info(),

		elem.Script(
			prop.Src("bower_components/todomvc-common/base.js"),
		),
	)
}

func header(m *model.Model) dom.Aspect {
	return elem.Header(
		prop.Id("header"),

		elem.Header1(
			dom.Text("todos"),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			dom.PreventDefault(event.Submit(m.AddItem)),

			elem.Input(
				prop.Id("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(),
				bind.Value(&m.AddItemTitle, m.Scope),
			),
		),
	)
}

func main(m *model.Model) dom.Aspect {
	return elem.Section(
		prop.Id("main"),

		elem.Input(
			prop.Id("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			bind.IfFunc(func() bool { return m.CompletedItemCount() == len(m.Items) }, m.Scope,
				prop.Checked(),
			),
			event.Change(m.ToggleAll),
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
						aspects.Add(item, itemElem(item, m))
					}
				}
			}),
		),
	)
}

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

func footer(m *model.Model) dom.Aspect {
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
				event.Click(m.ClearCompleted),
			),
		),
	)
}

func filterButton(label string, state model.FilterState, m *model.Model) dom.Aspect {
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

func info() dom.Aspect {
	return elem.Footer(
		prop.Id("info"),

		elem.Paragraph(
			dom.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			dom.Text("Template by "),
			elem.Anchor(
				prop.Href("http://github.com/sindresorhus"),
				dom.Text("Sindre Sorhus"),
			),
		),
		elem.Paragraph(
			dom.Text("Created by "),
			elem.Anchor(
				prop.Href("http://github.com/neelance"),
				dom.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			dom.Text("Part of "),
			elem.Anchor(
				prop.Href("http://todomvc.com"),
				dom.Text("TodoMVC"),
			),
		),
	)
}
