package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/examples/todomvc/actions"
	"github.com/gopherjs/vecty/examples/todomvc/dispatcher"
	"github.com/gopherjs/vecty/examples/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

type ItemView struct {
	vecty.Core

	Index     int
	Item      *model.Item
	editing   bool
	editTitle string
	input     *vecty.HTML
}

func (p *ItemView) Restore(prev vecty.Component) bool {
	if old, ok := prev.(*ItemView); ok {
		p.editing = old.editing
		p.editTitle = old.editTitle
	}
	return false
}

func (p *ItemView) onDestroy(event *vecty.Event) {
	dispatcher.Dispatch(&actions.DestroyItem{
		Index: p.Index,
	})
}

func (p *ItemView) onToggleCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetCompleted{
		Index:     p.Index,
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *ItemView) onStartEdit(event *vecty.Event) {
	p.editing = true
	p.editTitle = p.Item.Title
	vecty.Rerender(p)
	p.input.Node.Call("focus")
}

func (p *ItemView) onEditInput(event *vecty.Event) {
	p.editTitle = event.Target.Get("value").String()
	vecty.Rerender(p)
}

func (p *ItemView) onStopEdit(event *vecty.Event) {
	p.editing = false
	vecty.Rerender(p)
	dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Index,
		Title: p.editTitle,
	})
}

func (p *ItemView) Render() *vecty.HTML {
	p.input = elem.Input(
		prop.Class("edit"),
		prop.Value(p.editTitle),
		event.Input(p.onEditInput),
	)

	return elem.ListItem(
		vecty.ClassMap{
			"completed": p.Item.Completed,
			"editing":   p.editing,
		},

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(p.Item.Completed),
				event.Change(p.onToggleCompleted),
			),
			elem.Label(
				vecty.Text(p.Item.Title),
				event.DoubleClick(p.onStartEdit),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(p.onDestroy),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(p.onStopEdit).PreventDefault(),
			p.input,
		),
	)
}
