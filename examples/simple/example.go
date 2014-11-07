package main

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/examples/simple/model"
	"github.com/neelance/dom/examples/simple/view"
)

func main() {
	m := &model.Model{
		Scope: bind.NewScope(),
		Items: []*model.Item{
			&model.Item{"First item"},
			&model.Item{"Second item"},
			&model.Item{"Third item"},
		},
	}

	m.AppendItem = func() {
		m.Items = append(m.Items, &model.Item{"New item"})
		m.Scope.Digest()
	}

	m.PrependItem = func() {
		m.Items = append([]*model.Item{{"New item"}}, m.Items...)
		m.Scope.Digest()
	}

	itemIndex := func(item *model.Item) int {
		for i, item2 := range m.Items {
			if item == item2 {
				return i
			}
		}
		panic("item not found")
	}

	m.DeleteItem = func(item *model.Item) func() {
		return func() {
			i := itemIndex(item)
			copy(m.Items[i:], m.Items[i+1:])
			m.Items = m.Items[:len(m.Items)-1]
			m.Scope.Digest()
		}
	}

	view.Page(m).Apply(dom.Body())
}
