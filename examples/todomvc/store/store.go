package store

import (
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/storeutil"
)

type FilterState int

const (
	All FilterState = iota
	Active
	Completed
)

var (
	Items     []*Item
	Filter    FilterState = All
	Listeners             = storeutil.NewListenerRegistry()
)

func init() {
	Items = []*Item{
		{"Foo", false},
		{"Bar", true},
		{"Baz", false},
	}
	dispatcher.Register(onAction)
}

type Item struct {
	Title     string
	Completed bool
}

func ActiveItemCount() int {
	return count(false)
}

func CompletedItemCount() int {
	return count(true)
}

func count(completed bool) int {
	count := 0
	for _, item := range Items {
		if item.Completed == completed {
			count++
		}
	}
	return count
}

func ItemIndex(item *Item) int {
	for i, item2 := range Items {
		if item == item2 {
			return i
		}
	}
	panic("item not found")
}

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.SetCompleted:
		Items[a.Index].Completed = a.Completed
		Listeners.Fire()

	case *actions.SetAllCompleted:
		for _, item := range Items {
			item.Completed = a.Completed
		}
		Listeners.Fire()
	}
}
