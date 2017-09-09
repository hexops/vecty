package vecty

import (
	"fmt"
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

// Core implements the Context method of the Component interface, and is the
// core/central struct which all Component implementations should embed.
type Core struct {
	prevRenderComponent Component
	prevRender          *HTML
	unmounted           bool
}

// Context implements the Component interface.
func (c *Core) Context() *Core { return c }

// Component represents a single visual component within an application. To
// define a new component simply implement the Render method and embed the Core
// struct:
//
// 	type MyComponent struct {
// 		vecty.Core
// 		... additional component fields (state or properties) ...
// 	}
//
// 	func (c *MyComponent) Render() *vecty.HTML {
// 		... rendering ...
// 	}
//
type Component interface {
	// Render is responsible for building HTML which represents the component.
	//
	// If Render returns nil, the component will render as nothing (in reality,
	// a noscript tag, which has no display or action, and is compatible with
	// Vecty's diffing algorithm).
	Render() *HTML

	// Context returns the components context, which is used internally by
	// Vecty in order to store the previous component render for diffing.
	Context() *Core
}

// Copier is an optional interface that a Component can implement in order to
// copy itself. Vecty must internally copy components, and it does so by either
// invoking the Copy method of the Component or, if the component does not
// implement the Copier interface, a shallow copy is performed.
type Copier interface {
	// Copy returns a copy of the component.
	Copy() Component
}

// Mounter is an optional interface that a Component can implement in order
// to receive component mount events.
type Mounter interface {
	// Mount is called after the component has been mounted, after the DOM node
	// has been attached.
	Mount()
}

// Unmounter is an optional interface that a Component can implement in order
// to receive component unmount events.
type Unmounter interface {
	// Unmount is called before the component has been unmounted, before the
	// DOM node has been removed.
	Unmount()
}

// ComponentOrHTML represents one of:
//
//  Component
//  *HTML
//  List
//  nil
//
// If the underlying value is not one of these types, the code handling the
// value is expected to panic.
type ComponentOrHTML interface{}

// RenderSkipper is an optional interface that Component's can implement in
// order to short-circuit the reconciliation of a Component's rendered body.
//
// This is purely an optimization, and does not need to be implemented by
// Components for correctness. Without implementing this interface, only the
// difference between renders will be applied to the browser DOM. This
// interface allows components to bypass calculating the difference altogether
// and quickly state "nothing has changed, do not re-render".
type RenderSkipper interface {
	// SkipRender is called with a copy of the Component made the last time its
	// Render method was invoked. If it returns true, rendering of the
	// component will be skipped.
	//
	// The previous component may be of a different type than this
	// RenderSkipper itself, thus a type assertion should be used no action
	// taken if the type does not match.
	SkipRender(prev Component) bool
}

// HTML represents some form of HTML: an element with a specific tag, or some
// literal text (a TextNode).
type HTML struct {
	node jsObject

	namespace, tag, text, innerHTML string
	styles, dataset                 map[string]string
	properties, attributes          map[string]interface{}
	eventListeners                  []*EventListener
	children                        []ComponentOrHTML
}

// Node returns the underlying JavaScript Element or TextNode.
func (h *HTML) Node() *js.Object { return h.node.(wrappedObject).j }

func (h *HTML) createNode() {
	switch {
	case h.tag != "" && h.text != "":
		panic("vecty: only one of HTML.tag or HTML.text may be set")
	case h.tag == "" && h.innerHTML != "":
		panic("vecty: only HTML may have UnsafeHTML attribute")
	case h.tag != "" && h.namespace == "":
		h.node = global.Get("document").Call("createElement", h.tag)
	case h.tag != "" && h.namespace != "":
		h.node = global.Get("document").Call("createElementNS", h.namespace, h.tag)
	default:
		h.node = global.Get("document").Call("createTextNode", h.text)
	}
}

func (h *HTML) reconcileText(prev *HTML) {
	h.node = prev.node

	// Text modifications.
	if h.text != prev.text {
		h.node.Set("nodeValue", h.text)
	}
}

func (h *HTML) reconcile(prev *HTML) []Mounter {
	// Check for compatible tag and mutate previous instance on match, otherwise start fresh
	switch {
	case prev != nil && h.tag == "" && prev.tag == "":
		// Compatible text node
		h.reconcileText(prev)
		return nil
	case prev != nil && h.tag != "" && prev.tag != "" && h.tag == prev.tag && h.namespace == prev.namespace:
		// Compatible element node
		h.node = prev.node
	default:
		// Incompatible node, start fresh
		if prev == nil {
			prev = &HTML{}
		}
		h.createNode()
	}

	// Wrap event listeners
	for _, l := range h.eventListeners {
		l := l
		l.wrapper = func(jsEvent *js.Object) {
			if l.callPreventDefault {
				jsEvent.Call("preventDefault")
			}
			if l.callStopPropagation {
				jsEvent.Call("stopPropagation")
			}
			l.Listener(&Event{Object: jsEvent, Target: jsEvent.Get("target")})
		}
	}

	// Properties
	for name, value := range h.properties {
		var oldValue interface{}
		switch name {
		case "value":
			oldValue = h.node.Get("value").String()
		case "checked":
			oldValue = h.node.Get("checked").Bool()
		default:
			oldValue = prev.properties[name]
		}
		if value != oldValue {
			h.node.Set(name, value)
		}
	}
	new := h.node != prev.node
	if !new {
		for name := range prev.properties {
			if _, ok := h.properties[name]; !ok {
				h.node.Delete(name)
			}
		}
	}

	// Attributes
	for name, value := range h.attributes {
		if value != prev.attributes[name] {
			h.node.Call("setAttribute", name, value)
		}
	}
	if !new {
		for name := range prev.attributes {
			if _, ok := h.attributes[name]; !ok {
				h.node.Call("removeAttribute", name)
			}
		}
	}

	// Dataset
	dataset := h.node.Get("dataset")
	for name, value := range h.dataset {
		if value != prev.dataset[name] {
			dataset.Set(name, value)
		}
	}
	if !new {
		for name := range prev.dataset {
			if _, ok := h.dataset[name]; !ok {
				dataset.Delete(name)
			}
		}
	}

	// Styles
	style := h.node.Get("style")
	for name, value := range h.styles {
		oldValue := prev.styles[name]
		if value != oldValue {
			style.Call("setProperty", name, value)
		}
	}
	if !new {
		for name := range prev.styles {
			if _, ok := h.styles[name]; !ok {
				style.Call("removeProperty", name)
			}
		}
	}

	if !new {
		for _, l := range prev.eventListeners {
			h.node.Call("removeEventListener", l.Name, l.wrapper)
		}
	}
	for _, l := range h.eventListeners {
		h.node.Call("addEventListener", l.Name, l.wrapper)
	}

	if h.innerHTML != prev.innerHTML {
		h.node.Set("innerHTML", h.innerHTML)
	}

	return h.reconcileChildren(prev, nil)
}

// reconcileChildren reconciles children of the current HTML against a previous
// render's DOM nodes.
func (h *HTML) reconcileChildren(prev *HTML, insertBefore *HTML) (pendingMounts []Mounter) {
	// TODO better list element reuse
	for i, nextChild := range h.children {
		// TODO(pdf): Add tests for node equality
		new := h.node != prev.node
		if i >= len(prev.children) || new {
			if nextChildList, ok := nextChild.(List); ok {
				nextChildList.reconcile(h.node, insertBefore, nil)
				continue
			}
			nextChildRender, skip, mounters := render(nextChild, nil)
			if skip || nextChildRender == nil {
				continue
			}
			pendingMounts = append(pendingMounts, mounters...)
			if m, ok := nextChild.(Mounter); ok {
				pendingMounts = append(pendingMounts, m)
			}
			if insertBefore != nil {
				h.node.Call("insertBefore", nextChildRender.node, insertBefore.node)
				continue
			}
			h.node.Call("appendChild", nextChildRender.node)
			continue
		}

		prevChild := prev.children[i]
		var prevChildRender *HTML
		// If the previous child was not a list, extract the previous child
		// render.
		if _, isList := prevChild.(List); !isList {
			prevChildRender = extractHTML(prevChild)
		}

		// If the previous child render was nil, try to find the next DOM node
		// in the previous render so that we can insert this child at the
		// correct location.
		if prevChildRender == nil {
			for j := i + 1; j < len(prev.children) && insertBefore == nil; j++ {
				if prevChildList, ok := prev.children[j].(List); ok {
					insertBefore = prevChildList.firstHTML()
				} else {
					insertBefore = extractHTML(prev.children[j])
				}
			}
		}

		// If the next child is a list, reconcile its elements in-place, and
		// we're done.
		if nextChildList, ok := nextChild.(List); ok {
			nextChildList.reconcile(h.node, insertBefore, prevChild)
			continue
		}

		// If the previous child was a list, remove the list elements from the
		// previous render, since we no longer have a list.
		if prevChildList, ok := prevChild.(List); ok {
			prevChildList.remove(h.node)
		}

		// Determine the next child render.
		nextChildRender, skip, mounters := render(nextChild, prevChild)
		if nextChildRender != nil && prevChildRender != nil && nextChildRender == prevChildRender {
			panic("vecty: next child render must not equal previous child render (did the child Render illegally return a stored render variable?)")
		}

		// If the previous and next child are components of the same type, then
		// keep prevChildComponent as our child so that the next time we are
		// here prevChild will be the same pointer. We do this because
		// prevChildComponent is the persistent component instance.
		if prevChildComponent, ok := prevChild.(Component); ok {
			if nextChildComponent, ok := nextChild.(Component); ok && sameType(prevChildComponent, nextChildComponent) {
				h.children[i] = prevChild
				nextChild = prevChild
			}
		}
		if skip {
			continue
		}
		pendingMounts = append(pendingMounts, mounters...)

		// Perform the final reconciliation action for nextChildRender and
		// prevChildRender. Replace, remove, insert or append the DOM nodes.
		switch {
		case nextChildRender == nil && prevChildRender == nil:
			continue // nothing to do.
		case nextChildRender != nil && prevChildRender != nil:
			if m := mountUnmount(nextChild, prevChild); m != nil {
				pendingMounts = append(pendingMounts, m)
			}
			replaceNode(nextChildRender.node, prevChildRender.node)
		case nextChildRender == nil && prevChildRender != nil:
			unmount(prevChild)
			removeNode(prevChildRender.node)
		case nextChildRender != nil && prevChildRender == nil:
			if m, ok := nextChild.(Mounter); ok {
				pendingMounts = append(pendingMounts, m)
			}
			if insertBefore != nil {
				h.node.Call("insertBefore", nextChildRender.node, insertBefore.node)
				continue
			}
			h.node.Call("appendChild", nextChildRender.node)
		default:
			panic("vecty: internal error (unexpected switch state)")
		}
	}
	h.removeChildren(prev)
	return pendingMounts
}

// removeChildren removes child elements from the previous render pass that no
// longer exist on the current HTML children.
func (h *HTML) removeChildren(prev *HTML) {
	if prev == nil {
		return
	}
	// Every previous child that h.children does not have in common.
	for i := len(h.children); i < len(prev.children); i++ {
		prevChild := prev.children[i]
		if prevChildList, ok := prevChild.(List); ok {
			// Previous child was a list, so remove all DOM nodes in it.
			prevChildList.remove(h.node)
			continue
		}
		prevChildRender := extractHTML(prevChild)
		if prevChildRender == nil {
			continue
		}
		unmount(prevChild)
		removeNode(prevChildRender.node)
	}
}

// List represents a list of components or HTML.
type List []ComponentOrHTML

// reconcile reconciles the List against the DOM node in isolation. If
// insertBefore is non-nil, the List elements will be inserted into the DOM
// before that node.
func (l List) reconcile(node jsObject, insertBefore *HTML, prev ComponentOrHTML) []Mounter {
	nextHTML := &HTML{node: node, children: l}
	switch c := prev.(type) {
	case List:
		return nextHTML.reconcileChildren(&HTML{node: node, children: c}, insertBefore)
	default:
		if prev == nil {
			return nextHTML.reconcileChildren(&HTML{node: node}, insertBefore)
		}
		return nextHTML.reconcileChildren(&HTML{node: node, children: []ComponentOrHTML{prev}}, insertBefore)
	}
}

// firstHTML returns the content of the first element from which a *HTML can be
// extracted. Returns nil if not found.
func (l List) firstHTML() (h *HTML) {
	for _, v := range l {
		if listChild, ok := v.(List); ok {
			h = listChild.firstHTML()
		} else {
			h = extractHTML(v)
		}
		if h != nil {
			return h
		}
	}
	return nil
}

// remove the List's elements from the DOM node
func (l List) remove(node jsObject) {
	nextHTML := &HTML{node: node}
	nextHTML.removeChildren(&HTML{node: node, children: l})
}

// Tag returns an HTML element with the given tag name. Generally, this
// function is not used directly but rather the elem subpackage (which is type
// safe) is used instead.
func Tag(tag string, m ...MarkupOrChild) *HTML {
	h := &HTML{
		tag: tag,
	}
	for _, m := range m {
		apply(m, h)
	}
	return h
}

// Text returns a TextNode with the given literal text. Because the returned
// HTML represents a TextNode, the text does not have to be escaped (arbitrary
// user input fed into this function will always be safely rendered).
func Text(text string, m ...MarkupOrChild) *HTML {
	h := &HTML{
		text: text,
	}
	for _, m := range m {
		apply(m, h)
	}
	return h
}

// Rerender causes the body of the given component (i.e. the HTML returned by
// the Component's Render method) to be re-rendered.
//
// If the component has not been rendered before, Rerender panics. If the
// component was previously unmounted, Rerender is no-op.
func Rerender(c Component) {
	if c == nil {
		panic("vecty: Rerender illegally called with a nil Component argument")
	}
	prevRender := c.Context().prevRender
	if prevRender == nil {
		panic("vecty: Rerender invoked on Component that has never been rendered")
	}
	if c.Context().unmounted {
		return
	}
	nextRender, skip, pendingMounts := renderComponent(c, prevRender)
	if skip {
		return
	}
	// mountUnmount to check for node replacement
	mountUnmount(nextRender, prevRender)
	replaceNode(nextRender.node, prevRender.node)
	mount(pendingMounts...)
}

// extractHTML returns the *HTML from a ComponentOrHTML.
func extractHTML(e ComponentOrHTML) *HTML {
	switch v := e.(type) {
	case nil:
		return nil
	case *HTML:
		return v
	case Component:
		return v.Context().prevRender
	default:
		panic(fmt.Sprintf("vecty: encountered invalid ComponentOrHTML %T", e))
	}
}

// sameType returns whether first and second components are of the same type
func sameType(first, second Component) bool {
	return reflect.TypeOf(first) == reflect.TypeOf(second)
}

// doCopy makes a copy of the given component.
func doCopy(c Component) Component {
	if c == nil {
		panic("vecty: cannot copy nil Component")
	}

	// If the Component implements the Copier interface, then use that to
	// perform the copy.
	if copier, ok := c.(Copier); ok {
		cpy := copier.Copy()
		if cpy == c {
			panic("vecty: Component.Copy returned an identical *MyComponent pointer")
		}
		return cpy
	}

	// Component does not implement the Copier interface, so perform a shallow
	// copy.
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		panic(fmt.Sprintf("vecty: Component must be pointer to struct, found %T", c))
	}
	cpy := reflect.New(v.Elem().Type())
	cpy.Elem().Set(v.Elem())
	return cpy.Interface().(Component)
}

// copyProps copies all struct fields from src to dst that are tagged with
// `vecty:"prop"`.
//
// If src and dst are different types, copyProps is no-op.
func copyProps(src, dst Component) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dst)
	if s.Type() != d.Type() {
		return
	}
	for i := 0; i < s.Elem().NumField(); i++ {
		sf := s.Elem().Field(i)
		if s.Elem().Type().Field(i).Tag.Get("vecty") == "prop" {
			df := d.Elem().Field(i)
			if sf.Type() != df.Type() {
				panic("vecty: internal error (should never be possible, struct types are identical)")
			}
			df.Set(sf)
		}
	}
}

// render handles rendering the next child into HTML. If skip is returned,
// the component's SkipRender method has signaled to skip rendering.
//
// In specific, render handles six cases:
//
// 1. nextChild == *HTML && prevChild == *HTML
// 2. nextChild == *HTML && prevChild == Component
// 3. nextChild == *HTML && prevChild == nil
// 4. nextChild == Component && prevChild == Component
// 5. nextChild == Component && prevChild == *HTML
// 6. nextChild == Component && prevChild == nil
//
func render(next, prev ComponentOrHTML) (h *HTML, skip bool, pendingMounts []Mounter) {
	switch v := next.(type) {
	case *HTML:
		// Cases 1, 2 and 3 above. Reconcile against the prevRender.
		pendingMounts := v.reconcile(extractHTML(prev))
		return v, false, pendingMounts
	case Component:
		// Cases 4, 5, and 6 above.
		return renderComponent(v, prev)
	case nil:
		return nil, false, nil
	default:
		panic(fmt.Sprintf("vecty: encountered invalid ComponentOrHTML %T", next))
	}
}

// renderComponent handles rendering the given Component into *HTML. If skip ==
// true is returned, the Component's SkipRender method has signaled the
// component does not need to be rendered and h == nil is returned.
func renderComponent(next Component, prev ComponentOrHTML) (h *HTML, skip bool, pendingMounts []Mounter) {
	// If we had a component last render, and it's of compatible type, operate
	// on the previous instance.
	if prevComponent, ok := prev.(Component); ok && sameType(next, prevComponent) {
		// Copy `vecty:"prop"` fields from the newly rendered component (next)
		// into the persistent component instance (prev) so that it is aware of
		// what properties the parent has specified during SkipRender/Render
		// below.
		copyProps(next, prevComponent)
		next = prevComponent
	}

	// Before rendering, consult the Component's SkipRender method to see if we
	// should skip rendering or not.
	if rs, ok := next.(RenderSkipper); ok {
		prevRenderComponent := next.Context().prevRenderComponent
		if prevRenderComponent != nil {
			if next == prevRenderComponent {
				panic("vecty: internal error (SkipRender called with identical prev component)")
			}
			if rs.SkipRender(prevRenderComponent) {
				return nil, true, nil
			}
		}
	}

	// Render the component into HTML, handling nil renders.
	nextRender := next.Render()
	if nextRender == nil {
		// nil renders are translated into noscript tags.
		nextRender = Tag("noscript")
	}

	// Reconcile the actual rendered HTML.
	pendingMounts = nextRender.reconcile(extractHTML(prev))

	// Update the context to consider this render.
	next.Context().prevRender = nextRender
	next.Context().prevRenderComponent = doCopy(next)
	next.Context().unmounted = false
	return nextRender, false, pendingMounts
}

// mountUnmount determines whether a mount or unmount event should occur,
// actions unmounts recursively if appropriate, and returns either a Mounter,
// or nil.
func mountUnmount(next, prev ComponentOrHTML) Mounter {
	if next == prev {
		return nil
	}
	if prevHTML := extractHTML(prev); prevHTML != nil {
		if nextHTML := extractHTML(next); nextHTML == nil || prevHTML.node != nextHTML.node {
			for _, child := range prevHTML.children {
				unmount(child)
			}
		}
	}
	if u, ok := prev.(Unmounter); ok {
		u.Unmount()
	}
	if m, ok := next.(Mounter); ok {
		return m
	}
	return nil
}

// mount all pending Mounters
func mount(pendingMounts ...Mounter) {
	for _, mounter := range pendingMounts {
		if mounter == nil {
			continue
		}
		mounter.Mount()
	}
}

// unmount recursively unmounts the provided ComponentOrHTML, and any children
// that satisfy the Unmounter interface.
func unmount(e ComponentOrHTML) {
	if c, ok := e.(Component); ok {
		c.Context().unmounted = true
	}

	if list, ok := e.(List); ok {
		for _, child := range list {
			unmount(child)
		}
		return
	}

	if h := extractHTML(e); h != nil {
		for _, child := range h.children {
			unmount(child)
		}
	}

	if u, ok := e.(Unmounter); ok {
		u.Unmount()
	}
}

// RenderBody renders the given component as the document body. The given
// Component's Render method must return a "body" element.
func RenderBody(body Component) {
	nextRender, skip, pendingMounts := renderComponent(body, nil)
	if skip {
		panic("vecty: RenderBody Component.SkipRender returned true")
	}
	if nextRender.tag != "body" {
		panic(fmt.Sprintf("vecty: RenderBody expected Component.Render to return a body tag, found %q", nextRender.tag))
	}
	doc := global.Get("document")
	if doc.Get("readyState").String() == "loading" {
		doc.Call("addEventListener", "DOMContentLoaded", func() { // avoid duplicate body
			doc.Set("body", nextRender.node)
			mount(pendingMounts...)
			if m, ok := body.(Mounter); ok {
				mount(m)
			}
		})
		return
	}
	doc.Set("body", nextRender.node)
	mount(pendingMounts...)
	if m, ok := body.(Mounter); ok {
		mount(m)
	}
}

// SetTitle sets the title of the document.
func SetTitle(title string) {
	global.Get("document").Set("title", title)
}

// AddStylesheet adds an external stylesheet to the document.
func AddStylesheet(url string) {
	link := global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	global.Get("document").Get("head").Call("appendChild", link)
}

var (
	global    = wrapObject(js.Global)
	undefined = wrappedObject{js.Undefined}
)

type jsObject interface {
	Set(key string, value interface{})
	Get(key string) jsObject
	Delete(key string)
	Call(name string, args ...interface{}) jsObject
	String() string
	Bool() bool
}

func wrapObject(j *js.Object) jsObject {
	if j == nil {
		return nil
	}
	if j == js.Undefined {
		return undefined
	}
	return wrappedObject{j}
}

type wrappedObject struct {
	j *js.Object
}

func (w wrappedObject) Set(key string, value interface{}) {
	if v, ok := value.(wrappedObject); ok {
		value = v.j
	}
	w.j.Set(key, value)
}

func (w wrappedObject) Get(key string) jsObject {
	return wrapObject(w.j.Get(key))
}

func (w wrappedObject) Delete(key string) {
	w.j.Delete(key)
}

func (w wrappedObject) Call(name string, args ...interface{}) jsObject {
	for i, arg := range args {
		if v, ok := arg.(wrappedObject); ok {
			args[i] = v.j
		}
	}
	return wrapObject(w.j.Call(name, args...))
}

func (w wrappedObject) String() string {
	return w.j.String()
}
func (w wrappedObject) Bool() bool {
	return w.j.Bool()
}

var isTest bool
