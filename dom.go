package vecty

import (
	"fmt"
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

// Core implements the Context method of the Component interface, and is the
// core/central struct which all Component implementations should embed.
type Core struct {
	prevComponent, prevRenderComponent Component
	prevRender                         *HTML
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

// Unmounter is an optional interface that a Component can implement in order
// to receive component unmount events.
type Unmounter interface {
	// Unmount is called after the component has been unmounted, after the DOM
	// element has been removed.
	Unmount()
}

// ComponentOrHTML represents one of:
//
//  Component
//  *HTML
//
// If the underlying value is not one of these types, the code handling the
// value is expected to panic.
type ComponentOrHTML interface{}

// Restorer is an optional interface that Component's can implement in order to
// restore state during component reconciliation.
type Restorer interface {
	// Restore is called when the component should restore itself against a
	// previous instance of a component.
	//
	// The previous component may be nil or of a different type than this
	// Restorer itself, thus a type assertion should be used no action taken
	// if the type does not match.
	Restore(prev Component)
}

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
	new                             bool
}

// Node returns the underlying JavaScript Element or TextNode.
func (h *HTML) Node() *js.Object { return h.node.(wrappedObject).j }

func (h *HTML) createNode() {
	h.new = true
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

func (h *HTML) reconcile(prev *HTML) {
	// Check for compatible tag and mutate previous instance on match, otherwise start fresh
	switch {
	case prev != nil && h.tag == "" && prev.tag == "":
		// Compatible text node
		h.reconcileText(prev)
		return
	case prev != nil && h.tag != "" && prev.tag != "" && h.tag == prev.tag && h.namespace == prev.namespace:
		// Compatible element node
		h.node = prev.node
	default:
		// Incompatible node, start fresh
		if prev == nil {
			prev = &HTML{}
		}
		h.createNode()
		defer func() {
			h.new = false
		}()
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
	if !h.new {
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
	if !h.new {
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
	if !h.new {
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
	if !h.new {
		for name := range prev.styles {
			if _, ok := h.styles[name]; !ok {
				style.Call("removeProperty", name)
			}
		}
	}

	if !h.new {
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

	// TODO better list element reuse
	for i, nextChild := range h.children {
		// TODO(pdf): Add tests for h.new usage
		if i >= len(prev.children) || h.new {
			nextChildRender, skip := render(nextChild, nil)
			if skip {
				continue
			}
			h.node.Call("appendChild", nextChildRender.node)
			continue
		}

		prevChild := prev.children[i]
		prevChildRender, ok := prevChild.(*HTML)
		if !ok {
			prevChildRender = prevChild.(Component).Context().prevRender
		}

		nextChildRender, skip := render(nextChild, prevChildRender)
		if nextChildRender == prevChildRender {
			panic("vecty: next child render must not equal previous child render (did the child Render illegally return a stored render variable?)")
		}
		if skip {
			continue
		}
		replaceNode(nextChildRender.node, prevChildRender.node)
	}
	for i := len(h.children); i < len(prev.children); i++ {
		prevChild := prev.children[i]
		prevChildRender, ok := prevChild.(*HTML)
		if !ok {
			prevChildRender = prevChild.(Component).Context().prevRender
		}
		removeNode(prevChildRender.node)
		if u, ok := prevChild.(Unmounter); ok {
			u.Unmount()
		}
	}
}

// Tag returns an HTML element with the given tag name. Generally, this
// function is not used directly but rather the elem subpackage (which is type
// safe) is used instead.
func Tag(tag string, m ...MarkupOrComponentOrHTML) *HTML {
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
func Text(text string, m ...MarkupOrComponentOrHTML) *HTML {
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
// If the component has not been rendered before, Rerender panics.
func Rerender(c Component) {
	if c == nil {
		panic("vecty: Rerender illegally called with a nil Component argument")
	}
	prevRender := c.Context().prevRender
	if prevRender == nil {
		panic("vecty: Rerender invoked on Component that has never been rendered")
	}
	nextRender, skip := renderComponent(c, prevRender)
	if skip {
		return
	}
	replaceNode(nextRender.node, prevRender.node)
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

// render the ComponentOrHTML. In the case of *HTML, its Restore method is
// invoked with the specified prevRender and c.(*HTML) is returned.
//
// In the case of a Component, renderComponent(c.(Component)) is returned.
func render(c ComponentOrHTML, prevRender *HTML) (h *HTML, skip bool) {
	switch v := c.(type) {
	case *HTML:
		v.reconcile(prevRender)
		return v, false
	case Component:
		return renderComponent(v, prevRender)
	default:
		panic(fmt.Sprintf("vecty: encountered invalid ComponentOrHTML %T", c))
	}
}

// renderComponent handles rendering the given Component into *HTML. If skip ==
// true is returned, the Component's SkipRender method has signaled the
// component does not need to be rendered and h == nil is returned.
func renderComponent(comp Component, prevRender *HTML) (h *HTML, skip bool) {
	// Now that we know we are rendering the component, restore friendly
	// relations between the component and the previous component.
	if comp != comp.Context().prevComponent {
		if r, ok := comp.(Restorer); ok {
			r.Restore(comp.Context().prevComponent)
		}
	}

	// Before rendering, consult the Component's SkipRender method to see if we
	// should skip rendering or not.
	if rs, ok := comp.(RenderSkipper); ok {
		prev := comp.Context().prevRenderComponent
		if prev != nil {
			if comp == prev {
				panic("vecty: internal error (SkipRender called with identical prev component)")
			}
			if rs.SkipRender(prev) {
				return nil, true
			}
		}
	}

	// Render the component into HTML, handling nil renders.
	nextRender := comp.Render()
	if nextRender == nil {
		// nil renders are translated into noscript tags.
		nextRender = Tag("noscript")
	}

	// Restore the actual rendered HTML.
	nextRender.reconcile(prevRender)

	// Update the context to consider this render.
	comp.Context().prevRender = nextRender
	comp.Context().prevComponent = comp
	comp.Context().prevRenderComponent = doCopy(comp)
	return nextRender, false
}

// RenderBody renders the given component as the document body. The given
// Component's Render method must return a "body" element.
func RenderBody(body Component) {
	nextRender, skip := renderComponent(body, nil)
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
		})
		return
	}
	doc.Set("body", nextRender.node)
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
