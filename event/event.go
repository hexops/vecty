package event

import (
	"github.com/neelance/dom"
)

func Input(f func(*dom.ElemAspect)) *dom.EventAspect {
	return dom.Event("input", f)
}

func KeyDown(f func(*dom.ElemAspect)) *dom.EventAspect {
	return dom.Event("keydown", f)
}
