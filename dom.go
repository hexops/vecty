package dom

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom/domutil"
)

type Markup interface {
	Apply(element *Element)
}

type Instance interface {
	Node() *js.Object
}

type Spec interface {
	Reconcile(oldSpec Spec)
	Instance
}

func Render(spec Spec, container *js.Object) {
	spec.Reconcile(nil)
	container.Call("appendChild", spec.Node())
}

func RenderAsBody(spec Spec) {
	body := js.Global.Get("document").Call("createElement", "body")
	Render(spec, body)
	js.Global.Get("document").Set("body", body)
}

type TextSpec struct {
	text string
	node *js.Object
}

func Text(text string) *TextSpec {
	return &TextSpec{text: text}
}

func (s *TextSpec) Apply(element *Element) {
	element.Children = append(element.Children, s)
}

func (s *TextSpec) Reconcile(oldSpec Spec) {
	if oldText, ok := oldSpec.(*TextSpec); ok {
		s.node = oldText.node
		if oldText.text != s.text {
			s.node.Set("nodeValue", s.text)
		}
		return
	}

	s.node = js.Global.Get("document").Call("createTextNode", s.text)
}

func (s *TextSpec) Node() *js.Object {
	return s.node
}

type Element struct {
	TagName        string
	Properties     []*Property
	Styles         []*Style
	EventListeners []*EventListener
	Children       []Spec
	node           *js.Object
}

func (e *Element) AddChild(s Spec) {
	e.Children = append(e.Children, s)
}

func (e *Element) Apply(element *Element) {
	element.Children = append(element.Children, e)
}

func (e *Element) Reconcile(oldSpec Spec) {
	if oldElement, ok := oldSpec.(*Element); ok && oldElement.TagName == e.TagName {
		e.node = oldElement.node
		// TODO update properties, etc.
		for i, newChild := range e.Children {
			oldChild := oldElement.Children[i]
			newChild.Reconcile(oldChild)
			domutil.ReplaceNode(newChild.Node(), oldChild.Node())
		}
		return
	}

	e.node = js.Global.Get("document").Call("createElement", e.TagName)
	for _, p := range e.Properties {
		e.node.Set(p.Name, p.Value)
	}
	style := e.node.Get("style")
	for _, s := range e.Styles {
		style.Call("setProperty", s.Name, s.Value)
	}
	for _, l := range e.EventListeners {
		e.node.Call("addEventListener", l.Name, func(jsEvent *js.Object) {
			l.Listener(&Event{Target: jsEvent.Get("target")})
		})
	}
	for _, c := range e.Children {
		c.Reconcile(nil)
		e.node.Call("appendChild", c.Node())
	}
}

func (e *Element) Node() *js.Object {
	return e.node
}

type Property struct {
	Name  string
	Value interface{}
}

func (p *Property) Apply(element *Element) {
	element.Properties = append(element.Properties, p)
}

type Style struct {
	Name  string
	Value interface{}
}

func (p *Style) Apply(element *Element) {
	element.Styles = append(element.Styles, p)
}

type EventListener struct {
	Name     string
	Listener func(*Event)
}

func (l *EventListener) Apply(element *Element) {
	element.EventListeners = append(element.EventListeners, l)
}

type Event struct {
	Target *js.Object
}
