package actions

import "github.com/neelance/dom/examples/todomvc/store/model"

type AddItem struct {
	Title string
}

type DestroyItem struct {
	Index int
}

type SetCompleted struct {
	Index     int
	Completed bool
}

type SetAllCompleted struct {
	Completed bool
}

type ClearCompleted struct{}

type SetFilter struct {
	Filter model.FilterState
}
