package dom

import (
	"github.com/gopherjs/gopherjs/js"
)

type Aspect interface {
	Apply(node js.Object)
	Revert()
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

func (g groupAspect) Revert() {
	for _, a := range g {
		a.Revert()
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

func (e *elemAspect) Revert() {
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

func (a *propAspect) Apply(node js.Object) {
	node.Set(a.name, a.value)
}

func (a *propAspect) Revert() {
	// TODO
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

func (a *styleAspect) Apply(node js.Object) {
	node.Get("style").Set(a.name, a.value)
}

func (a *styleAspect) Revert() {
	// TODO
}

type Listener func(c *EventContext)

type EventContext struct {
	Node  js.Object
	Event js.Object
}

type eventAspect struct {
	eventType string
	listener  func(event js.Object)
	node      js.Object
}

func Event(eventType string, listener Listener) Aspect {
	var a *eventAspect
	a = &eventAspect{
		eventType: eventType,
		listener: func(event js.Object) {
			go listener(&EventContext{
				Node:  a.node,
				Event: event,
			})
		},
	}
	return a
}

func (a *eventAspect) Apply(node js.Object) {
	if a.node == nil {
		a.node = node
		a.node.Call("addEventListener", a.eventType, a.listener)
	}
}

func (a *eventAspect) Revert() {
	if a.node != nil {
		a.node.Call("removeEventListener", a.eventType, a.listener)
		a.node = nil
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

func (a *textAspect) Revert() {
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
