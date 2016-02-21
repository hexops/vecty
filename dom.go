package vecty

import "github.com/gopherjs/gopherjs/js"

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

type textComponent struct {
	text string
	node *js.Object
}

func Text(text string) Component {
	return &textComponent{text: text}
}

func (s *textComponent) Apply(element *Element) {
	element.Children = append(element.Children, s)
}

func (s *textComponent) Reconcile(oldComp Component) {
	if oldText, ok := oldComp.(*textComponent); ok {
		s.node = oldText.node
		if oldText.text != s.text {
			s.node.Set("nodeValue", s.text)
		}
		return
	}

	s.node = js.Global.Get("document").Call("createTextNode", s.text)
}

func (s *textComponent) Node() *js.Object {
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
	for _, l := range e.EventListeners {
		l.wrapper = func(jsEvent *js.Object) {
			if l.CallPreventDefault {
				jsEvent.Call("preventDefault")
			}
			l.Listener(&Event{Target: jsEvent.Get("target")})
		}
	}

	if oldElement, ok := oldComp.(*Element); ok && oldElement.TagName == e.TagName {
		e.node = oldElement.node
		for name, value := range e.Properties {
			oldValue := oldElement.Properties[name]
			if value != oldValue || name == "value" || name == "checked" {
				e.node.Set(name, value)
			}
		}
		for name := range oldElement.Properties {
			if _, ok := e.Properties[name]; !ok {
				e.node.Set(name, nil)
			}
		}

		// TODO fix style reset
		style := e.node.Get("style")
		for name, value := range e.Style {
			style.Call("setProperty", name, value)
		}

		for _, l := range oldElement.EventListeners {
			e.node.Call("removeEventListener", l.Name, l.wrapper)
		}
		for _, l := range e.EventListeners {
			e.node.Call("addEventListener", l.Name, l.wrapper)
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
			replaceNode(newChild.Node(), oldChild.Node())
		}
		for i := len(e.Children); i < len(oldElement.Children); i++ {
			removeNode(oldElement.Children[i].Node())
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
		e.node.Call("addEventListener", l.Name, l.wrapper)
	}
	for _, c := range e.Children {
		c.Reconcile(nil)
		e.node.Call("appendChild", c.Node())
	}
}

func (e *Element) Node() *js.Object {
	return e.node
}

// SetTitle sets the title of the document.
func SetTitle(title string) {
	js.Global.Get("document").Set("title", title)
}

// AddStylesheet adds an external stylesheet to the document.
func AddStylesheet(url string) {
	link := js.Global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	js.Global.Get("document").Get("head").Call("appendChild", link)
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
	if oldBody != nil {
		replaceNode(c.Body.Node(), oldBody.Node())
	}
}
