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
			main(m),
			footer(m),
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

		elem.H1(
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
			prop.Type("checkbox"),
			bind.IfFunc(func() bool { return m.CompletedItemCount() == len(m.Items) }, m.Scope,
				prop.Checked(),
			),
			event.Change(m.ToggleAll),
		),
		elem.Label(
			prop.For("toggle-all"),
			dom.Text("Mark all as complete"),
		),

		elem.UL(
			prop.Id("todo-list"),

			bind.Dynamic(m.Scope, func(aspects *bind.Aspects) {
				for _, item := range m.Items {
					if !aspects.Reuse(item) {
						aspects.Add(item, elem.LI(
							bind.IfPtr(&item.Completed, m.Scope,
								prop.Class("Completed"),
							),

							elem.Div(
								prop.Class("view"),

								elem.Input(
									prop.Class("toggle"),
									prop.Type("checkbox"),
									bind.Checked(&item.Completed, m.Scope),
								),
								elem.Label(
									bind.TextPtr(&item.Title, m.Scope),
								),
								elem.Button(
									prop.Class("destroy"),
								),
							),
							elem.Input(
								prop.Class("edit"),
								prop.Value(item.Title), // TODO fixme
							),
						))
					}
				}
			}),
		),
	)
}

func footer(m *model.Model) dom.Aspect {
	return elem.Footer(
		prop.Id("footer"),

		elem.Span(
			prop.Id("todo-count"),

			elem.Strong(
				bind.TextFunc(bind.Itoa(m.IncompleteItemCount), m.Scope),
			),
			dom.Text(" item left"),
		),

		elem.UL(
			prop.Id("filters"),

			elem.LI(
				elem.A(
					prop.Class("selected"),
					dom.Text("All"),
				),
				dom.Text(" "),
			),
			elem.LI(
				elem.A(
					dom.Text("Active"),
				),
				dom.Text(" "),
			),
			elem.LI(
				elem.A(
					dom.Text("Completed"),
				),
				dom.Text(" "),
			),
		),

		bind.IfFunc(func() bool { return m.CompletedItemCount() != 0 }, m.Scope,
			elem.Button(
				prop.Id("clear-completed"),
				dom.Text("Clear completed ("),
				bind.TextFunc(bind.Itoa(m.CompletedItemCount), m.Scope),
				dom.Text(")"),
			),
		),
	)
}

func info() dom.Aspect {
	return elem.Footer(
		prop.Id("info"),

		elem.P(
			dom.Text("Double-click to edit a todo"),
		),
		elem.P(
			dom.Text("Template by "),
			elem.A(
				prop.HRef("http://github.com/sindresorhus"),
				dom.Text("Sindre Sorhus"),
			),
		),
		elem.P(
			dom.Text("Created by "),
			elem.A(
				prop.HRef("http://github.com/neelance"),
				dom.Text("Richard Musiol"),
			),
		),
		elem.P(
			dom.Text("Part of "),
			elem.A(
				prop.HRef("http://todomvc.com"),
				dom.Text("TodoMVC"),
			),
		),
	)
}
