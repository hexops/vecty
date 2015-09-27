//go:generate gencomponent github.com/neelance/dom/example/components/spec

package impl

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/event"
	"github.com/neelance/dom/prop"
)

type GreeterImpl struct {
	GreeterAccessors
}

func (s *GreeterImpl) onInput(event *dom.Event) {
	s.State().SetName(event.Target.Get("value").String())
}

func (s *GreeterImpl) Render() dom.Spec {
	return elem.Div(
		elem.Header1(
			dom.Text(s.Props().Greeting()+" "+s.State().Name()),
		),

		dom.Text("Your name: "),
		elem.Input(
			prop.Value(s.State().Name()),
			event.Input(s.onInput),
		),
	)
}
