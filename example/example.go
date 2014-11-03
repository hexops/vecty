package main

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/prop"
	"github.com/neelance/dom/style"
)

type Item struct {
	Label string
}

var name string
var items []*Item
var scope = bind.NewScope()

func main() {
	items = []*Item{
		&Item{"First item"},
		&Item{"Second item"},
		&Item{"Third item"},
	}

	PageView().Apply(dom.Body())
}

func PageView() *dom.ElemAspect {
	return elem.Div(
		GreetingView(),
		ItemsView(),
	)
}

func GreetingView() *dom.ElemAspect {
	return elem.Div(
		dom.Text("Your name: "),
		elem.Input(
			prop.Type("text"),
			bind.InputValue(&name, scope),
		),
		elem.H1(
			style.Color("blue"),
			dom.Text("Hello "),
			bind.Text(&name, scope),
			dom.Text("!"),
		),
	)
}

func ItemsView() *dom.ElemAspect {
	return elem.Div(
		elem.Ul(
			bind.Dynamic(scope, func(aspects *bind.Aspects) {
				for _, item := range items {
					if !aspects.Reuse(item) {
						aspects.Add(item, elem.Li(
							bind.Text(&item.Label, scope),
							dom.Text(" "),
							elem.Button(
								dom.Text("Delete"),
								event.Click(DeleteItem(item)),
							),
						))
					}
				}
			}),
		),

		elem.Button(
			dom.Text("Append item"),
			event.Click(AppendItem),
		),
		elem.Button(
			dom.Text("Prepend item"),
			event.Click(PrependItem),
		),
	)
}

func AppendItem() {
	items = append(items, &Item{"New item"})
	scope.Digest()
}

func PrependItem() {
	items = append([]*Item{{"New item"}}, items...)
	scope.Digest()
}

func DeleteItem(item *Item) func() {
	return func() {
		i := ItemIndex(item)
		copy(items[i:], items[i+1:])
		items = items[:len(items)-1]
		scope.Digest()
	}
}

func ItemIndex(item *Item) int {
	for i, item2 := range items {
		if item == item2 {
			return i
		}
	}
	panic("item not found")
}
