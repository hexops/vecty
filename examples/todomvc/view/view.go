package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/prop"
)

func Page(model *model.Model) dom.Aspect {
	return dom.Group(
		elem.Section(
			prop.Id("todoapp"),

			header(model),
			main(model),
			footer(model),
		),

		info(),

		elem.Script(
			prop.Src("bower_components/todomvc-common/base.js"),
		),
	)
}

func header(model *model.Model) dom.Aspect {
	return elem.Header(
		prop.Id("header"),

		elem.H1(
			dom.Text("todos"),
		),
		elem.Input(
			prop.Id("new-todo"),
			prop.Placeholder("What needs to be done?"),
			prop.Autofocus(),
		),
	)
}

func main(model *model.Model) dom.Aspect {
	return elem.Section(
		prop.Id("main"),

		elem.Input(
			prop.Id("toggle-all"),
			prop.Type("checkbox"),
		),
		elem.Label(
			prop.For("toggle-all"),
			dom.Text("Mark all as complete"),
		),

		elem.UL(
			prop.Id("todo-list"),

			bind.Dynamic(model.Scope, func(aspects *bind.Aspects) {
				for _, item := range model.Items {
					if !aspects.Reuse(item) {
						aspects.Add(item, elem.LI(
							bind.If(&item.Completed, model.Scope,
								prop.Class("Completed"),
							),

							elem.Div(
								prop.Class("view"),

								elem.Input(
									prop.Class("toggle"),
									prop.Type("checkbox"),
									bind.Checked(&item.Completed, model.Scope),
								),
								elem.Label(
									bind.Text(&item.Text, model.Scope),
								),
								elem.Button(
									prop.Class("destroy"),
								),
							),
							elem.Input(
								prop.Class("edit"),
								prop.Value(item.Text), // TODO fixme
							),
						))
					}
				}
			}),
		),
	)
}

func footer(model *model.Model) dom.Aspect {
	return elem.Footer(
		prop.Id("footer"),

		elem.Span(
			prop.Id("todo-count"),

			elem.Strong(
				dom.Text("1"),
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

		elem.Button(
			prop.Id("clear-completed"),
			dom.Text("Clear completed (1)"),
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
