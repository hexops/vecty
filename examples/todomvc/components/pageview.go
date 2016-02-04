package components

import (
	"fmt"

	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/examples/todomvc/store/model"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

type PageView struct {
	dom.Composite

	Items        []*model.Item
	newItemTitle string
}

func (p *PageView) Apply(element *dom.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(oldComp dom.Component) {
	if oldComp, ok := oldComp.(*PageView); ok {
		p.Body = oldComp.Body
		p.newItemTitle = oldComp.newItemTitle
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) onNewItemTitleInput(event *dom.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	p.ReconcileBody()
}

func (p *PageView) onAdd(event *dom.Event) {
	dispatcher.Dispatch(&actions.AddItem{
		Title: p.newItemTitle,
	})
	p.newItemTitle = ""
	p.ReconcileBody()
}

func (p *PageView) onClearCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.ClearCompleted{})
}

func (p *PageView) onToggleAllCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *PageView) render() dom.Component {
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

func (p *PageView) renderHeader() dom.Markup {
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

func (p *PageView) renderFooter() dom.Component {
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
			&FilterButton{Label: "All", Filter: model.All},
			dom.Text(" "),
			&FilterButton{Label: "Active", Filter: model.Active},
			dom.Text(" "),
			&FilterButton{Label: "Completed", Filter: model.Completed},
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

func (p *PageView) renderInfo() dom.Component {
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

func (p *PageView) renderItemList() dom.Component {
	var items dom.List
	for i, item := range store.Items {
		if (store.Filter == model.Active && item.Completed) || (store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, &ItemView{Index: i, Item: item})
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
