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

func InputValue(ptr *string, scope *Scope) (a *dom.EventAspect) {
	a = event.Input(func() {
		*ptr = a.Element.Node.Get("value").Str()
		scope.Digest()
	})
	return
}

type Aspects struct {
	Current  []dom.Aspect
	OldCache map[interface{}]dom.Aspect
	NewCache map[interface{}]dom.Aspect
}

func (a *Aspects) Add(key interface{}, aspect dom.Aspect) {
	a.Current = append(a.Current, aspect)
	a.NewCache[key] = aspect
}

func (a *Aspects) Reuse(key interface{}) bool {
	if cached, ok := a.OldCache[key]; ok {
		a.Add(key, cached)
		delete(a.OldCache, key)
		return true
	}
	return false
}

type DynamicAspect struct {
	Scope *Scope
	Fun   func(*Aspects)
	cache map[interface{}]dom.Aspect
}

func (d *DynamicAspect) Apply(parent *dom.ElemAspect) {
	update := func() {
		aspects := &Aspects{
			OldCache: d.cache,
			NewCache: make(map[interface{}]dom.Aspect),
		}

		d.Fun(aspects)
		d.cache = aspects.NewCache

		for _, a := range aspects.OldCache {
			if ra, ok := a.(dom.RevokableAspect); ok {
				ra.Revoke()
			}
		}

		for _, a := range aspects.Current {
			a.Apply(parent)
		}
	}
	update()
	d.Scope.Listeners = append(d.Scope.Listeners, update)
}

func Dynamic(scope *Scope, fun func(*Aspects)) *DynamicAspect {
	return &DynamicAspect{
		Scope: scope,
		Fun:   fun,
	}
}
