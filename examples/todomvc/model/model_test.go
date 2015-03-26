package model

import "testing"

func TestActiveItemCount(t *testing.T) {
	l := ItemList{
		Items: []*Item{
			&Item{Completed: false},
			&Item{Completed: true},
			&Item{Completed: false},
		},
	}
	if l.ActiveItemCount() != 2 {
		t.Fail()
	}
}

func TestCompletedItemCount(t *testing.T) {
	l := ItemList{
		Items: []*Item{
			&Item{Completed: true},
			&Item{Completed: false},
			&Item{Completed: true},
		},
	}
	if l.CompletedItemCount() != 2 {
		t.Fail()
	}
}

func TestItemIndex(t *testing.T) {
	item := &Item{}
	l := ItemList{
		Items: []*Item{
			&Item{},
			item,
			&Item{},
		},
	}
	if l.ItemIndex(item) != 1 {
		t.Fail()
	}
}
