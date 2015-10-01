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
		Index: p.Props().Index(),
	})
}

func (p *ItemViewImpl) onToggleCompleted(event *dom.Event) {
	dispatcher.Dispatch(&actions.SetCompleted{
		Index:     p.Props().Index(),
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *ItemViewImpl) onStartEdit(event *dom.Event) {
	p.State().SetEditing(true)
	p.State().SetEditTitle(p.Props().Item().Title)
}

func (p *ItemViewImpl) onEditInput(event *dom.Event) {
	p.State().SetEditTitle(event.Target.Get("value").String())
}

func (p *ItemViewImpl) onStopEdit(event *dom.Event) {
	p.State().SetEditing(false)
	dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Props().Index(),
		Title: p.State().EditTitle(),
	})
}

func (p *ItemViewImpl) Render() dom.Spec {
	return elem.ListItem(
		dom.ClassMap{
			"completed": p.Props().Item().Completed,
			"editing":   p.State().Editing(),
		},

		elem.Div(
			prop.Class("view"),

			elem.Input(
				prop.Class("toggle"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(p.Props().Item().Completed),
				event.Change(p.onToggleCompleted),
			),
			elem.Label(
				dom.Text(p.Props().Item().Title),
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
				prop.Value(p.State().EditTitle()),
				event.Input(p.onEditInput),
			),
		),
	)
}
