package elem

import (
	"github.com/neelance/dom"
)

func Div(mutators ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("div", mutators...)
}

func H1(mutators ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("h1", mutators...)
}

func Input(mutators ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("input", mutators...)
}
