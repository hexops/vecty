package vecty

import (
	"fmt"
	"strings"

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

// PreventDefault prevents the default behavior of the event from occuring.
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

// Apply implements the Markup interface.
func (l *EventListener) Apply(h *HTML) {
	h.eventListeners = append(h.eventListeners, l)
}

// Event represents a DOM event.
type Event struct {
	Target *js.Object
}

// MarkupOrComponentOrHTML represents one of:
//
//  Markup
//  Component
//  *HTML
//
// If the underlying value is not one of these types, the code handling the
// value is expected to panic.
type MarkupOrComponentOrHTML interface{}

func apply(m MarkupOrComponentOrHTML, h *HTML) {
	if m == nil {
		return
	}
	switch m := m.(type) {
	case Markup:
		m.Apply(h)
	case Component:
		h.children = append(h.children, m)
	case *HTML:
		h.children = append(h.children, m)
	default:
		panic(fmt.Sprintf("vecty: invalid type %T does not match MarkupOrComponent interface", m))
	}
}

// Markup represents some type of markup (a style, property, data, etc) which
// can be applied to a given HTML element or text node.
type Markup interface {
	// Apply applies the markup to the given HTML element or text node.
	Apply(h *HTML)
}

type markupFunc func(h *HTML)

func (m markupFunc) Apply(h *HTML) { m(h) }

// Style returns Markup which applies the given CSS style. Generally, this
// function is not used directly but rather the style subpackage (which is type
// safe) is used instead.
func Style(key, value string) Markup {
	return markupFunc(func(h *HTML) {
		if h.styles == nil {
			h.styles = make(map[string]string)
		}
		h.styles[key] = value
	})
}

// Property returns Markup which applies the given JavaScript property to an
// HTML element or text node. Generally, this function is not used directly but
// rather the style subpackage (which is type safe) is used instead.
func Property(key string, value interface{}) Markup {
	return markupFunc(func(h *HTML) {
		if h.properties == nil {
			h.properties = make(map[string]interface{})
		}
		h.properties[key] = value
	})
}

// Data returns Markup which applies the given data attribute.
func Data(key, value string) Markup {
	return markupFunc(func(h *HTML) {
		h.dataset[key] = value
	})
}

// ClassMap is markup that specifies classes to be applied to an element if
// their boolean value are true.
type ClassMap map[string]bool

// Apply implements the Markup interface.
func (m ClassMap) Apply(h *HTML) {
	var classes []string
	for name, active := range m {
		if active {
			classes = append(classes, name)
		}
	}
	Property("className", strings.Join(classes, " ")).Apply(h)
}

// List represents a list of Markup, Component, or HTML which is individually
// applied to an HTML element or text node.
type List []MarkupOrComponentOrHTML

// Apply implements the Markup interface.
func (l List) Apply(h *HTML) {
	for _, m := range l {
		apply(m, h)
	}
}

// If returns nil if cond is false, otherwise it returns the given markup.
func If(cond bool, markup ...MarkupOrComponentOrHTML) MarkupOrComponentOrHTML {
	if cond {
		return List(markup)
	}
	return nil
}
