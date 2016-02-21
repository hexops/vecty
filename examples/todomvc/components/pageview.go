package components

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/examples/todomvc/actions"
	"github.com/gopherjs/vecty/examples/todomvc/dispatcher"
	"github.com/gopherjs/vecty/examples/todomvc/store"
	"github.com/gopherjs/vecty/examples/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

type PageView struct {
	vecty.Composite

	Items        []*model.Item
	newItemTitle string
}

// Apply implements the vecty.Markup interface.
func (p *PageView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*PageView); ok {
		p.Body = oldComp.Body
		p.newItemTitle = oldComp.newItemTitle
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) onNewItemTitleInput(event *vecty.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	p.ReconcileBody()
}

func (p *PageView) onAdd(event *vecty.Event) {
	dispatcher.Dispatch(&actions.AddItem{
		Title: p.newItemTitle,
	})
	p.newItemTitle = ""
	p.ReconcileBody()
}

func (p *PageView) onClearCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.ClearCompleted{})
}

func (p *PageView) onToggleAllCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *PageView) render() vecty.Component {
	return elem.Div(
		elem.Section(
			prop.Class("todoapp"),

			p.renderHeader(),
			vecty.If(len(store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() vecty.Markup {
	return elem.Header(
		prop.Class("header"),

		elem.Header1(
			vecty.Text("todos"),
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

func (p *PageView) renderFooter() vecty.Component {
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
				vecty.Text(fmt.Sprintf("%d", count)),
			),
			vecty.Text(itemsLeftText),
		),

		elem.UnorderedList(
			prop.Class("filters"),
			&FilterButton{Label: "All", Filter: model.All},
			vecty.Text(" "),
			&FilterButton{Label: "Active", Filter: model.Active},
			vecty.Text(" "),
			&FilterButton{Label: "Completed", Filter: model.Completed},
		),

		vecty.If(store.CompletedItemCount() > 0,
			elem.Button(
				prop.Class("clear-completed"),
				vecty.Text(fmt.Sprintf("Clear completed (%d)", store.CompletedItemCount())),
				event.Click(p.onClearCompleted),
			),
		),
	)
}

func (p *PageView) renderInfo() vecty.Component {
	return elem.Footer(
		prop.Class("info"),

		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				prop.Href("http://github.com/neelance"),
				vecty.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				prop.Href("http://todomvc.com"),
				vecty.Text("TodoMVC"),
			),
		),
	)
}

func (p *PageView) renderItemList() vecty.Component {
	var items vecty.List
	for i, item := range store.Items {
		if (store.Filter == model.Active && item.Completed) || (store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, &ItemView{Index: i, Item: item})
	}

	return elem.Section(
		prop.Class("main"),

		elem.Input(
			prop.ID("toggle-all"),
			prop.Class("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			prop.Checked(store.CompletedItemCount() == len(store.Items)),
			event.Change(p.onToggleAllCompleted),
		),
		elem.Label(
			prop.For("toggle-all"),
			vecty.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			prop.Class("todo-list"),
			items,
		),
	)
}
