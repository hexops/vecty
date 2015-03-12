package prop

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
)

type InputType string

const (
	TypeButton        InputType = "button"
	TypeCheckbox                = "checkbox"
	TypeColor                   = "color"
	TypeDate                    = "date"
	TypeDatetime                = "datetime"
	TypeDatetimeLocal           = "datetime-local"
	TypeEmail                   = "email"
	TypeFile                    = "file"
	TypeHidden                  = "hidden"
	TypeImage                   = "image"
	TypeMonth                   = "month"
	TypeNumber                  = "number"
	TypePassword                = "password"
	TypeRadio                   = "radio"
	TypeRange                   = "range"
	TypeMin                     = "min"
	TypeMax                     = "max"
	TypeValue                   = "value"
	TypeStep                    = "step"
	TypeReset                   = "reset"
	TypeSearch                  = "search"
	TypeSubmit                  = "submit"
	TypeTel                     = "tel"
	TypeText                    = "text"
	TypeTime                    = "time"
	TypeUrl                     = "url"
	TypeWeek                    = "week"
)

func Autofocus() dom.Aspect {
	return dom.ToggleProperty("autofocus")
}

func Checked() dom.Aspect {
	return dom.ToggleProperty("checked")
}

type classAspect struct {
	classes   []string
	classList *js.Object
}

func (a *classAspect) Apply(node *js.Object, p, r float64) {
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
	return dom.SetProperty("htmlFor", id)
}

func Href(url string) dom.Aspect {
	return dom.SetProperty("href", url)
}

func Id(id string) dom.Aspect {
	return dom.SetProperty("id", id)
}

func Placeholder(text string) dom.Aspect {
	return dom.SetProperty("placeholder", text)
}

func Src(url string) dom.Aspect {
	return dom.SetProperty("src", url)
}

func Type(t InputType) dom.Aspect {
	return dom.SetProperty("type", string(t))
}

func Value(v string) dom.Aspect {
	return dom.SetProperty("value", v)
}
