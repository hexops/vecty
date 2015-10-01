package spec

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/examples/todomvc/store/model"
)

type PageView struct {
	dom.Instance
	items        []*model.Item
	newItemTitle string
}

type ItemView struct {
	dom.Instance
	Index     int
	Item      *model.Item
	editing   bool
	editTitle string
}

type FilterButton struct {
	dom.Instance
	Label  string
	Filter model.FilterState
}

func (s *PageView) Apply(element *dom.Element)     { element.AddChild(s) }
func (s *ItemView) Apply(element *dom.Element)     { element.AddChild(s) }
func (s *FilterButton) Apply(element *dom.Element) { element.AddChild(s) }

func (s *PageView) Reconcile(oldSpec dom.Spec)     { ReconcilePageView(s, oldSpec) }
func (s *ItemView) Reconcile(oldSpec dom.Spec)     { ReconcileItemView(s, oldSpec) }
func (s *FilterButton) Reconcile(oldSpec dom.Spec) { ReconcileFilterButton(s, oldSpec) }

var ReconcilePageView func(newSpec *PageView, oldSpec dom.Spec)
var ReconcileItemView func(newSpec *ItemView, oldSpec dom.Spec)
var ReconcileFilterButton func(newSpec *FilterButton, oldSpec dom.Spec)
