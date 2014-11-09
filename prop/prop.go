package prop

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
)

func Autofocus() dom.Aspect {
	return dom.ToggleProp("autofocus")
}

func Checked() dom.Aspect {
	return dom.ToggleProp("checked")
}

type classAspect struct {
	classes   []string
	classList js.Object
}

func (a *classAspect) Apply(node js.Object, p, r float64) {
	a.classList = node.Get("classList")
	for _, c := range a.classes {
		a.classList.Call("add", c)
	}
}

func (a *classAspect) Revert() {
	if a.classList == nil {
		return
	}
	for _, c := range a.classes {
		a.classList.Call("remove", c)
	}
}

func Class(classes ...string) dom.Aspect {
	return &classAspect{classes: classes}
}

func For(id string) dom.Aspect {
	return dom.SetProp("htmlFor", id)
}

func HRef(url string) dom.Aspect {
	return dom.SetProp("href", url)
}

func Id(id string) dom.Aspect {
	return dom.SetProp("id", id)
}

func Placeholder(text string) dom.Aspect {
	return dom.SetProp("placeholder", text)
}

func Src(url string) dom.Aspect {
	return dom.SetProp("src", url)
}

func Type(t string) dom.Aspect {
	return dom.SetProp("type", t)
}

func Value(v string) dom.Aspect {
	return dom.SetProp("value", v)
}
