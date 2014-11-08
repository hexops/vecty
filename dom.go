package dom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Aspect interface {
	Apply(node js.Object)
}

type RevokableAspect interface {
	Aspect
	Revoke()
}

type groupAspect []Aspect

func Group(aspects ...Aspect) Aspect {
	if len(aspects) == 1 {
		return aspects[0]
	}
	return groupAspect(aspects)
}

func (g groupAspect) Apply(node js.Object) {
	for _, a := range g {
		a.Apply(node)
	}
}

func (g groupAspect) Revoke() {
	for _, a := range g {
		if ra, ok := a.(RevokableAspect); ok {
			ra.Revoke()
		}
	}
}

type elemAspect struct {
	tagName string
	aspect  Aspect
	node    js.Object
}

func Elem(tagName string, aspects ...Aspect) Aspect {
	return &elemAspect{
		tagName: tagName,
		aspect:  Group(aspects...),
	}
}

func (e *elemAspect) Apply(node js.Object) {
	if e.node == nil {
		e.node = js.Global.Get("document").Call("createElement", e.tagName)
	}
	e.aspect.Apply(e.node)
	node.Call("appendChild", e.node)
}

func (e *elemAspect) Revoke() {
	e.node.Call("remove")
}

type propAspect struct {
	name  string
	value interface{}
}

func Prop(name string, value interface{}) Aspect {
	return &propAspect{
		name:  name,
		value: value,
	}
}

func (p *propAspect) Apply(node js.Object) {
	node.Set(p.name, p.value)
}

type styleAspect struct {
	name  string
	value interface{}
}

func Style(name string, value interface{}) Aspect {
	return &styleAspect{
		name:  name,
		value: value,
	}
}

func (s *styleAspect) Apply(node js.Object) {
	node.Get("style").Set(s.name, s.value)
}

type Listener func(c *EventContext)

type EventContext struct {
	Node  js.Object
	Event js.Object
}

type eventAspect struct {
	eventType string
	listener  Listener
	node      js.Object
}

func Event(eventType string, listener Listener) Aspect {
	return &eventAspect{
		eventType: eventType,
		listener:  listener,
	}
}

func (l *eventAspect) Apply(node js.Object) {
	if l.node == nil {
		l.node = node
		node.Call("addEventListener", l.eventType, func(event js.Object) {
			go l.listener(&EventContext{
				Node:  l.node,
				Event: event,
			})
		})
	}
}

type textAspect struct {
	content string
	node    js.Object
}

func (a *textAspect) Apply(node js.Object) {
	if a.node == nil {
		a.node = js.Global.Get("document").Call("createTextNode", a.content)
	}
	node.Call("appendChild", a.node)
}

func (a *textAspect) Revoke() {
	a.node.Call("remove")
}

func Text(content string) Aspect {
	return &textAspect{
		content: content,
	}
}

func Body() js.Object {
	return js.Global.Get("document").Get("body")
}
