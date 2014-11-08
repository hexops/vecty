package model

import "github.com/neelance/dom/bind"

type Model struct {
	Scope *bind.Scope

	Items []*Item
}

type Item struct {
	Text      string
	Completed bool
}
