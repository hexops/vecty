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

type RevokableAspect interface {
	Aspect
	Revoke()
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

func (e *ElemAspect) Revoke() {
	e.Node.Call("remove")
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
	Fun       func()
	Element   *ElemAspect
}

func Event(eventType string, fun func()) *EventAspect {
	return &EventAspect{
		EventType: eventType,
		Fun:       fun,
	}
}

func (l *EventAspect) Apply(parent *ElemAspect) {
	l.Element = parent
	parent.Node.Call("addEventListener", l.EventType, func() {
		l.Fun()
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
