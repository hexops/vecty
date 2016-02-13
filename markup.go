package vecty

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

type Markup interface {
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

type EventListener struct {
	Name               string
	Listener           func(*Event)
	callPreventDefault bool
	wrapper            func(jsEvent *js.Object)
}

func (l *EventListener) PreventDefault() *EventListener {
	l.callPreventDefault = true
	return l
}

// Apply implements the Markup interface.
func (l *EventListener) Apply(element *Element) {
	element.EventListeners = append(element.EventListeners, l)
}

type Event struct {
	Target *js.Object
}

type List []Markup

// Apply implements the Markup interface.
func (g List) Apply(element *Element) {
	for _, m := range g {
		if m != nil {
			m.Apply(element)
		}
	}
}

func If(cond bool, markup ...Markup) Markup {
	if cond {
		return List(markup)
	}
	return nil
}
