package bind

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/prop"
)

type Scope struct {
	age             int
	listeners       map[int]func()
	listenerCounter int
}

func NewScope() *Scope {
	return &Scope{
		age:             0,
		listeners:       make(map[int]func()),
		listenerCounter: 0,
	}
}

func (s *Scope) Age() int {
	return s.age
}

func (s *Scope) Digest() {
	s.age++
	for _, l := range s.listeners {
		l()
	}
}

type Listener struct {
	id    int
	fun   func()
	scope *Scope
}

func (s *Scope) NewListener(fun func()) *Listener {
	s.listenerCounter++
	s.listeners[s.listenerCounter] = fun
	return &Listener{id: s.listenerCounter, fun: fun, scope: s}
}

func (l *Listener) Call() {
	l.fun()
}

func (l *Listener) Remove() {
	delete(l.scope.listeners, l.id)
}

func (s *Scope) CacheString(fun func() string) func() string {
	cachedAge := -1
	cache := ""
	return func() string {
		age := s.Age()
		if age != cachedAge {
			cachedAge = age
			cache = fun()
		}
		return cache
	}
}

type ifAspect struct {
	condition *bool
	scope     *Scope
	aspect    dom.Aspect
	listener  *Listener
}

func If(condition *bool, scope *Scope, aspects ...dom.Aspect) dom.Aspect {
	return &ifAspect{condition: condition, scope: scope, aspect: dom.Group(aspects...)}
}

func (a *ifAspect) Apply(node js.Object) {
	if a.listener != nil {
		return
	}
	a.listener = a.scope.NewListener(func() {
		switch *a.condition {
		case true:
			a.aspect.Apply(node)
		case false:
			a.aspect.Revert()
		}
	})
	a.listener.Call()
}

func (a *ifAspect) Revert() {
	a.aspect.Revert()
	a.listener.Remove()
	a.listener = nil
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
	scope    *Scope
	fun      func(*Aspects)
	cache    map[interface{}]dom.Aspect
	listener *Listener
}

func (a *dynamicAspect) Apply(node js.Object) {
	if a.listener != nil {
		return
	}
	a.listener = a.scope.NewListener(func() {
		aspects := &Aspects{
			oldCache: a.cache,
			newCache: make(map[interface{}]dom.Aspect),
		}

		a.fun(aspects)
		a.cache = aspects.newCache

		for _, a := range aspects.oldCache {
			a.Revert()
		}

		for _, a := range aspects.current {
			a.Apply(node)
		}
	})
	a.listener.Call()
}

func (a *dynamicAspect) Revert() {
	for _, a := range a.cache {
		a.Revert()
	}
	a.listener.Remove()
	a.listener = nil
}

func Dynamic(scope *Scope, fun func(*Aspects)) dom.Aspect {
	return &dynamicAspect{
		scope: scope,
		fun:   fun,
	}
}

func TextFunc(fun func() string, scope *Scope) dom.Aspect {
	var current string
	return elem.Span(
		Dynamic(scope, func(aspects *Aspects) {
			text := fun()
			if text == current {
				aspects.Reuse("")
				return
			}
			aspects.Add("", dom.Text(text))
			current = text
		}),
	)
}

func TextPtr(ptr *string, scope *Scope) dom.Aspect {
	return TextFunc(func() string { return *ptr }, scope)
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
