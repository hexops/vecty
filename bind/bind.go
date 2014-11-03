package bind

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/event"
)

type Scope struct {
	Listeners []func()
}

func (s *Scope) Digest() {
	for _, l := range s.Listeners {
		l()
	}
}

func NewScope() *Scope {
	return &Scope{}
}

type TextBinding struct {
	Ptr   *string
	Scope *Scope
	Text  *dom.TextAspect
}

func Text(ptr *string, scope *Scope) dom.Aspect {
	return &TextBinding{Ptr: ptr, Scope: scope, Text: dom.Text(*ptr)}
}

func (b *TextBinding) Apply(parent *dom.ElemAspect) {
	b.Text.Apply(parent)
	b.Scope.Listeners = append(b.Scope.Listeners, func() {
		b.Text.Node.Set("textContent", *b.Ptr)
	})
}

func InputValue(ptr *string, scope *Scope) *dom.EventAspect {
	return event.Input(func(e *dom.ElemAspect) {
		*ptr = e.Node.Get("value").Str()
		scope.Digest()
	})
}
