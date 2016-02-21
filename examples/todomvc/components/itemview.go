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
	vecty.Composite

	Index     int
	Item      *model.Item
	editing   bool
	editTitle string
}

// Apply implements the vecty.Markup interface.
func (p *ItemView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *ItemView) Reconcile(oldComp vecty.Component) {
	if oldComp, ok := oldComp.(*ItemView); ok {
		p.Body = oldComp.Body
		p.editing = oldComp.editing
		p.editTitle = oldComp.editTitle
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
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
	p.ReconcileBody()
}

func (p *ItemView) onEditInput(event *vecty.Event) {
	p.editTitle = event.Target.Get("value").String()
	p.ReconcileBody()
}

func (p *ItemView) onStopEdit(event *vecty.Event) {
	p.editing = false
	p.ReconcileBody()
	dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Index,
		Title: p.editTitle,
	})
}

func (p *ItemView) render() vecty.Component {
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
			elem.Input(
				prop.Class("edit"),
				prop.Value(p.editTitle),
				event.Input(p.onEditInput),
			),
		),
	)
}
