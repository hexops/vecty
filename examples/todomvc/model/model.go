package model

import "github.com/neelance/dom/bind"

type FilterState int

const (
	All FilterState = iota
	Active
	Completed
)

type ItemList struct {
	Scope *bind.Scope

	Items        []*Item
	AddItemTitle string
	EditItem     *Item
	Filter       FilterState
}

type Item struct {
	Scope *bind.Scope `json:"-"`

	Title     string
	Completed bool
}

func (m *ItemList) ActiveItemCount() int {
	return m.count(false)
}

func (m *ItemList) CompletedItemCount() int {
	return m.count(true)
}

func (m *ItemList) count(completed bool) int {
	count := 0
	for _, item := range m.Items {
		if item.Completed == completed {
			count++
		}
	}
	return count
}

func (m *ItemList) ItemIndex(item *Item) int {
	for i, item2 := range m.Items {
		if item == item2 {
			return i
		}
	}
	panic("item not found")
}
