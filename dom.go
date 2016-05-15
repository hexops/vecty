package vecty

import (
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

// Renderable is a user-defined renderable component.
type Renderable interface {
	Component
	Render() Component
}

// Core is the struct which all components embed.
type Core struct {
	component Renderable
	body      Component
}

// Apply implements the Markup interface.
func (c *Core) Apply(e *Element) {
	e.AddChild(c.component)
}

// Unmount implements the Component interface.
func (c *Core) Unmount() {
	c.body.Unmount()
}

// Node implements the Component interface.
func (c *Core) Node() *js.Object {
	return c.body.Node()
}

// Reconcile implements the Component interface.
func (c *Core) Reconcile(prev Component) {
	// This function is the generic equivalent to:
	//
	// 	if p, ok := prev.(*UserComponent); ok {
	// 		c.body = p.body
	// 	}
	// 	b.Rerender()
	//
	if reflect.TypeOf(c.component) == reflect.TypeOf(prev) {
		prevCore := reflect.Indirect(reflect.ValueOf(prev)).FieldByName("Core").Interface().(*Core)
		c.body = prevCore.body
	}
	c.Rerender()
}

// Rerender causes the component to rerender the body and reconcile it.
func (c *Core) Rerender() {
	oldBody := c.body
	c.body = c.component.Render()
	c.body.Reconcile(oldBody)
	if oldBody != nil {
		replaceNode(c.body.Node(), oldBody.Node())
	}
}

// New creates a new Core component object.
func New(component Renderable) *Core {
	return &Core{component: component}
}

// Component represents a Vecty component.
type Component interface {
	Markup

	// Unmount is called before the component is unmounted. i.e., when all of
	// its event listeners and DOM elements will be removed.
	Unmount()

	Reconcile(oldComp Component)
	Node() *js.Object
}

// Render renders a component into the given container element. It is appended
// as a child element.
func Render(comp Component, container *js.Object) {
	comp.Reconcile(nil)
	container.Call("appendChild", comp.Node())
}

// RenderAsBody renders the given component as the body of the page, replacing
// whatever existing content in the page body there may be.
func RenderAsBody(comp Component) {
	body := js.Global.Get("document").Call("createElement", "body")
	Render(comp, body)
	js.Global.Get("document").Set("body", body)
}

type textComponent struct {
	text string
	node *js.Object
}

// Apply implements the Markup interface.
func (s *textComponent) Apply(element *Element) {
	element.Children = append(element.Children, s)
}

// Unmount implements the Component interface and is no-op.
func (s *textComponent) Unmount() {}

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

// Text returns a component which renders the given text. The text is always
// escaped, and as such feeding arbitrary user input to this function is safe.
func Text(text string) Component {
	return &textComponent{text: text}
}

// Element is a Component which virtually represents a DOM element.
type Element struct {
	TagName        string
	Properties     map[string]interface{}
	Style          map[string]interface{}
	Dataset        map[string]string
	EventListeners []*EventListener
	Children       []Component
	node           *js.Object
}

// AddChild adds a child component.
func (e *Element) AddChild(s Component) {
	e.Children = append(e.Children, s)
}

// Apply implements the Markup interface.
func (e *Element) Apply(element *Element) {
	element.Children = append(element.Children, e)
}

// Unmount unmounts this Element component by calling unmount on all children
// and removing all event listeners.
func (e *Element) Unmount() {
	for _, child := range e.Children {
		child.Unmount()
	}
	for _, l := range e.EventListeners {
		e.node.Call("removeEventListener", l.Name, l.wrapper)
	}
}

// Reconcile implements the Component interface.
func (e *Element) Reconcile(oldComp Component) {
	for _, l := range e.EventListeners {
		l.wrapper = func(jsEvent *js.Object) {
			if l.callPreventDefault {
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

		style := e.node.Get("style")
		for name, value := range e.Style {
			style.Call("setProperty", name, value)
		}
		for name := range oldElement.Style {
			if _, ok := e.Style[name]; !ok {
				style.Call("removeProperty", name)
			}
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
			oldElement.Children[i].Unmount()
			removeNode(oldElement.Children[i].Node())
		}
		return
	}

	e.node = js.Global.Get("document").Call("createElement", e.TagName)
	for name, value := range e.Properties {
		e.node.Set(name, value)
	}
	for name, value := range e.Dataset {
		e.node.Get("dataset").Set(name, value)
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

// Node implements the Component interface.
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
