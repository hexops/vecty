package prop

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
)

func Autofocus() dom.Aspect {
	return dom.Prop("autofocus", true)
}

func Checked() dom.Aspect {
	return dom.Prop("checked", true)
}

type classAspect struct {
	classes   []string
	classList js.Object
}

func (a *classAspect) Apply(node js.Object) {
	a.classList = node.Get("classList")
	for _, c := range a.classes {
		a.classList.Call("add", c)
	}
}

func (a *classAspect) Revoke() {
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
	return dom.Prop("htmlFor", id)
}

func HRef(url string) dom.Aspect {
	return dom.Prop("href", url)
}

func Id(id string) dom.Aspect {
	return dom.Prop("id", id)
}

func Placeholder(text string) dom.Aspect {
	return dom.Prop("placeholder", text)
}

func Src(url string) dom.Aspect {
	return dom.Prop("src", url)
}

func Type(t string) dom.Aspect {
	return dom.Prop("type", t)
}

func Value(v interface{}) dom.Aspect {
	return dom.Prop("value", v)
}
