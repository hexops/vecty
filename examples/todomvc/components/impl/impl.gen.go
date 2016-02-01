// GENERATED, DO NOT CHANGE

package impl

import (
	"github.com/neelance/dom"

	"github.com/neelance/dom/componentutil"

	"github.com/neelance/dom/examples/todomvc/components/spec"

	"github.com/neelance/dom/examples/todomvc/store/model"
)

func init() {
	spec.ReconcileFilterButton = reconcileFilterButton
	spec.ReconcileItemView = reconcileItemView
	spec.ReconcilePageView = reconcilePageView
}

type FilterButtonImpl struct {
	componentutil.Core
	componentutil.EmptyLifecycle

	Label  string
	Filter model.FilterState
}

func (c *FilterButtonImpl) applyProps(spec *spec.FilterButton) {
	c.Label = spec.Label
	c.Filter = spec.Filter
	c.DoRender()
}

func reconcileFilterButton(newSpec *spec.FilterButton, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.FilterButton); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*FilterButtonImpl).applyProps(newSpec)
		return
	}

	inst := &FilterButtonImpl{}
	inst.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	inst.applyProps(newSpec)
	inst.ComponentDidMount()
}

type ItemViewImpl struct {
	componentutil.Core
	componentutil.EmptyLifecycle

	Index int
	Item  *model.Item

	editing   bool
	editTitle string
}

func (c *ItemViewImpl) applyProps(spec *spec.ItemView) {
	c.Index = spec.Index
	c.Item = spec.Item
	c.DoRender()
}

func reconcileItemView(newSpec *spec.ItemView, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.ItemView); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*ItemViewImpl).applyProps(newSpec)
		return
	}

	inst := &ItemViewImpl{}
	inst.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	inst.applyProps(newSpec)
	inst.ComponentDidMount()
}

type PageViewImpl struct {
	componentutil.Core
	componentutil.EmptyLifecycle

	items        []*model.Item
	newItemTitle string
}

func (c *PageViewImpl) applyProps(spec *spec.PageView) {
	c.DoRender()
}

func reconcilePageView(newSpec *spec.PageView, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.PageView); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*PageViewImpl).applyProps(newSpec)
		return
	}

	inst := &PageViewImpl{}
	inst.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	inst.applyProps(newSpec)
	inst.ComponentDidMount()
}
