package event

import (
	"github.com/neelance/dom"
)

func Click(f func()) *dom.EventAspect {
	return dom.Event("click", f)
}

func Input(f func()) *dom.EventAspect {
	return dom.Event("input", f)
}

func KeyDown(f func()) *dom.EventAspect {
	return dom.Event("keydown", f)
}
