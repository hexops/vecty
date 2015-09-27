// GENERATED, DO NOT CHANGE

package impl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/componentutil"
	"github.com/neelance/dom/example/components/spec"
)

func init() {
	spec.ReconcileGreeter = reconcileGreeter
}

type GreeterAccessors interface {
	Props() GreeterProps
	State() GreeterState
	Node() *js.Object
}

type GreeterProps interface {
	Greeting() string
}

type GreeterState interface {
	Name() string
	SetName(name string)
}

type greeterCore struct {
	componentutil.Core

	_greeting string

	_name string
}

func (c *greeterCore) Props() GreeterProps {
	return c
}

func (c *greeterCore) Greeting() string {
	return c._greeting
}

func (c *greeterCore) State() GreeterState {
	return c
}

func (c *greeterCore) Name() string {
	return c._name
}

func (c *greeterCore) SetName(name string) {
	c._name = name
	c.Update()
}

func reconcileGreeter(newSpec *spec.Greeter, oldSpec dom.Spec) {
	if oldSpec, ok := oldSpec.(*spec.Greeter); ok {
		newSpec.Instance = oldSpec.Instance
		newSpec.Instance.(*GreeterImpl).GreeterAccessors.(*greeterCore).DoRender()
		return
	}

	c := &greeterCore{
		_greeting: newSpec.Greeting,
	}
	inst := &GreeterImpl{c}
	c.Lifecycle = inst
	newSpec.Instance = inst
	c.DoRender()
}
