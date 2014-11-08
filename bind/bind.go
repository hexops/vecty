package bind

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/prop"
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

type ifBinding struct {
	condition *bool
	scope     *Scope
	aspect    dom.Aspect
}

func If(condition *bool, scope *Scope, aspects ...dom.Aspect) dom.Aspect {
	return &ifBinding{condition: condition, scope: scope, aspect: dom.Group(aspects...)}
}

func (b *ifBinding) Apply(node js.Object) {
	l := func() {
		switch *b.condition {
		case true:
			b.aspect.Apply(node)
		case false:
			if ra, ok := b.aspect.(dom.RevokableAspect); ok {
				ra.Revoke()
			}
		}
	}
	b.scope.Listeners = append(b.scope.Listeners, l)
	l()
}

type Aspects struct {
	current  []dom.Aspect
	oldCache map[interface{}]dom.Aspect
	newCache map[interface{}]dom.Aspect
}

func (a *Aspects) Add(key interface{}, aspect dom.Aspect) {
	a.current = append(a.current, aspect)
	a.newCache[key] = aspect
}

func (a *Aspects) Reuse(key interface{}) bool {
	if cached, ok := a.oldCache[key]; ok {
		a.Add(key, cached)
		delete(a.oldCache, key)
		return true
	}
	return false
}

type dynamicAspect struct {
	scope *Scope
	fun   func(*Aspects)
	cache map[interface{}]dom.Aspect
}

func (d *dynamicAspect) Apply(node js.Object) {
	update := func() {
		aspects := &Aspects{
			oldCache: d.cache,
			newCache: make(map[interface{}]dom.Aspect),
		}

		d.fun(aspects)
		d.cache = aspects.newCache

		for _, a := range aspects.oldCache {
			if ra, ok := a.(dom.RevokableAspect); ok {
				ra.Revoke()
			}
		}

		for _, a := range aspects.current {
			a.Apply(node)
		}
	}
	update()
	d.scope.Listeners = append(d.scope.Listeners, update)
}

func Dynamic(scope *Scope, fun func(*Aspects)) dom.Aspect {
	return &dynamicAspect{
		scope: scope,
		fun:   fun,
	}
}

func Text(ptr *string, scope *Scope) dom.Aspect {
	return elem.Span(
		Dynamic(scope, func(aspects *Aspects) {
			aspects.Add("", dom.Text(*ptr))
		}),
	)
}

func InputValue(ptr *string, scope *Scope) dom.Aspect {
	return event.Input(func(c *dom.EventContext) {
		*ptr = c.Node.Get("value").Str()
		scope.Digest()
	})
}

func Checked(condition *bool, scope *Scope) dom.Aspect {
	return dom.Group(
		If(condition, scope,
			prop.Checked(),
		),
		event.Change(func(c *dom.EventContext) {
			*condition = c.Node.Get("checked").Bool()
			scope.Digest()
		}),
	)
}
