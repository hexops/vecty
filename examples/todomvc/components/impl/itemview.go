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

func (p *ItemViewImpl) Render() dom.Spec {
	return elem.ListItem(
		dom.If(p.Props().Item().Completed,
			prop.Class("completed"),
		),
		// bind.IfFunc(editing, item.Scope,
		// 	prop.Class("editing"),
		// ),

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
				// event.DblClick(l.StartEdit(item)),
			),
			elem.Button(
				prop.Class("destroy"),
				event.Click(p.onDestroy),
			),
		),
		elem.Form(
			style.Margin(style.Px(0)),
			// dom.PreventDefault(event.Submit(l.StopEdit)),
			elem.Input(
				prop.Class("edit"),
				prop.Value(p.Props().Item().Title),
			),
		),
	)
}
