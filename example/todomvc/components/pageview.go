package components

import (
	"fmt"

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
func (p *PageView) Render() *vecty.HTML {
	return elem.Body(
		elem.Section(
			vecty.Markup(
				prop.Class("todoapp"),
			),

			p.renderHeader(),
			vecty.If(len(store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() *vecty.HTML {
	return elem.Header(
		vecty.Markup(
			prop.Class("header"),
		),

		elem.Heading1(
			vecty.Text("todos"),
		),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onAdd).PreventDefault(),
			),

			elem.Input(
				vecty.Markup(
					prop.Class("new-todo"),
					prop.Placeholder("What needs to be done?"),
					prop.Autofocus(true),
					prop.Value(p.newItemTitle),
					event.Input(p.onNewItemTitleInput),
				),
			),
		),
	)
}

func (p *PageView) renderFooter() *vecty.HTML {
	count := store.ActiveItemCount()
	var itemsLeftText = " items left"
	if count == 1 {
		itemsLeftText = " item left"
	}

	return elem.Footer(
		vecty.Markup(
			prop.Class("footer"),
		),

		elem.Span(
			vecty.Markup(
				prop.Class("todo-count"),
			),

			elem.Strong(
				vecty.Text(fmt.Sprintf("%d", count)),
			),
			vecty.Text(itemsLeftText),
		),

		elem.UnorderedList(
			vecty.Markup(
				prop.Class("filters"),
			),
			&FilterButton{Label: "All", Filter: model.All},
			vecty.Text(" "),
			&FilterButton{Label: "Active", Filter: model.Active},
			vecty.Text(" "),
			&FilterButton{Label: "Completed", Filter: model.Completed},
		),

		vecty.If(store.CompletedItemCount() > 0,
			elem.Button(
				vecty.Markup(
					prop.Class("clear-completed"),
					event.Click(p.onClearCompleted),
				),
				vecty.Text(fmt.Sprintf("Clear completed (%d)", store.CompletedItemCount())),
			),
		),
	)
}

func (p *PageView) renderInfo() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(
			prop.Class("info"),
		),

		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://github.com/neelance"),
				),
				vecty.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://todomvc.com"),
				),
				vecty.Text("TodoMVC"),
			),
		),
	)
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
		vecty.Markup(
			prop.Class("main"),
		),

		elem.Input(
			vecty.Markup(
				prop.ID("toggle-all"),
				prop.Class("toggle-all"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(store.CompletedItemCount() == len(store.Items)),
				event.Change(p.onToggleAllCompleted),
			),
		),
		elem.Label(
			vecty.Markup(
				prop.For("toggle-all"),
			),
			vecty.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			vecty.Markup(
				prop.Class("todo-list"),
			),
			items,
		),
	)
}
