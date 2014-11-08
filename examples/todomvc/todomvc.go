package main

import (
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/examples/todomvc/view"

	"github.com/neelance/dom"
)

func main() {
	model := &model.Model{
		Scope: bind.NewScope(),
		Items: []*model.Item{
			&model.Item{Text: "Create a TodoMVC template", Completed: true},
			&model.Item{Text: "Rule the web", Completed: false},
		},
	}

	view.Page(model).Apply(dom.Body())
}
