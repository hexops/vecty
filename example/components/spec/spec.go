package spec

import "github.com/neelance/dom"

type Greeter struct {
	dom.Instance

	Greeting string // uppercase -> property
	name     string // lowercase -> state
}

func (s *Greeter) Apply(element *dom.Element) { element.AddChild(s) }
func (s *Greeter) Reconcile(oldSpec dom.Spec) { ReconcileGreeter(s, oldSpec) }

var ReconcileGreeter func(newSpec *Greeter, oldSpec dom.Spec)
