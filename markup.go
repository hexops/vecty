package vecty

import "github.com/gopherjs/gopherjs/js"

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

// Class returns an Applyer which applies the provided classes. Subsequent
// calls to this function will append additional classes. To toggle classes,
// use ClassMap instead.
func Class(class ...string) Applyer {
	return markupFunc(func(h *HTML) {
		if h.classes == nil {
			h.classes = make(map[string]struct{})
		}
		for _, name := range class {
			h.classes[name] = struct{}{}
		}
	})
}

// ClassMap is markup that specifies classes to be applied to an element if
// their boolean value are true.
type ClassMap map[string]bool

// Apply implements the Applyer interface.
func (m ClassMap) Apply(h *HTML) {
	if h.classes == nil {
		h.classes = make(map[string]struct{})
	}
	for name, active := range m {
		if !active {
			delete(h.classes, name)
			continue
		}
		h.classes[name] = struct{}{}
	}
}

// markupList represents a list of Applyer which is individually
// applied to an HTML element or text node.
type markupList struct {
	list []Applyer
}

// Apply implements the Applyer interface.
func (m markupList) Apply(h *HTML) {
	for _, a := range m.list {
		if a == nil {
			continue
		}
		a.Apply(h)
	}
}

// If returns nil if cond is false, otherwise it returns the given children.
func If(cond bool, children ...ComponentOrHTML) ComponentOrHTML {
	if cond {
		return List(children)
	}
	return nil
}

// MarkupIf returns nil if cond is false, otherwise it returns the given markup.
//
// TODO(slimsag): we could offer (*HTML).WithMarkupIf(cond, markup...) instead?
func MarkupIf(cond bool, markup ...Applyer) Applyer {
	if cond {
		return markupList{list: markup}
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
