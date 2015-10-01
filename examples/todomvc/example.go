package main

import (
	"encoding/json"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/examples/todomvc/actions"
	_ "github.com/neelance/dom/examples/todomvc/components/impl"
	"github.com/neelance/dom/examples/todomvc/components/spec"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store"
	"github.com/neelance/dom/examples/todomvc/store/model"
)

func main() {
	attachLocalStorage()

	dom.SetTitle("GopherJS • TodoMVC")
	dom.AddStylesheet("assets/base.css")
	dom.RenderAsBody(&spec.PageView{})
}

func attachLocalStorage() {
	store.Listeners.Add(nil, func() {
		data, err := json.Marshal(store.Items)
		if err != nil {
			fmt.Printf("failed to store items: %s\n", err)
		}
		js.Global.Get("localStorage").Set("items", string(data))
	})

	if data := js.Global.Get("localStorage").Get("items"); data != js.Undefined {
		var items []*model.Item
		if err := json.Unmarshal([]byte(data.String()), &items); err != nil {
			fmt.Printf("failed to load items: %s\n", err)
		}
		dispatcher.Dispatch(&actions.ReplaceItems{
			Items: items,
		})
	}
}
