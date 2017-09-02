package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/example/todomvc/actions"
	"github.com/gopherjs/vecty/example/todomvc/dispatcher"
	"github.com/gopherjs/vecty/example/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

// ItemView is a vecty.Component which represents a single item in the TODO
// list.
type ItemView struct {
	vecty.Core

	Index     int         `vecty:"prop"`
	Item      *model.Item `vecty:"prop"`
	editing   bool
	editTitle string
	input     *vecty.HTML
}

// Restore implements the vecty.Restorer interface.
func (p *ItemView) Restore(prev vecty.Component) {
	if old, ok := prev.(*ItemView); ok {
		p.editing = old.editing
		p.editTitle = old.editTitle
	}
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
	p.input.Node().Call("focus")
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

// Render implements the vecty.Component interface.
func (p *ItemView) Render() *vecty.HTML {
	p.input = elem.Input(
		vecty.Markup(
			prop.Class("edit"),
			prop.Value(p.editTitle),
			event.Input(p.onEditInput),
		),
	)

	return elem.ListItem(
		vecty.Markup(
			vecty.ClassMap{
				"completed": p.Item.Completed,
				"editing":   p.editing,
			},
		),

		elem.Div(
			vecty.Markup(
				prop.Class("view"),
			),

			elem.Input(
				vecty.Markup(
					prop.Class("toggle"),
					prop.Type(prop.TypeCheckbox),
					prop.Checked(p.Item.Completed),
					event.Change(p.onToggleCompleted),
				),
			),
			elem.Label(
				vecty.Markup(
					event.DoubleClick(p.onStartEdit),
				),
				vecty.Text(p.Item.Title),
			),
			elem.Button(
				vecty.Markup(
					prop.Class("destroy"),
					event.Click(p.onDestroy),
				),
			),
		),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onStopEdit).PreventDefault(),
			),
			p.input,
		),
	)
}
