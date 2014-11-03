package dom

import "github.com/gopherjs/gopherjs/js"

type ElemAspect struct {
	TagName string
	Aspects []Aspect
	Node    js.Object
}

type Aspect interface {
	Apply(parent *ElemAspect)
}

func Elem(tagName string, mutators ...Aspect) *ElemAspect {
	return &ElemAspect{
		TagName: tagName,
		Aspects: mutators,
	}
}

func (e *ElemAspect) Apply(parent *ElemAspect) {
	if e.Node == nil {
		e.Node = js.Global.Get("document").Call("createElement", e.TagName)
	}
	for _, m := range e.Aspects {
		m.Apply(e)
	}
	parent.Node.Call("appendChild", e.Node)
}

type PropAspect struct {
	Name  string
	Value interface{}
}

func Prop(name string, value interface{}) *PropAspect {
	return &PropAspect{
		Name:  name,
		Value: value,
	}
}

func (p *PropAspect) Apply(parent *ElemAspect) {
	parent.Node.Set(p.Name, p.Value)
}

type StyleAspect struct {
	Name  string
	Value interface{}
}

func Style(name string, value interface{}) *StyleAspect {
	return &StyleAspect{
		Name:  name,
		Value: value,
	}
}

func (s *StyleAspect) Apply(parent *ElemAspect) {
	parent.Node.Get("style").Set(s.Name, s.Value)
}

type EventAspect struct {
	EventType string
	Func      func(*ElemAspect)
}

func Event(eventType string, f func(*ElemAspect)) *EventAspect {
	return &EventAspect{
		EventType: eventType,
		Func:      f,
	}
}

func (l *EventAspect) Apply(parent *ElemAspect) {
	parent.Node.Call("addEventListener", l.EventType, func() {
		l.Func(parent)
	})
}

type TextAspect struct {
	Content string
	Node    js.Object
}

func (m *TextAspect) Apply(parent *ElemAspect) {
	if m.Node == nil {
		m.Node = js.Global.Get("document").Call("createTextNode", m.Content)
	}
	parent.Node.Call("appendChild", m.Node)
}

func Text(content string) *TextAspect {
	return &TextAspect{
		Content: content,
	}
}

func Body() *ElemAspect {
	return &ElemAspect{
		TagName: "body",
		Node:    js.Global.Get("document").Get("body"),
	}
}
