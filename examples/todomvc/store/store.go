package store

import (
	"github.com/neelance/dom/examples/todomvc/actions"
	"github.com/neelance/dom/examples/todomvc/dispatcher"
	"github.com/neelance/dom/examples/todomvc/store/model"
	"github.com/neelance/dom/storeutil"
)

var (
	Items     []*model.Item
	Filter    model.FilterState = model.All
	Listeners                   = storeutil.NewListenerRegistry()
)

func init() {
	dispatcher.Register(onAction)
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

func ItemIndex(item *model.Item) int {
	for i, item2 := range Items {
		if item == item2 {
			return i
		}
	}
	panic("item not found")
}

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.ReplaceItems:
		Items = a.Items
		Listeners.Fire()

	case *actions.AddItem:
		Items = append(Items, &model.Item{Title: a.Title, Completed: false})
		Listeners.Fire()

	case *actions.DestroyItem:
		copy(Items[a.Index:], Items[a.Index+1:])
		Items = Items[:len(Items)-1]
		Listeners.Fire()

	case *actions.SetTitle:
		Items[a.Index].Title = a.Title
		Listeners.Fire()

	case *actions.SetCompleted:
		Items[a.Index].Completed = a.Completed
		Listeners.Fire()

	case *actions.SetAllCompleted:
		for _, item := range Items {
			item.Completed = a.Completed
		}
		Listeners.Fire()

	case *actions.ClearCompleted:
		var activeItems []*model.Item
		for _, item := range Items {
			if !item.Completed {
				activeItems = append(activeItems, item)
			}
		}
		Items = activeItems
		Listeners.Fire()

	case *actions.SetFilter:
		Filter = a.Filter
		Listeners.Fire()

	}
}
