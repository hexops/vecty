package actions

import "github.com/gopherjs/vecty/examples/todomvc/store/model"

type ReplaceItems struct {
	Items []*model.Item
}

type AddItem struct {
	Title string
}

type DestroyItem struct {
	Index int
}

type SetTitle struct {
	Index int
	Title string
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
