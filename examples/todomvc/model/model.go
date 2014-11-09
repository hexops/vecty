package model

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
)

type FilterState int

const (
	All FilterState = iota
	Active
	Completed
)

type Model struct {
	Scope *bind.Scope

	Items        []*Item
	AddItemTitle string
	EditItem     *Item
	Filter       FilterState

	// derived
	ActiveItemCount    func() int
	CompletedItemCount func() int

	// listeners
	AddItem        dom.Listener
	DestroyItem    func(*Item) dom.Listener
	ClearCompleted dom.Listener
	ToggleAll      dom.Listener
}

type Item struct {
	Title     string
	Completed bool
}
