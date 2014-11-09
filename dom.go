package dom

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

type Aspect interface {
	Apply(node js.Object, p, r float64)
	Revert()
}

type groupAspect []Aspect

func Group(aspects ...Aspect) Aspect {
	if len(aspects) == 1 {
		return aspects[0]
	}
	return groupAspect(aspects)
}

func (g groupAspect) Apply(node js.Object, p, r float64) {
	r2 := r / float64(len(g))
	for i, a := range g {
		a.Apply(node, p+r2*float64(i), r2)
	}
}

func (g groupAspect) Revert() {
	for _, a := range g {
		a.Revert()
	}
}

type nodeAspect struct {
	node  js.Object
	child Aspect
}

func Elem(tagName string, aspects ...Aspect) Aspect {
	return &nodeAspect{
		node:  js.Global.Get("document").Call("createElement", tagName),
		child: Group(aspects...),
	}
}

func (e *nodeAspect) Apply(node js.Object, p, r float64) {
	if !e.node.Get("previousSibling").IsNull() && e.node.Get("previousSibling").Get("gopherjsDomPosition").Float() > p {
		e.node.Call("remove")
	}
	if e.node.Get("parentNode").IsNull() {
		e.node.Set("gopherjsDomPosition", p)
		c := node.Get("firstChild")
		for !c.IsNull() && c.Get("gopherjsDomPosition").Float() < p {
			c = c.Get("nextSibling")
		}
		node.Call("insertBefore", e.node, c)
	}
	if e.child != nil {
		e.child.Apply(e.node, 0, 1)
	}
}

func (e *nodeAspect) Revert() {
	e.node.Call("remove")
}

func Text(content string) Aspect {
	return &nodeAspect{
		node: js.Global.Get("document").Call("createTextNode", content),
	}
}

type setPropAspect struct {
	name  string
	value string
}

func SetProp(name string, value string) Aspect {
	return &setPropAspect{name: name, value: value}
}

func (a *setPropAspect) Apply(node js.Object, p, r float64) {
	if node.Get(a.name).Str() != a.value {
		node.Set(a.name, a.value)
	}
}

func (a *setPropAspect) Revert() {
	// no reset
}

type togglePropAspect struct {
	name string
	node js.Object
}

func ToggleProp(name string) Aspect {
	return &togglePropAspect{name: name}
}

func (a *togglePropAspect) Apply(node js.Object, p, r float64) {
	a.node = node
	node.Set(a.name, true)
}

func (a *togglePropAspect) Revert() {
	a.node.Set(a.name, false)
}

type styleAspect struct {
	name  string
	value string
	style js.Object
}

func Style(name string, value string) Aspect {
	return &styleAspect{
		name:  name,
		value: value,
	}
}

func (a *styleAspect) Apply(node js.Object, p, r float64) {
	a.style = node.Get("style")
	a.style.Call("setProperty", a.name, a.value, "important")
}

func (a *styleAspect) Revert() {
	a.style.Call("removeProperty", a.name)
}

type Listener func(c *EventContext)

type EventContext struct {
	Node  js.Object
	Event js.Object
}

type eventAspect struct {
	eventType      string
	listener       func(event js.Object)
	preventDefault bool
	node           js.Object
}

func Event(eventType string, listener Listener) Aspect {
	var a *eventAspect
	a = &eventAspect{
		eventType: eventType,
		listener: func(event js.Object) {
			if a.preventDefault {
				event.Call("preventDefault")
			}
			go listener(&EventContext{
				Node:  a.node,
				Event: event,
			})
		},
		preventDefault: false,
	}
	return a
}

func PreventDefault(aspect Aspect) Aspect {
	aspect.(*eventAspect).preventDefault = true
	return aspect
}

func (a *eventAspect) Apply(node js.Object, p, r float64) {
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

type debugAspect struct {
	msg interface{}
}

func Debug(msg interface{}) Aspect {
	return &debugAspect{msg: msg}
}

func (a *debugAspect) Apply(node js.Object, p, r float64) {
	println("Apply:", fmt.Sprint(a.msg), node)
}

func (a *debugAspect) Revert() {
	println("Revert:", fmt.Sprint(a.msg))
}

func AddToBody(aspects ...Aspect) {
	Group(aspects...).Apply(js.Global.Get("document").Get("body"), 0, 1)
}
