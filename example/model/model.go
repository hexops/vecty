package model

import "github.com/neelance/dom/bind"

type Item struct {
	Label string
}

type Model struct {
	Scope *bind.Scope

	Name  string
	Items []*Item

	AppendItem  func()
	PrependItem func()
	DeleteItem  func(*Item) func()
}
