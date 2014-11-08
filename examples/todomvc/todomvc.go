package main

import (
	"strconv"

	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/examples/todomvc/view"

	"github.com/neelance/dom"
)

func main() {
	scope := bind.NewScope()

	model := &model.Model{
		Scope: scope,
		Items: []*model.Item{
			&model.Item{Text: "Create a TodoMVC template", Completed: true},
			&model.Item{Text: "Rule the web", Completed: false},
		},
	}

	count := func(completed bool) func() string {
		return scope.CacheString(func() string {
			count := 0
			for _, item := range model.Items {
				if item.Completed == completed {
					count++
				}
			}
			return strconv.Itoa(count)
		})
	}
	model.IncompleteItemCount = count(false)
	model.CompletedItemCount = count(true)

	view.Page(model).Apply(dom.Body())
}
