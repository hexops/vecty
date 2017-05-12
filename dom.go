package vecty

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

var (
	errMissingParent = errors.New("missing parent node")
	errEmptyElement  = errors.New("empty element or node")
)

// Core implements the private context method of the Component interface, and
// is the core/central struct which all Component implementations should embed.
type Core struct {
	prevRender   *HTML
	prevInstance Component
	mounted      bool
	unmounted    bool
}

// context implements the Component interface.
func (c *Core) context() *Core { return c }

// Clone is a general implementation for the Updater interface. This uses
// reflection, so it incurs a performance penalty, and will panic if the
// Component is not a pointer to struct. Components should provide their own
// implementation to avoid these limitations.
//
// See documentation for the Updater interface for an example
// of a component-specific implementation.
func (c *Core) Clone(current Component) Component {
	v := reflect.ValueOf(current)
	if v.Kind() != reflect.Ptr || v.IsNil() || v.Elem().Kind() != reflect.Struct {
		panic("Clone expects pointer to struct")
	}

	copy := reflect.New(v.Elem().Type())
	copy.Elem().Set(v.Elem())

	return copy.Interface().(Component)
}

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

	// context returns the component's context, which is used internally by
	// Vecty in order to store the previous component render for diffing.
	context() *Core
}

// Mounter is an optional interface that a Component can implement in order
// to receive component mount events.
type Mounter interface {
	Component
	// Mount is called after the component has been mounted, after the DOM
	// element has been added.
	Mount()
}

// Unmounter is an optional interface that a Component can implement in order
// to receive component unmount events.
type Unmounter interface {
	Component
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

// Updater is an optional interface that Components can implement in order to
// restore state during component reconciliation and also to short-circuit
// the reconciliation of a Component's output.
//
// Example implementation:
//  type MyComponent struct {
//  	vecty.Core
//  	state string
//  }
//
//  func (c *MyComponent) ShouldUpdate(prev Component) bool {
//  	p, ok := prev.(*MyComponent)
//  	if !ok {
//  		return true
//  	}
//  	if p.state != c.state {
//  		return true
//  	}
//  	return false
//  }
//
// You may optionally (and preferably) implement the Clone method (more on this
// below):
//
//  func (c *MyComponent) Clone(current vecty.Component) vecty.Component {
//  	m := &MyComponent{}
//  	*m = *c
//  	return m
//  }
type Updater interface {
	// ShouldUpdate is called before the component is requested to render, and
	// is passed a previous instance of a component. The previous component may
	// be nil.
	//
	// If ShouldUpdate determines that an update of the Component's rendered
	// output is required, `true` should be returned, and the component will be
	// rendered. If properties or state that would affect the component's output
	// have not changed, returning `false` will skip rendering of the component
	// for this pass.
	ShouldUpdate(prev Component) bool

	// Clone returns a copy of the current state of the component, used during
	// reconcilliation by the Updater interface.
	//
	// If a Component does not implement this method, a general implementation
	// from vecty.Core will be used, however this uses reflection, so it will
	// be slower, and will panic if the Component is not a pointer to struct.
	//
	// The `current` argument is guaranteed to be equal to the Component, and so
	// may be ignored by implementers - it is only there to serve the generic
	// implementation in vecty.Core.
	Clone(current Component) Component
}

// Keyer is an optional interface that Components may implement to uniquely
// identify an instance amongst a list of its siblings. Implementing Keyer
// will provide a significant performance boost when rendering siblings that
// may change order.
//
// Inside lists of Markup (e.g. vecty.List), vecty.Key(string) may be used
// instead.
type Keyer interface {
	// Key must return a string that is unique amongst sibiling elements when
	// implemented, otherwise the renderer will panic.
	Key() string
}

// HTML represents some form of HTML: an element with a specific tag, or some
// literal text (a TextNode).
type HTML struct {
	node jsObject

	namespace, tag, text, innerHTML, key string
	styles, dataset                      map[string]string
	properties, attributes               map[string]interface{}
	eventListeners                       []*EventListener
	children                             []ComponentOrHTML
	childIndex                           map[string]int
	lifecycleEvents                      []lifecycleEvent
	new                                  bool
}

// Key satisfies the Keyer interface
func (h *HTML) Key() string {
	return h.key
}

func (h *HTML) Node() *js.Object { return h.node.(wrappedObject).j }

// newNode returns a new element node or panics on invalid HTML.
func (h *HTML) newNode() jsObject {
	h.new = true
	switch {
	case h.tag != "" && h.text != "":
		panic("vecty: only one of HTML.tag or HTML.text may be set")
	case h.text != "" && h.innerHTML != "":
		panic("vecty: only HTML may have UnsafeHTML attribute")
	case h.tag != "" && h.namespace == "":
		return global.Get("document").Call("createElement", h.tag)
	case h.tag != "" && h.namespace != "":
		return global.Get("document").Call("createElementNS", h.namespace, h.tag)
	default:
		return global.Get("document").Call("createTextNode", h.text)
	}
}

// parentNode returns the parent Node for a ComponentOrHTML
func (h *HTML) parentNode(el ComponentOrHTML) (jsObject, error) {
	e := assertHTML(el)
	if e == nil || e.node == nil {
		return nil, errEmptyElement
	}

	parent := e.node.Get(`parentNode`)
	// TODO: parent == nil || parent == js.Undefined
	if parent == nil {
		return nil, errMissingParent
	}

	return parent, nil
}

// appendChild to this HTML
func (h *HTML) appendChild(next ComponentOrHTML) {
	n := assertHTML(next)
	if n == nil {
		return
	}
	h.node.Call("appendChild", n.node)
}

// replace prev with self in prev's parent
func (h *HTML) replace(prev ComponentOrHTML) error {
	if prev == nil {
		return errEmptyElement
	}
	p := assertHTML(prev)
	if p == nil || p.node == nil {
		return errEmptyElement
	}

	if h.node == p.node {
		return nil
	}

	parent, err := h.parentNode(p)
	if err != nil {
		return err
	}

	parent.Call("replaceChild", h.node, p.node)
	return nil
}

// replaceChild on prev parent Node, or append to this HTML on failure
func (h *HTML) replaceChild(next, prev ComponentOrHTML) error {
	if next == prev {
		return nil
	} else if h.new {
		h.appendChild(next)
		return nil
	}

	n, p := assertHTML(next), assertHTML(prev)
	if n == nil || n.node == nil || p == nil || p.node == nil {
		return errEmptyElement
	}
	if n.node == p.node {
		return nil
	}
	parent, err := h.parentNode(p)
	if err != nil {
		return err
	}

	parent.Call("replaceChild", n.node, p.node)
	return nil
}

// insert next element before prev element in prev parent
func (h *HTML) insertChildBefore(next, prev ComponentOrHTML) error {
	parent, err := h.parentNode(prev)
	if err != nil {
		return err
	}

	n, p := assertHTML(next), assertHTML(prev)
	if n == nil || n.node == nil {
		return errEmptyElement
	}
	if p == nil {
		p = &HTML{}
	}

	parent.Call("insertBefore", n.node, p.node)
	return nil
}

// insert next element after prev element in prev parent
func (h *HTML) insertChildAfter(next, prev ComponentOrHTML) error {
	p := assertHTML(prev)
	if p == nil {
		return errEmptyElement
	}

	return h.insertChildBefore(next, p.node.Get("nextSibling"))
}

// remove previous child
func (h *HTML) removeChild(prev ComponentOrHTML) error {
	if prev == nil {
		return nil
	}
	p, ok := prev.(*HTML)
	if !ok {
		p = prev.(Component).context().prevRender
		prev.(Component).context().prevRender = nil
	}
	if p.node == nil {
		return nil
	}
	parent, err := h.parentNode(p)
	if err != nil {
		return err
	}
	parent.Call("removeChild", p.node)

	return nil
}

// mutate the prev HTML Node to the desired state for the current element.
func (h *HTML) mutate(prev *HTML) *HTML {
	// On compatible tag, mutate previous node, otherwise start from clean element
	if (h.text != "" && prev.text != "") || (h.tag != "" && prev.tag != "" && h.tag == prev.tag && h.namespace == prev.namespace) {
		h.node = prev.node
	} else {
		h.node = h.newNode()
	}

	// Mutate text element
	if h.text != "" {
		if h.text != prev.text {
			h.node.Set("nodeValue", h.text)
		}
		return h
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
				// TODO: Set(name, js.Undefined)
				h.node.Set(name, nil)
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

	// Event listeners
	if !h.new {
		for _, l := range prev.eventListeners {
			h.node.Call("removeEventListener", l.Name, l.wrapper)
		}
	}
	for _, l := range h.eventListeners {
		h.node.Call("addEventListener", l.Name, l.wrapper)
	}

	// Unsafe innerHTML
	if h.innerHTML != prev.innerHTML {
		h.node.Set("innerHTML", h.innerHTML)
	}

	// If there are no children, we're finished.
	if len(h.children) == 0 && len(prev.children) == 0 {
		return h
	}

	// Iterate over child elements and render recursively
	h.childIndex = make(map[string]int)
	h.lifecycleEvents = make([]lifecycleEvent, 0)
	nextKeyed := false
	prevKeyed := len(prev.children) == len(prev.childIndex) && len(prev.children) > 0
	focus := global.Get("document").Get("activeElement")
	for i, nextChild := range h.children {
		var (
			key        string
			prevChild  ComponentOrHTML
			prevIndex  int
			foundKey   bool
			nextRender *HTML
		)

		// key detection
		if keyer, ok := nextChild.(Keyer); ok {
			key = keyer.Key()
			if key != "" {
				nextKeyed = true
				h.childIndex[key] = i
			} else if nextKeyed {
				panic("All siblings must have keys when using keyed elements")
			}
		}

		// Find prevChild by key or index, otherwise it is nil
		switch {
		case prevKeyed && nextKeyed:
			if prevIndex, foundKey = prev.childIndex[key]; foundKey {
				prevChild = prev.children[prevIndex]
				delete(prev.childIndex, key)
			}
		case i < len(prev.children):
			prevChild = prev.children[i]
		}

		// Render new child
		nextRender = renderComponentOrHTML(nextChild, prevChild)

		// Select best insertion method
		switch {
		case i >= len(prev.children) || h.new:
			// Append the child to this container
			h.appendChild(nextRender)
		case prevKeyed && nextKeyed:
			// If we couldn't re-use the element, didn't find an existing key,
			// or the element position wasn't stable, insert at the correct
			// position. Insert will move the element for us.
			if h.new || !foundKey || i != prevIndex {
				if err := h.insertChildAfter(nextRender, prev.children[i]); err != nil {
					h.appendChild(nextRender)
				}
			}
		default:
			// Otherwise, replace elements where necessary
			if err := h.replaceChild(nextChild, prevChild); err != nil {
				panic(err)
			}
		}

		h.lifecycleEvents = append(h.lifecycleEvents, &eventMountUnmount{next: nextChild, prev: prevChild})
	}

	// TODO: focus != nil && focus != js.Undefined
	if focus != nil {
		focus.Call("focus")
	}

	// Remove dangling children
	if prevKeyed && nextKeyed {
		for _, i := range prev.childIndex {
			if err := h.removeChild(prev.children[i]); err != nil {
				panic(err)
			}
			h.lifecycleEvents = append(h.lifecycleEvents, &eventUnmount{prev: prev.children[i]})
		}
		return h
	}

	for i := len(h.children); i < len(prev.children); i++ {
		if err := h.removeChild(prev.children[i]); err != nil {
			panic(err)
		}
		h.lifecycleEvents = append(h.lifecycleEvents, &eventUnmount{prev: prev.children[i]})
	}

	return h
}

// render generates the DOM representation of the HTML.
func (h *HTML) render(prev *HTML) *HTML {
	// Wrap eventListeners
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

	// Populate prev if empty
	if prev == nil {
		prev = &HTML{}
	}
	if prev.node == nil {
		prev.node = h.newNode()
	}

	// Mutate element
	return h.mutate(prev)
}

// shouldUpdate component?
func shouldUpdate(c Component) bool {
	// Always render new components
	if c.context().prevRender == nil {
		return true
	}
	// Check with component, when supported
	if updater, ok := c.(Updater); ok {
		return updater.ShouldUpdate(c.context().prevInstance)
	}
	// Update by default
	return true
}

// renderComponent returns the HTML output from the Component.
func renderComponent(c Component, prev ComponentOrHTML) Component {
	if !shouldUpdate(c) {
		return c
	}

	// Render component HTML
	render := renderComponentOrHTML(c.Render(), prev)
	// Store current instance after render, in case render updates state
	if updater, ok := c.(Updater); ok {
		c.context().prevInstance = updater.Clone(c)
	}
	// Store previous render
	c.context().prevRender = render

	return c
}

// renderComponentOrHTML is the main entry point for the render chain.
func renderComponentOrHTML(next, prev ComponentOrHTML) *HTML {
	// Empty output rendered as `noscript` tag
	if next == nil {
		next = Tag("noscript")
	}

	n, ok := next.(*HTML)
	if !ok {
		// Render Component
		c := next.(Component)
		return renderComponent(c, prev).context().prevRender
	}

	// Render HTML with mutation
	return n.render(assertHTML(prev))
}

// Rerender causes the body of the given component (i.e. the HTML returned by
// the Component's Render method) to be re-rendered.
func Rerender(c Component) {
	prevRender := c.context().prevRender
	nextRender := renderComponentOrHTML(c, prevRender)
	if prevRender != nil && nextRender.new {
		if err := nextRender.replace(prevRender); err != nil {
			panic(err)
		}
	}
	e := &eventMountUnmount{next: c, prev: c}
	e.trigger()
}

// RenderBody renders the given component as the document body. The given
// Component's Render method must return a "body" element.
func RenderBody(body Component) {
	render := renderComponentOrHTML(body, nil)
	if render.tag != "body" {
		panic(fmt.Sprintf("vecty: RenderBody expected Component.Render to return a body tag, found %q", render.tag))
	}
	doc := global.Get("document")
	if doc.Get("readyState").String() != "loading" {
		doc.Set("body", render.node)
		e := &eventMount{next: body}
		e.trigger()
		return
	}
	doc.Call("addEventListener", "DOMContentLoaded", func() { // avoid duplicate body
		doc.Set("body", render.node)
		e := &eventMount{next: body}
		e.trigger()
	})
}

type lifecycleEvent interface {
	trigger()
}

// eventMountUnmount recursively triggers the Mount and Unmount handlers of
// next and prev and their children respectively, if appropriate
type eventMountUnmount struct {
	prev ComponentOrHTML
	next ComponentOrHTML
}

func (e *eventMountUnmount) trigger() {
	h := assertHTML(e.next)
	if h != nil {
		for _, evt := range h.lifecycleEvents {
			evt.trigger()
		}
		h.lifecycleEvents = nil
	}

	var shouldMount, shouldUnmount bool
	nextComponent, nextIsComponent := e.next.(Component)
	prevComponent, prevIsComponent := e.prev.(Component)
	switch {
	case e.prev == nil && e.next != nil && nextIsComponent:
		// Had nil, now have Component, mount next
		shouldMount = true
	case nextIsComponent != prevIsComponent:
		if prevIsComponent {
			// Had Component, now have HTML, unmount prev
			shouldUnmount = true
		} else {
			// Had HTML, now have Component, mount next
			shouldMount = true
		}
	case nextIsComponent && nextComponent != prevComponent:
		// Have inequal Components, unmount prev, mount next
		shouldUnmount = true
		shouldMount = true
	}
	if shouldUnmount {
		unmount(e.prev)
	}
	if shouldMount {
		mount(e.next)
	}
}

// eventMount triggers the Mount event of next if appropriate
type eventMount struct {
	next ComponentOrHTML
}

func (e *eventMount) trigger() {
	h := assertHTML(e.next)
	if h != nil {
		for _, evt := range h.lifecycleEvents {
			evt.trigger()
		}
		h.lifecycleEvents = nil
	}
	mount(e.next)
}

// eventUnmount triggers the Unmount event of prev if appropriate
type eventUnmount struct {
	prev ComponentOrHTML
}

func (e *eventUnmount) trigger() {
	h := assertHTML(e.prev)
	if h != nil {
		for _, evt := range h.lifecycleEvents {
			evt.trigger()
		}
		h.lifecycleEvents = nil
	}
	unmount(e.prev)
}

// mount calls the Mount function on Mounter components
func mount(e ComponentOrHTML) {
	if m, ok := e.(Mounter); ok {
		if m.context().mounted {
			unmount(e)
		}
		m.context().mounted = true
		m.context().unmounted = false
		m.Mount()
	}
}

// unmount calls the Unmount function on Unmounter components
func unmount(e ComponentOrHTML) {
	if u, ok := e.(Unmounter); ok {
		if u.context().unmounted {
			return
		}
		u.context().unmounted = true
		u.context().mounted = false
		u.Unmount()
	}
}

// assertHTML is a convenience method for casting to *HTML
func assertHTML(e ComponentOrHTML) *HTML {
	if e == nil {
		return nil
	}
	h, ok := e.(*HTML)
	if !ok {
		if c, ok := e.(Component); ok {
			h = c.context().prevRender
		}
	}

	return h
}

var global jsObject = wrapObject(js.Global)

type jsObject interface {
	Set(key string, value interface{})
	Get(key string) jsObject
	Call(name string, args ...interface{}) jsObject
	String() string
	Bool() bool
}

func wrapObject(j *js.Object) jsObject {
	if j == nil {
		return nil
	}
	if j == js.Undefined {
		panic("TODO")
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
