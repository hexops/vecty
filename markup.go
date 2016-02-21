package vecty

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

// Markup represents some markup that can be applied to a DOM element. For
// example, styles like font size, properties like checked status of an input,
// or child elements.
type Markup interface {
	// Apply should apply the markup to the given element.
	Apply(element *Element)
}

type property struct {
	Name  string
	Value interface{}
}

// Apply implements the Markup interface.
func (p *property) Apply(element *Element) {
	if element.Properties == nil {
		element.Properties = make(map[string]interface{})
	}
	if _, ok := element.Properties[p.Name]; ok {
		panic(fmt.Sprintf("duplicate property: %s", p.Name))
	}
	element.Properties[p.Name] = p.Value
}

// Property returns Markup which applies the given value to the named property
// of an DOM element.
func Property(name string, value interface{}) Markup {
	return &property{Name: name, Value: value}
}

type data struct {
	name  string
	value string
}

// Apply implements the Markup interface.
func (d *data) Apply(element *Element) {
	if element.Dataset == nil {
		element.Dataset = make(map[string]string)
	}
	if _, ok := element.Dataset[d.name]; ok {
		panic(fmt.Sprintf("duplicate data: %s", d.name))
	}
	element.Dataset[d.name] = d.value
}

// Data returns Markup which applies the given value to the named custom data
// attribute of a DOM element.
func Data(name, value string) Markup {
	return &data{name: name, value: value}
}

// ClassMap is markup that specifies classes to be applied to an element if
// their boolean value are true.
type ClassMap map[string]bool

// Apply implements the Markup interface.
func (m ClassMap) Apply(element *Element) {
	var classes []string
	for name, active := range m {
		if active {
			classes = append(classes, name)
		}
	}
	Property("className", strings.Join(classes, " ")).Apply(element)
}

type style struct {
	Name  string
	Value interface{}
}

// Apply implements the Markup interface.
func (s *style) Apply(element *Element) {
	if element.Style == nil {
		element.Style = make(map[string]interface{})
	}
	if _, ok := element.Style[s.Name]; ok {
		panic(fmt.Sprintf("duplicate style: %s", s.Name))
	}
	element.Style[s.Name] = s.Value
}

// Style returns Markup which applies the style with the given value to an
// element.
func Style(name string, value interface{}) Markup {
	return &style{Name: name, Value: value}
}

// EventListener is markup that specifies a callback function to be invoked when
// the named DOM event is fired.
type EventListener struct {
	Name               string
	Listener           func(*Event)
	callPreventDefault bool
	wrapper            func(jsEvent *js.Object)
}

// PreventDefault prevents the default behavior of the event from occuring.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault.
func (l *EventListener) PreventDefault() *EventListener {
	l.callPreventDefault = true
	return l
}

// Apply implements the Markup interface.
func (l *EventListener) Apply(element *Element) {
	element.EventListeners = append(element.EventListeners, l)
}

// Event represents a DOM event.
type Event struct {
	Target *js.Object
}

// List represents a list of markup which will all be applied to an element.
type List []Markup

// Apply implements the Markup interface.
func (g List) Apply(element *Element) {
	for _, m := range g {
		if m != nil {
			m.Apply(element)
		}
	}
}

// If returns nil if cond is false, otherwise it returns a list of markup.
func If(cond bool, markup ...Markup) Markup {
	if cond {
		return List(markup)
	}
	return nil
}
