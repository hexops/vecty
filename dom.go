package dom

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom/domutil"
)

type Markup interface {
	Apply(element *Element)
}

type Component interface {
	Markup
	Reconcile(oldComp Component)
	Node() *js.Object
}

func Render(comp Component, container *js.Object) {
	comp.Reconcile(nil)
	container.Call("appendChild", comp.Node())
}

func RenderAsBody(comp Component) {
	body := js.Global.Get("document").Call("createElement", "body")
	Render(comp, body)
	js.Global.Get("document").Set("body", body)
}

type TextComponent struct {
	text string
	node *js.Object
}

func Text(text string) *TextComponent {
	return &TextComponent{text: text}
}

func (s *TextComponent) Apply(element *Element) {
	element.Children = append(element.Children, s)
}

func (s *TextComponent) Reconcile(oldComp Component) {
	if oldText, ok := oldComp.(*TextComponent); ok {
		s.node = oldText.node
		if oldText.text != s.text {
			s.node.Set("nodeValue", s.text)
		}
		return
	}

	s.node = js.Global.Get("document").Call("createTextNode", s.text)
}

func (s *TextComponent) Node() *js.Object {
	return s.node
}

type Element struct {
	TagName        string
	Properties     map[string]interface{}
	Style          map[string]interface{}
	EventListeners []*EventListener
	Children       []Component
	node           *js.Object
}

func (e *Element) AddChild(s Component) {
	e.Children = append(e.Children, s)
}

func (e *Element) Apply(element *Element) {
	element.Children = append(element.Children, e)
}

func (e *Element) Reconcile(oldComp Component) {
	if oldElement, ok := oldComp.(*Element); ok && oldElement.TagName == e.TagName {
		e.node = oldElement.node
		for name, value := range e.Properties {
			oldValue := oldElement.Properties[name]
			if name == "value" { // TODO is there a nicer way than to compial case this?
				oldValue = e.node.Get("value").String()
			}
			if value != oldValue {
				e.node.Set(name, value)
			}
		}
		for name, _ := range oldElement.Properties {
			if _, ok := e.Properties[name]; !ok {
				e.node.Set(name, nil)
			}
		}

		// TODO fix style reset
		style := e.node.Get("style")
		for name, value := range e.Style {
			style.Call("setProperty", name, value)
		}

		// TODO better list element reuse
		for i, newChild := range e.Children {
			if i >= len(oldElement.Children) {
				newChild.Reconcile(nil)
				e.node.Call("appendChild", newChild.Node())
				continue
			}
			oldChild := oldElement.Children[i]
			newChild.Reconcile(oldChild)
			domutil.ReplaceNode(newChild.Node(), oldChild.Node())
		}
		for i := len(e.Children); i < len(oldElement.Children); i++ {
			domutil.RemoveNode(oldElement.Children[i].Node())
		}
		return
	}

	e.node = js.Global.Get("document").Call("createElement", e.TagName)
	for name, value := range e.Properties {
		e.node.Set(name, value)
	}
	style := e.node.Get("style")
	for name, value := range e.Style {
		style.Call("setProperty", name, value)
	}
	for _, l := range e.EventListeners {
		e.node.Call("addEventListener", l.Name, func(jsEvent *js.Object) {
			if l.CallPreventDefault {
				jsEvent.Call("preventDefault")
			}
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
	if element.Properties == nil {
		element.Properties = make(map[string]interface{})
	}
	if _, ok := element.Properties[p.Name]; ok {
		panic(fmt.Sprintf("duplicate property: %s", p.Name))
	}
	element.Properties[p.Name] = p.Value
}

type ClassMap map[string]bool

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
}

func (l *EventListener) PreventDefault() *EventListener {
	l.CallPreventDefault = true
	return l
}

func (l *EventListener) Apply(element *Element) {
	element.EventListeners = append(element.EventListeners, l)
}

type Event struct {
	Target *js.Object
}

// SetTitle sets the title of the document.
func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}

// AddStylesheed adds an external stylesheet to the document.
func AddStylesheet(url string) {
	link := js.Global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	js.Global.Get("document").Get("head").Call("appendChild", link)
}

type List []Markup

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

type Composite struct {
	RenderFunc func() Component
	Body       Component
}

func (c *Composite) Node() *js.Object {
	return c.Body.Node()
}

func (c *Composite) ReconcileBody() {
	oldBody := c.Body
	c.Body = c.RenderFunc()
	c.Body.Reconcile(oldBody)
}
