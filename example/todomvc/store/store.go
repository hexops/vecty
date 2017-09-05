package store

import (
	"github.com/gopherjs/vecty/example/todomvc/actions"
	"github.com/gopherjs/vecty/example/todomvc/dispatcher"
	"github.com/gopherjs/vecty/example/todomvc/store/model"
	"github.com/gopherjs/vecty/example/todomvc/store/storeutil"
)

var (
	// Items represents all of the TODO items in the store.
	Items []*model.Item

	// Filter represents the active viewing filter for items.
	Filter = model.All

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()
)

func init() {
	dispatcher.Register(onAction)
}

// ActiveItemCount returns the current number of items that are not completed.
func ActiveItemCount() int {
	return count(false)
}

// CompletedItemCount returns the current number of items that are completed.
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

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.ReplaceItems:
		Items = a.Items

	case *actions.AddItem:
		Items = append(Items, &model.Item{Title: a.Title, Completed: false})

	case *actions.DestroyItem:
		copy(Items[a.Index:], Items[a.Index+1:])
		Items = Items[:len(Items)-1]

	case *actions.SetTitle:
		Items[a.Index].Title = a.Title

	case *actions.SetCompleted:
		Items[a.Index].Completed = a.Completed

	case *actions.SetAllCompleted:
		for _, item := range Items {
			item.Completed = a.Completed
		}

	case *actions.ClearCompleted:
		var activeItems []*model.Item
		for _, item := range Items {
			if !item.Completed {
				activeItems = append(activeItems, item)
			}
		}
		Items = activeItems

	case *actions.SetFilter:
		Filter = a.Filter

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}
