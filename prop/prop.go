package prop

import "github.com/neelance/dom"

func ClassList(names ...string) *dom.PropAspect {
	return dom.Prop("classList", names)
}

func Type(t string) *dom.PropAspect {
	return dom.Prop("type", t)
}

func Value(v interface{}) *dom.PropAspect {
	return dom.Prop("value", v)
}
