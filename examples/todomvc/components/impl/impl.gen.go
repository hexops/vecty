// GENERATED, DO NOT CHANGE

package impl

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/neelance/dom"

	"github.com/neelance/dom/componentutil"

	"github.com/neelance/dom/examples/todomvc/components/spec"

	"github.com/neelance/dom/examples/todomvc/store"
)

func init() {
	spec.ReconcileFilterButton = reconcileFilterButton
	spec.ReconcileItemView = reconcileItemView
	spec.ReconcilePageView = reconcilePageView
}

type FilterButtonImpl struct {
	FilterButtonAccessors
	componentutil.EmptyLifecycle
}

type FilterButtonAccessors interface {
	Props() FilterButtonProps
	State() FilterButtonState
	Node() *js.Object
}

type FilterButtonProps interface {
	Label() string
	Selected() bool
}

type FilterButtonState interface {
}

type filterButtonCore struct {
	componentutil.Core

	_label    string
	_selected bool
}

func (c *filterButtonCore) Props() FilterButtonProps {
	return c
}

func (c *filterButtonCore) Label() string {
	return c._label
}

func (c *filterButtonCore) Selected() bool {
	return c._selected
}

func (c *filterButtonCore) State() FilterButtonState {
	return c
}

func (c *filterButtonCore) applyProps(spec *spec.FilterButton) {
	c._label = spec.Label
	c._selected = spec.Selected
	c.DoRender()
}

func reconcileFilterButton(newSpec *spec.FilterButton, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.FilterButton); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*FilterButtonImpl).FilterButtonAccessors.(*filterButtonCore).applyProps(newSpec)
		return
	}

	c := &filterButtonCore{}
	inst := &FilterButtonImpl{FilterButtonAccessors: c}
	c.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	c.applyProps(newSpec)
	inst.ComponentDidMount()
}

type ItemViewImpl struct {
	ItemViewAccessors
	componentutil.EmptyLifecycle
}

type ItemViewAccessors interface {
	Props() ItemViewProps
	State() ItemViewState
	Node() *js.Object
}

type ItemViewProps interface {
	Index() int
	Item() *store.Item
}

type ItemViewState interface {
}

type itemViewCore struct {
	componentutil.Core

	_index int
	_item  *store.Item
}

func (c *itemViewCore) Props() ItemViewProps {
	return c
}

func (c *itemViewCore) Index() int {
	return c._index
}

func (c *itemViewCore) Item() *store.Item {
	return c._item
}

func (c *itemViewCore) State() ItemViewState {
	return c
}

func (c *itemViewCore) applyProps(spec *spec.ItemView) {
	c._index = spec.Index
	c._item = spec.Item
	c.DoRender()
}

func reconcileItemView(newSpec *spec.ItemView, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.ItemView); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*ItemViewImpl).ItemViewAccessors.(*itemViewCore).applyProps(newSpec)
		return
	}

	c := &itemViewCore{}
	inst := &ItemViewImpl{ItemViewAccessors: c}
	c.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	c.applyProps(newSpec)
	inst.ComponentDidMount()
}

type PageViewImpl struct {
	PageViewAccessors
	componentutil.EmptyLifecycle
}

type PageViewAccessors interface {
	Props() PageViewProps
	State() PageViewState
	Node() *js.Object
}

type PageViewProps interface {
}

type PageViewState interface {
	Items() []*store.Item
	SetItems(items []*store.Item)
}

type pageViewCore struct {
	componentutil.Core

	_items []*store.Item
}

func (c *pageViewCore) Props() PageViewProps {
	return c
}

func (c *pageViewCore) State() PageViewState {
	return c
}

func (c *pageViewCore) Items() []*store.Item {
	return c._items
}

func (c *pageViewCore) SetItems(items []*store.Item) {
	c._items = items
	c.Update()
}

func (c *pageViewCore) applyProps(spec *spec.PageView) {
	c.DoRender()
}

func reconcilePageView(newSpec *spec.PageView, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.PageView); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*PageViewImpl).PageViewAccessors.(*pageViewCore).applyProps(newSpec)
		return
	}

	c := &pageViewCore{}
	inst := &PageViewImpl{PageViewAccessors: c}
	c.Lifecycle = inst
	newSpec.Instance = inst
	inst.ComponentWillMount()
	c.applyProps(newSpec)
	inst.ComponentDidMount()
}
