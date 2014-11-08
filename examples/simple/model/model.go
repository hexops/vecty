package model

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
)

type Item struct {
	Label string
}

type Model struct {
	Scope *bind.Scope

	Name  string
	Items []*Item

	AppendItem  dom.Listener
	PrependItem dom.Listener
	DeleteItem  func(*Item) dom.Listener
}
