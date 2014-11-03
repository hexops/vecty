package elem

import (
	"github.com/neelance/dom"
)

func Button(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("button", aspects...)
}

func Div(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("div", aspects...)
}

func H1(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("h1", aspects...)
}

func Input(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("input", aspects...)
}

func Li(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("li", aspects...)
}

func Ul(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("ul", aspects...)
}
