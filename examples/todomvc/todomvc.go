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
	m.IncompleteItemCount = count(false)
	m.CompletedItemCount = count(true)

	m.AddItem = func(c *dom.EventContext) {
		m.Items = append(m.Items, &model.Item{Title: m.AddItemTitle, Completed: false})
		m.AddItemTitle = ""
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
