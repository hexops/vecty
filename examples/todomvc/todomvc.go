package main

import (
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/examples/todomvc/view"

	"github.com/neelance/dom"
)

func main() {
	scope := bind.NewScope()

	m := &model.Model{
		Scope: scope,
		Items: []*model.Item{
			&model.Item{Title: "Create a TodoMVC template", Completed: true},
			&model.Item{Title: "Rule the web", Completed: false},
		},
	}

	count := func(completed bool) func() int {
		return scope.CacheInt(func() int {
			count := 0
			for _, item := range m.Items {
				if item.Completed == completed {
					count++
				}
			}
			return count
		})
	}
	m.ActiveItemCount = count(false)
	m.CompletedItemCount = count(true)

	itemIndex := func(item *model.Item) int {
		for i, item2 := range m.Items {
			if item == item2 {
				return i
			}
		}
		panic("item not found")
	}

	m.AddItem = func(c *dom.EventContext) {
		m.Items = append(m.Items, &model.Item{Title: m.AddItemTitle, Completed: false})
		m.AddItemTitle = ""
		scope.Digest()
	}

	m.DestroyItem = func(item *model.Item) dom.Listener {
		return func(c *dom.EventContext) {
			i := itemIndex(item)
			copy(m.Items[i:], m.Items[i+1:])
			m.Items = m.Items[:len(m.Items)-1]
			m.Scope.Digest()
		}
	}

	m.ClearCompleted = func(c *dom.EventContext) {
		var incomplete []*model.Item
		for _, item := range m.Items {
			if !item.Completed {
				incomplete = append(incomplete, item)
			}
		}
		m.Items = incomplete
		scope.Digest()
	}

	m.ToggleAll = func(c *dom.EventContext) {
		checked := c.Node.Get("checked").Bool()
		for _, item := range m.Items {
			item.Completed = checked
		}
		scope.Digest()
	}

	view.Page(m).Apply(dom.Body())
}
