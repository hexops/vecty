package prop

import "github.com/gopherjs/vecty"

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

func Autofocus(autofocus bool) *vecty.Prop {
	return &vecty.Prop{Name: "autofocus", Value: autofocus}
}

func Checked(checked bool) *vecty.Prop {
	return &vecty.Prop{Name: "checked", Value: checked}
}

func Class(class string) *vecty.Prop {
	return &vecty.Prop{Name: "className", Value: class}
}

func For(id string) *vecty.Prop {
	return &vecty.Prop{Name: "htmlFor", Value: id}
}

func Href(url string) *vecty.Prop {
	return &vecty.Prop{Name: "href", Value: url}
}

func Id(id string) *vecty.Prop {
	return &vecty.Prop{Name: "id", Value: id}
}

func Placeholder(text string) *vecty.Prop {
	return &vecty.Prop{Name: "placeholder", Value: text}
}

func Src(url string) *vecty.Prop {
	return &vecty.Prop{Name: "src", Value: url}
}

func Type(t InputType) *vecty.Prop {
	return &vecty.Prop{Name: "type", Value: string(t)}
}

func Value(v string) *vecty.Prop {
	return &vecty.Prop{Name: "value", Value: v}
}
