package components

import (
	"strconv"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/example/todomvc/actions"
	"github.com/gopherjs/vecty/example/todomvc/dispatcher"
	"github.com/gopherjs/vecty/example/todomvc/store"
	"github.com/gopherjs/vecty/example/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

// PageView is a vecty.Component which represents the entire page.
type PageView struct {
	vecty.Core

	Items        []*model.Item `vecty:"prop"`
	newItemTitle string
}

// Restore implements the vecty.Restorer interface.
func (p *PageView) Restore(prev vecty.Component) {
	if old, ok := prev.(*PageView); ok {
		p.newItemTitle = old.newItemTitle
	}
}

func (p *PageView) onNewItemTitleInput(event *vecty.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	vecty.Rerender(p)
}

func (p *PageView) onAdd(event *vecty.Event) {
	dispatcher.Dispatch(&actions.AddItem{
		Title: p.newItemTitle,
	})
	p.newItemTitle = ""
	vecty.Rerender(p)
}

func (p *PageView) onClearCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.ClearCompleted{})
}

func (p *PageView) onToggleAllCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Section(
			p.renderHeader(),
			vecty.If(len(store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		).WithMarkup(vecty.Class("todoapp")),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() *vecty.HTML {
	return elem.Header(
		elem.Heading1(
			vecty.Text("todos"),
		),
		elem.Form(
			elem.Input().WithMarkup(
				vecty.Class("new-todo"),
				prop.Placeholder("What needs to be done?"),
				prop.Autofocus(true),
				prop.Value(p.newItemTitle),
				event.Input(p.onNewItemTitleInput),
			),
		).WithMarkup(
			style.Margin(style.Px(0)),
			event.Submit(p.onAdd).PreventDefault(),
		),
	).WithMarkup(
		vecty.Class("header"),
	)
}

func (p *PageView) renderFooter() *vecty.HTML {
	count := store.ActiveItemCount()
	var itemsLeftText = " items left"
	if count == 1 {
		itemsLeftText = " item left"
	}

	return elem.Footer(
		elem.Span(
			elem.Strong(
				vecty.Text(strconv.Itoa(count)),
			),
			vecty.Text(itemsLeftText),
		).WithMarkup(vecty.Class("todo-count")),

		elem.UnorderedList(
			&FilterButton{Label: "All", Filter: model.All},
			vecty.Text(" "),
			&FilterButton{Label: "Active", Filter: model.Active},
			vecty.Text(" "),
			&FilterButton{Label: "Completed", Filter: model.Completed},
		).WithMarkup(vecty.Class("filters")),

		vecty.If(store.CompletedItemCount() > 0,
			elem.Button(
				vecty.Text("Clear completed ("+strconv.Itoa(store.CompletedItemCount())+")"),
			).WithMarkup(
				vecty.Class("clear-completed"),
				event.Click(p.onClearCompleted),
			),
		),
	).WithMarkup(vecty.Class("footer"))
}

func (p *PageView) renderInfo() *vecty.HTML {
	return elem.Footer(
		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				vecty.Text("Richard Musiol"),
			).WithMarkup(prop.Href("http://github.com/neelance")),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				vecty.Text("TodoMVC"),
			).WithMarkup(prop.Href("http://todomvc.com")),
		),
	).WithMarkup(vecty.Class("info"))
}

func (p *PageView) renderItemList() *vecty.HTML {
	var items vecty.List
	for i, item := range store.Items {
		if (store.Filter == model.Active && item.Completed) || (store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, &ItemView{Index: i, Item: item})
	}

	return elem.Section(
		elem.Input().WithMarkup(
			vecty.Class("toggle-all"),
			prop.ID("toggle-all"),
			prop.Type(prop.TypeCheckbox),
			prop.Checked(store.CompletedItemCount() == len(store.Items)),
			event.Change(p.onToggleAllCompleted),
		),
		elem.Label(
			vecty.Text("Mark all as complete"),
		).WithMarkup(prop.For("toggle-all")),

		elem.UnorderedList(
			items,
		).WithMarkup(vecty.Class("todo-list")),
	).WithMarkup(vecty.Class("main"))
}
