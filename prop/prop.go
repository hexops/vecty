package prop

import "github.com/neelance/dom"

func Autofocus() dom.Aspect {
	return dom.Prop("autofocus", true)
}

func Checked() dom.Aspect {
	return dom.Prop("checked", true)
}

type classAspect struct {
	classes []string
}

func (a *classAspect) Apply(parent *dom.ElemAspect) {
	for _, c := range a.classes {
		parent.Node.Get("classList").Call("add", c)
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
