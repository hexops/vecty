package model

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
)

type Model struct {
	Scope *bind.Scope

	Items        []*Item
	AddItemTitle string

	// derived
	IncompleteItemCount func() int
	CompletedItemCount  func() int

	// listeners
	AddItem   dom.Listener
	ToggleAll dom.Listener
}

type Item struct {
	Title     string
	Completed bool
}
