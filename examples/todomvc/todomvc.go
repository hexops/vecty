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
	dom.SetTitle("GopherJS • TodoMVC")
	dom.AddStylesheet("bower_components/todomvc-common/base.css")
	dom.SetBody(view.Page(createModel()))
}

func createModel() *model.Model {
	m := &model.Model{
		Scope: bind.NewScope(),
	}

	attachLocalStorage(m)

	m.AddItem = func(c *dom.EventContext) {
		m.Items = append(m.Items, &model.Item{Title: m.AddItemTitle, Completed: false})
		m.AddItemTitle = ""
		m.Scope.Digest()
	}

	m.DestroyItem = func(item *model.Item) dom.Listener {
		return func(c *dom.EventContext) {
			i := itemIndex(item, m)
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
		m.Scope.Digest()
	}

	m.ToggleAll = func(c *dom.EventContext) {
		checked := c.Node.Get("checked").Bool()
		for _, item := range m.Items {
			item.Completed = checked
		}
		m.Scope.Digest()
	}

	m.ActiveItemCount = countFunc(m, false)
	m.CompletedItemCount = countFunc(m, true)

	return m
}

func countFunc(m *model.Model, completed bool) func() int {
	return m.Scope.CacheInt(func() int {
		count := 0
		for _, item := range m.Items {
			if item.Completed == completed {
				count++
			}
		}
		return count
	})
}

func itemIndex(item *model.Item, m *model.Model) int {
	for i, item2 := range m.Items {
		if item == item2 {
			return i
		}
	}
	panic("item not found")
}

func attachLocalStorage(m *model.Model) {
	if data := js.Global.Get("localStorage").Get("items"); data != js.Undefined {
		if err := json.Unmarshal([]byte(data.String()), &m.Items); err != nil {
			fmt.Printf("failed to load items: %s\n", err)
		}
	}

	m.Scope.NewListener(func() {
		data, err := json.Marshal(&m.Items)
		if err != nil {
			fmt.Printf("failed to store items: %s\n", err)
		}
		js.Global.Get("localStorage").Set("items", string(data))
	})
}
