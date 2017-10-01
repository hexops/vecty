package vecty

import (
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

// EventListener is markup that specifies a callback function to be invoked when
// the named DOM event is fired.
type EventListener struct {
	Name                string
	Listener            func(*Event)
	callPreventDefault  bool
	callStopPropagation bool
	wrapper             func(jsEvent *js.Object)
}

// PreventDefault prevents the default behavior of the event from occurring.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault.
func (l *EventListener) PreventDefault() *EventListener {
	l.callPreventDefault = true
	return l
}

// StopPropagation prevents further propagation of the current event in the
// capturing and bubbling phases.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation.
func (l *EventListener) StopPropagation() *EventListener {
	l.callStopPropagation = true
	return l
}

// Apply implements the Applyer interface.
func (l *EventListener) Apply(h *HTML) {
	h.eventListeners = append(h.eventListeners, l)
}

// Event represents a DOM event.
type Event struct {
	*js.Object
	Target *js.Object
}

// MarkupOrChild represents one of:
//
//  Component
//  *HTML
//  List
//  KeyedList
//  nil
//  MarkupList
//
// If the underlying value is not one of these types, the code handling the
// value is expected to panic.
type MarkupOrChild interface{}

func apply(m MarkupOrChild, h *HTML) {
	switch m := m.(type) {
	case MarkupList:
		m.Apply(h)
	case Component, *HTML, List, KeyedList, nil:
		h.children = append(h.children, m)
	default:
		panic("vecty: invalid type " + reflect.TypeOf(m).String() + " does not match MarkupOrChild interface")
	}
}

// Applyer represents some type of markup (a style, property, data, etc) which
// can be applied to a given HTML element or text node.
type Applyer interface {
	// Apply applies the markup to the given HTML element or text node.
	Apply(h *HTML)
}

type markupFunc func(h *HTML)

func (m markupFunc) Apply(h *HTML) { m(h) }

// Style returns an Applyer which applies the given CSS style. Generally, this
// function is not used directly but rather the style subpackage (which is type
// safe) should be used instead.
func Style(key, value string) Applyer {
	return markupFunc(func(h *HTML) {
		if h.styles == nil {
			h.styles = make(map[string]string)
		}
		h.styles[key] = value
	})
}

// Key returns an Applyer that uniquely identifies the HTML element amongst its
// siblings. When used, all other sibling elements and components must also be
// keyed.
func Key(key interface{}) Applyer {
	return markupFunc(func(h *HTML) {
		h.key = key
	})
}

// Property returns an Applyer which applies the given JavaScript property to an
// HTML element or text node. Generally, this function is not used directly but
// rather the prop and style subpackages (which are type safe) should be used instead.
//
// To set style, use style package or Style. Property panics if key is "style".
func Property(key string, value interface{}) Applyer {
	if key == "style" {
		panic(`vecty: Property called with key "style"; style package or Style should be used instead`)
	}
	return markupFunc(func(h *HTML) {
		if h.properties == nil {
			h.properties = make(map[string]interface{})
		}
		h.properties[key] = value
	})
}

// Attribute returns an Applyer which applies the given attribute to an element.
//
// In most situations, you should use Property function, or the prop subpackage
// (which is type-safe) instead. There are only a few attributes (aria-*, role,
// etc) which do not have equivalent properties. Always opt for the property
// first, before relying on an attribute.
func Attribute(key string, value interface{}) Applyer {
	return markupFunc(func(h *HTML) {
		if h.attributes == nil {
			h.attributes = make(map[string]interface{})
		}
		h.attributes[key] = value
	})
}

// Data returns an Applyer which applies the given data attribute.
func Data(key, value string) Applyer {
	return markupFunc(func(h *HTML) {
		if h.dataset == nil {
			h.dataset = make(map[string]string)
		}
		h.dataset[key] = value
	})
}

// ClassMap is markup that specifies classes to be applied to an element if
// their boolean value are true.
type ClassMap map[string]bool

// Apply implements the Applyer interface.
func (m ClassMap) Apply(h *HTML) {
	var classes []string
	for name, active := range m {
		if active {
			classes = append(classes, name)
		}
	}
	Property("className", join(classes, " ")).Apply(h)
}

// MarkupList represents a list of Applyer which is individually
// applied to an HTML element or text node.
//
// It may only be created through the Markup function.
type MarkupList struct {
	list []Applyer
}

// Apply implements the Applyer interface.
func (m MarkupList) Apply(h *HTML) {
	for _, a := range m.list {
		if a == nil {
			continue
		}
		a.Apply(h)
	}
}

// Markup wraps a list of Applyer which is individually
// applied to an HTML element or text node.
func Markup(m ...Applyer) MarkupList {
	// returns public non-pointer struct value with private field so that users
	// must acquire a MarkupList only from this function, and so that it can
	// never be nil (which would make it indistinguishable from (*HTML)(nil) in
	// a call to e.g. Tag).
	return MarkupList{list: m}
}

// If returns nil if cond is false, otherwise it returns the given markup.
func If(cond bool, markup ...ComponentOrHTML) MarkupOrChild {
	if cond {
		return List(markup)
	}
	return nil
}

// MarkupIf returns nil if cond is false, otherwise it returns the given markup.
func MarkupIf(cond bool, markup ...Applyer) Applyer {
	if cond {
		return Markup(markup...)
	}
	return nil
}

// UnsafeHTML is Applyer which unsafely sets the inner HTML of an HTML element.
//
// It is entirely up to the caller to ensure the input HTML is properly
// sanitized.
//
// It is akin to innerHTML in standard JavaScript and dangerouslySetInnerHTML
// in React, and is said to be unsafe because Vecty makes no effort to validate
// or ensure the HTML is safe for insertion in the DOM. If the HTML came from a
// user, for example, it would create a cross-site-scripting (XSS) exploit in
// the application.
//
// The returned Applyer can only be applied to HTML, not vecty.Text, or else a
// panic will occur.
func UnsafeHTML(html string) Applyer {
	return markupFunc(func(h *HTML) {
		h.innerHTML = html
	})
}

// Namespace is Applyer which sets the namespace URI to associate with the
// created element. This is primarily used when working with, e.g., SVG.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Document/createElementNS#Valid Namespace URIs
func Namespace(uri string) Applyer {
	return markupFunc(func(h *HTML) {
		h.namespace = uri
	})
}

// join is extracted from the stdlib `strings` package
func join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}
