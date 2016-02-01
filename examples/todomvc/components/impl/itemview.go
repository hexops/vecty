package impl

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

func (p *ItemViewImpl) onDestroy(event *dom.Event) {
	dispatcher.Dispatch(&actions.DestroyItem{
		Index: p.Index,
	})
}

func (p *ItemViewImpl) onToggleCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetCompleted{
		Index:     p.Index,
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *ItemViewImpl) onStartEdit(event *dom.Event) {
	p.editing = true
	p.editTitle = p.Item.Title
	p.Update()
}

func (p *ItemViewImpl) onEditInput(event *dom.Event) {
	p.editTitle = event.Target.Get("value").String()
	p.Update()
}

func (p *ItemViewImpl) onStopEdit(event *dom.Event) {
	p.editing = false
	p.Update()
	dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Index,
		Title: p.editTitle,
	})
}

func (p *ItemViewImpl) Render() dom.Spec {
	return elem.ListItem(
		dom.ClassMap{
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
				dom.Text(p.Item.Title),
				event.DblClick(p.onStartEdit),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(p.onDestroy),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			event.Submit(p.onStopEdit).PreventDefault(),
			elem.Input(
				prop.Class("edit"),
				prop.Value(p.editTitle),
				event.Input(p.onEditInput),
			),
		),
	)
}
