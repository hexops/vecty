//go:generate gencomponent ../spec

package impl

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/components/spec"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func (p *PageViewImpl) ComponentDidMount() {
	store.Listeners.Add(p, p.update)
}

func (p *PageViewImpl) update() {
	p.State().SetItems(store.Items)
}

func (p *PageViewImpl) Render() dom.Spec {
	return elem.Div(
		elem.Section(
			prop.Id("todoapp"),

			p.renderHeader(),
			dom.If(len(store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),

		// elem.Script(
		// 	prop.Src("bower_components/todomvc-common/base.js"),
		// ),
	)
}

func (p *PageViewImpl) renderHeader() dom.Markup {
	return elem.Header(
		prop.Id("header"),

		elem.Header1(
			dom.Text("todos"),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			// dom.PreventDefault(event.Submit(l.AddItem)),

			elem.Input(
				prop.Id("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(true),
				// bind.Value(&m.AddItemTitle, m.Scope),
			),
		),
	)
}

func (p *PageViewImpl) renderFooter() dom.Spec {
	return elem.Footer(
		prop.Id("footer"),

		elem.Span(
			prop.Id("todo-count"),

			elem.Strong(
			// bind.TextFunc(bind.Itoa(m.ActiveItemCount), m.Scope),
			),
			// bind.IfFunc(func() bool { return m.ActiveItemCount() == 1 }, m.Scope,
			// 	dom.Text(" item left"),
			// ),
			// bind.IfFunc(func() bool { return m.ActiveItemCount() != 1 }, m.Scope,
			// 	dom.Text(" items left"),
			// ),
		),

		elem.UnorderedList(
			prop.Id("filters"),
			&spec.FilterButton{Label: "All", Selected: store.Filter == store.All},
			&spec.FilterButton{Label: "Active", Selected: store.Filter == store.Active},
			&spec.FilterButton{Label: "Completed", Selected: store.Filter == store.Completed},
			// filterButton("All", model.All, m),
			// filterButton("Active", model.Active, m),
			// filterButton("Completed", model.Completed, m),
		),

		// bind.IfFunc(func() bool { return m.CompletedItemCount() != 0 }, m.Scope,
		// 	elem.Button(
		// 		prop.Id("clear-completed"),
		// 		dom.Text("Clear completed ("),
		// 		bind.TextFunc(bind.Itoa(m.CompletedItemCount), m.Scope),
		// 		dom.Text(")"),
		// 		event.Click(l.ClearCompleted),
		// 	),
		// ),
	)
}

func (p *PageViewImpl) renderInfo() dom.Spec {
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

func (p *PageViewImpl) onToggleAllCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *PageViewImpl) renderItemList() dom.Spec {
	var items dom.List
	for i, item := range store.Items {
		items = append(items, &spec.ItemView{Index: i, Item: item})
	}

	return elem.Section(
		prop.Id("main"),

		elem.Input(
			prop.Id("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			prop.Checked(store.CompletedItemCount() == len(store.Items)),
			event.Change(p.onToggleAllCompleted),
		),
		elem.Label(
			prop.For("toggle-all"),
			dom.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			prop.Id("todo-list"),
			items,
			// bind.Dynamic(m.Scope, func(aspects *bind.Aspects) {
			// 	for _, item := range m.Items {
			// 		if !(m.Filter == model.All || (m.Filter == model.Active && !item.Completed) || (m.Filter == model.Completed && item.Completed)) {
			// 			continue
			// 		}
			// 		if !aspects.Reuse(item) {
			// 			theItem := item
			// 			editing := func() bool { return theItem == m.EditItem }
			// 			aspects.Add(item, itemElem(item, editing, l))
			// 		}
			// 	}
			// }),
		),
	)
}

func (p *FilterButtonImpl) Render() dom.Spec {
	return elem.ListItem(
		elem.Anchor(
			dom.If(p.Props().Selected(), prop.Class("selected")),
			prop.Href("#"),
			// dom.PreventDefault(event.Click(func(c *dom.EventContext) { m.Filter = state; m.Scope.Digest() })),

			dom.Text(p.Props().Label()),
		),
		dom.Text(" "),
	)
}
