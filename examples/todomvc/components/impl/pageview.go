//go:generate gencomponent ../spec

package impl

import (
	"fmt"

	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/components/spec"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/examples/todomvc/store/model"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func (p *PageViewImpl) ComponentDidMount() {
	store.Listeners.Add(p, func() {
		p.items = store.Items
		p.Update()
	})
}

func (p *PageViewImpl) onNewItemTitleInput(event *dom.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	p.Update()
}

func (p *PageViewImpl) onAdd(event *dom.Event) {
	dispatcher.Dispatch(&actions.AddItem{
		Title: p.newItemTitle,
	})
	p.newItemTitle = ""
	p.Update()
}

func (p *PageViewImpl) onClearCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.ClearCompleted{})
}

func (p *PageViewImpl) onToggleAllCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *PageViewImpl) Render() dom.Spec {
	return elem.Div(
		elem.Section(
			prop.Class("todoapp"),

			p.renderHeader(),
			dom.If(len(store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageViewImpl) renderHeader() dom.Markup {
	return elem.Header(
		prop.Class("header"),

		elem.Header1(
			dom.Text("todos"),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(p.onAdd).PreventDefault(),

			elem.Input(
				prop.Class("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(true),
				prop.Value(p.newItemTitle),
				event.Input(p.onNewItemTitleInput),
			),
		),
	)
}

func (p *PageViewImpl) renderFooter() dom.Spec {
	count := store.ActiveItemCount()
	var itemsLeftText = " items left"
	if count == 1 {
		itemsLeftText = " item left"
	}

	return elem.Footer(
		prop.Class("footer"),

		elem.Span(
			prop.Class("todo-count"),

			elem.Strong(
				dom.Text(fmt.Sprintf("%d", count)),
			),
			dom.Text(itemsLeftText),
		),

		elem.UnorderedList(
			prop.Class("filters"),
			&spec.FilterButton{Label: "All", Filter: model.All},
			dom.Text(" "),
			&spec.FilterButton{Label: "Active", Filter: model.Active},
			dom.Text(" "),
			&spec.FilterButton{Label: "Completed", Filter: model.Completed},
		),

		dom.If(store.CompletedItemCount() > 0,
			elem.Button(
				prop.Class("clear-completed"),
				dom.Text(fmt.Sprintf("Clear completed (%d)", store.CompletedItemCount())),
				event.Click(p.onClearCompleted),
			),
		),
	)
}

func (p *PageViewImpl) renderInfo() dom.Spec {
	return elem.Footer(
		prop.Class("info"),

		elem.Paragraph(
			dom.Text("Double-click to edit a todo"),
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

func (p *PageViewImpl) renderItemList() dom.Spec {
	var items dom.List
	for i, item := range store.Items {
		if (store.Filter == model.Active && item.Completed) || (store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, &spec.ItemView{Index: i, Item: item})
	}

	return elem.Section(
		prop.Class("main"),

		elem.Input(
			prop.Id("toggle-all"),
			prop.Class("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			prop.Checked(store.CompletedItemCount() == len(store.Items)),
			event.Change(p.onToggleAllCompleted),
		),
		elem.Label(
			prop.For("toggle-all"),
			dom.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			prop.Class("todo-list"),
			items,
		),
	)
}
