package model

import "github.com/neelance/dom/bind"

type Model struct {
	Scope *bind.Scope

	Items []*Item

	// derived
	IncompleteItemCount func() int
	CompletedItemCount  func() int
}

type Item struct {
	Text      string
	Completed bool
}
