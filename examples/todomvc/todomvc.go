package main

import (
	"encoding/json"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/examples/todomvc/model"
	"github.com/neelance/dom/examples/todomvc/view"

	"github.com/neelance/dom"
)

func main() {
	m := &model.ItemList{
		Scope: bind.NewScope(),
	}

	l := createListeners(m)

	attachLocalStorage(m)

	dom.SetTitle("GopherJS • TodoMVC")
	dom.AddStylesheet("bower_components/todomvc-common/base.css")
	dom.SetBody(view.Page(m, l))
}

func createListeners(m *model.ItemList) *view.PageListeners {
	l := &view.PageListeners{}

	l.AddItem = func(c *dom.EventContext) {
		m.Items = append(m.Items, &model.Item{Scope: m.Scope, Title: m.AddItemTitle, Completed: false})
		m.AddItemTitle = ""
		m.Scope.Digest()
	}

	l.DestroyItem = func(item *model.Item) dom.Listener {
		return func(c *dom.EventContext) {
			i := m.ItemIndex(item)
			copy(m.Items[i:], m.Items[i+1:])
			m.Items = m.Items[:len(m.Items)-1]
			m.Scope.Digest()
		}
	}

	l.StartEdit = func(item *model.Item) dom.Listener {
		return func(c *dom.EventContext) {
			m.EditItem = item
			m.Scope.Digest()
		}
	}

	l.StopEdit = func(c *dom.EventContext) {
		m.EditItem = nil
		m.Scope.Digest()
	}

	l.ClearCompleted = func(c *dom.EventContext) {
		var incomplete []*model.Item
		for _, item := range m.Items {
			if !item.Completed {
				incomplete = append(incomplete, item)
			}
		}
		m.Items = incomplete
		m.Scope.Digest()
	}

	l.ToggleAll = func(c *dom.EventContext) {
		checked := c.Node.Get("checked").Bool()
		for _, item := range m.Items {
			item.Completed = checked
		}
		m.Scope.Digest()
	}

	return l
}

func attachLocalStorage(m *model.ItemList) {
	if data := js.Global.Get("localStorage").Get("items"); data != js.Undefined {
		if err := json.Unmarshal([]byte(data.String()), &m.Items); err != nil {
			fmt.Printf("failed to load items: %s\n", err)
		}
	}
	for _, item := range m.Items {
		item.Scope = m.Scope
	}

	m.Scope.NewListener(func() {
		data, err := json.Marshal(&m.Items)
		if err != nil {
			fmt.Printf("failed to store items: %s\n", err)
		}
		js.Global.Get("localStorage").Set("items", string(data))
	})
}
