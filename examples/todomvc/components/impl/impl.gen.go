// GENERATED, DO NOT CHANGE

package impl

import (
	"github.com/gopherjs/gopherjs/js"

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
	Filter() model.FilterState
}

type FilterButtonState interface {
}

type filterButtonCore struct {
	componentutil.Core

	_label  string
	_filter model.FilterState
}

func (c *filterButtonCore) Props() FilterButtonProps {
	return c
}

func (c *filterButtonCore) Label() string {
	return c._label
}

func (c *filterButtonCore) Filter() model.FilterState {
	return c._filter
}

func (c *filterButtonCore) State() FilterButtonState {
	return c
}

func (c *filterButtonCore) applyProps(spec *spec.FilterButton) {
	c._label = spec.Label
	c._filter = spec.Filter
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
	Item() *model.Item
}

type ItemViewState interface {
	Editing() bool
	SetEditing(editing bool)
	EditTitle() string
	SetEditTitle(editTitle string)
}

type itemViewCore struct {
	componentutil.Core

	_index int
	_item  *model.Item

	_editing   bool
	_editTitle string
}

func (c *itemViewCore) Props() ItemViewProps {
	return c
}

func (c *itemViewCore) Index() int {
	return c._index
}

func (c *itemViewCore) Item() *model.Item {
	return c._item
}

func (c *itemViewCore) State() ItemViewState {
	return c
}

func (c *itemViewCore) Editing() bool {
	return c._editing
}

func (c *itemViewCore) SetEditing(editing bool) {
	c._editing = editing
	c.Update()
}

func (c *itemViewCore) EditTitle() string {
	return c._editTitle
}

func (c *itemViewCore) SetEditTitle(editTitle string) {
	c._editTitle = editTitle
	c.Update()
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
	Items() []*model.Item
	SetItems(items []*model.Item)
	NewItemTitle() string
	SetNewItemTitle(newItemTitle string)
}

type pageViewCore struct {
	componentutil.Core

	_items        []*model.Item
	_newItemTitle string
}

func (c *pageViewCore) Props() PageViewProps {
	return c
}

func (c *pageViewCore) State() PageViewState {
	return c
}

func (c *pageViewCore) Items() []*model.Item {
	return c._items
}

func (c *pageViewCore) SetItems(items []*model.Item) {
	c._items = items
	c.Update()
}

func (c *pageViewCore) NewItemTitle() string {
	return c._newItemTitle
}

func (c *pageViewCore) SetNewItemTitle(newItemTitle string) {
	c._newItemTitle = newItemTitle
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
