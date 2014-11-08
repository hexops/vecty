package model

import "github.com/neelance/dom/bind"

type Model struct {
	Scope *bind.Scope

	Items []*Item

	// derived
	IncompleteItemCount func() string
	CompletedItemCount  func() string
}

type Item struct {
	Text      string
	Completed bool
}
