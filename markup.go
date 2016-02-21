package vecty

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

type Markup interface {
	Apply(element *Element)
}

type Property struct {
	Name  string
	Value interface{}
}

// Apply implements the Markup interface.
func (p *Property) Apply(element *Element) {
	if element.Properties == nil {
		element.Properties = make(map[string]interface{})
	}
	if _, ok := element.Properties[p.Name]; ok {
		panic(fmt.Sprintf("duplicate property: %s", p.Name))
	}
	element.Properties[p.Name] = p.Value
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
	p := Property{
		Name:  "className",
		Value: strings.Join(classes, " "),
	}
	p.Apply(element)
}

type Style struct {
	Name  string
	Value interface{}
}

// Apply implements the Markup interface.
func (s *Style) Apply(element *Element) {
	if element.Style == nil {
		element.Style = make(map[string]interface{})
	}
	if _, ok := element.Style[s.Name]; ok {
		panic(fmt.Sprintf("duplicate style: %s", s.Name))
	}
	element.Style[s.Name] = s.Value
}

type EventListener struct {
	Name               string
	Listener           func(*Event)
	CallPreventDefault bool
	wrapper            func(jsEvent *js.Object)
}

func (l *EventListener) PreventDefault() *EventListener {
	l.CallPreventDefault = true
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
