package bind

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
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

func (s *Scope) CacheInt(fun func() int) func() int {
	cachedAge := -1
	cache := 0
	return func() int {
		age := s.Age()
		if age != cachedAge {
			cachedAge = age
			cache = fun()
		}
		return cache
	}
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

func (a *dynamicAspect) Apply(node *js.Object, p, r float64) {
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

		dom.Group(aspects.current...).Apply(node, p, r)
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

func IfFunc(condition func() bool, scope *Scope, aspects ...dom.Aspect) dom.Aspect {
	return Dynamic(scope, func(a *Aspects) {
		if condition() {
			if !a.Reuse("") {
				a.Add("", dom.Group(aspects...))
			}
		}
	})
}

func IfPtr(condition *bool, scope *Scope, aspects ...dom.Aspect) dom.Aspect {
	return IfFunc(func() bool { return *condition }, scope, aspects...)
}

func TextFunc(fun func() string, scope *Scope) dom.Aspect {
	current := ""
	return Dynamic(scope, func(aspects *Aspects) {
		text := fun()
		if text == current {
			aspects.Reuse("")
			return
		}
		aspects.Add("", dom.Text(text))
		current = text
	})
}

func TextPtr(ptr *string, scope *Scope) dom.Aspect {
	return TextFunc(func() string { return *ptr }, scope)
}

func Value(ptr *string, scope *Scope) dom.Aspect {
	return dom.Group(
		Dynamic(scope, func(aspects *Aspects) {
			aspects.Add("", prop.Value(*ptr))
		}),
		event.Input(func(c *dom.EventContext) {
			*ptr = c.Node.Get("value").String()
			scope.Digest()
		}),
	)
}

func Checked(condition *bool, scope *Scope) dom.Aspect {
	return dom.Group(
		IfPtr(condition, scope,
			prop.Checked(),
		),
		event.Change(func(c *dom.EventContext) {
			*condition = c.Node.Get("checked").Bool()
			scope.Digest()
		}),
	)
}

func Itoa(fun func() int) func() string {
	return func() string {
		return strconv.Itoa(fun())
	}
}
