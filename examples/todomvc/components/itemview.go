package components

import (
	"strconv"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/examples/todomvc/actions"
	"github.com/gopherjs/vecty/examples/todomvc/dispatcher"
	"github.com/gopherjs/vecty/examples/todomvc/store/model"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/vecty/style"
)

// ItemView component
type ItemView struct {
	vecty.Core

	Index              int
	Item               *model.Item
	completed, editing bool
	title, editTitle   string
	input              *vecty.HTML
}

func (i *ItemView) onDestroy(event *vecty.Event) {
	dispatcher.Dispatch(&actions.DestroyItem{
		Index: i.Index,
	})
}

func (i *ItemView) onToggleCompleted(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetCompleted{
		Index:     i.Index,
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (i *ItemView) onStartEdit(event *vecty.Event) {
	i.editing = true
	i.editTitle = i.title
	vecty.Rerender(i)
	i.input.Node().Call("focus")
}

func (i *ItemView) onEditInput(event *vecty.Event) {
	i.editTitle = event.Target.Get("value").String()
}

func (i *ItemView) onStopEdit(event *vecty.Event) {
	i.editing = false
	vecty.Rerender(i)
	dispatcher.Dispatch(&actions.SetTitle{
		Index: i.Index,
		Title: i.editTitle,
	})
}

// Key implements vecty.Keyer
func (i *ItemView) Key() string {
	return strconv.Itoa(i.Index)
}

// Clone implements vecty.Updater
func (i *ItemView) Clone(prev vecty.Component) vecty.Component {
	c := &ItemView{}
	*c = *i
	return c
}

// ShouldUpdate implements vecty.Updater
func (i *ItemView) ShouldUpdate(prev vecty.Component) bool {
	p, ok := prev.(*ItemView)
	if !ok {
		return true
	}
	if i.Item != p.Item ||
		i.title != p.title ||
		i.completed != p.completed ||
		i.editing != p.editing ||
		i.editTitle != p.editTitle {
		return true
	}

	return false
}

// Render implements vecty.Component
func (i *ItemView) Render() *vecty.HTML {
	i.title = i.Item.Title
	i.completed = i.Item.Completed
	i.input = elem.Input(
		prop.Class("edit"),
		prop.Value(i.editTitle),
		event.Input(i.onEditInput),
	)

	return elem.ListItem(
		vecty.ClassMap{
			"completed": i.completed,
			"editing":   i.editing,
		},

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(i.completed),
				event.Change(i.onToggleCompleted),
			),
			elem.Label(
				vecty.Text(i.title),
				event.DoubleClick(i.onStartEdit),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(i.onDestroy),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(i.onStopEdit).PreventDefault(),
			i.input,
		),
	)
}
