package actions

import "github.com/gopherjs/vecty/example/todomvc/store/model"

// ReplaceItems is an action that replaces all items with the specified ones.
type ReplaceItems struct {
	Items []*model.Item
}

// AddItem is an action which adds a single item with the specified title.
type AddItem struct {
	Title string
}

// DestroyItem is an action which destroys the item specified by the index.
type DestroyItem struct {
	Index int
}

// SetTitle is an action which specifies the title of an existing item.
type SetTitle struct {
	Index int
	Title string
}

// SetCompleted is an action which specifies the completed state of an existing
// item.
type SetCompleted struct {
	Index     int
	Completed bool
}

// SetAllCompleted is an action which marks all existing items as being
// completed or not.
type SetAllCompleted struct {
	Completed bool
}

// ClearCompleted is an action which clears the completed items.
type ClearCompleted struct{}

// SetFilter is an action which sets the filter for the viewed items.
type SetFilter struct {
	Filter model.FilterState
}
