package prop

import "github.com/neelance/dom"

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

func Autofocus(autofocus bool) *dom.Property {
	return &dom.Property{Name: "autofocus", Value: autofocus}
}

func Checked(checked bool) *dom.Property {
	return &dom.Property{Name: "checked", Value: checked}
}

func Class(class string) *dom.Property {
	return &dom.Property{Name: "className", Value: class}
}

func For(id string) *dom.Property {
	return &dom.Property{Name: "htmlFor", Value: id}
}

func Href(url string) *dom.Property {
	return &dom.Property{Name: "href", Value: url}
}

func Id(id string) *dom.Property {
	return &dom.Property{Name: "id", Value: id}
}

func Placeholder(text string) *dom.Property {
	return &dom.Property{Name: "placeholder", Value: text}
}

func Src(url string) *dom.Property {
	return &dom.Property{Name: "src", Value: url}
}

func Type(t InputType) *dom.Property {
	return &dom.Property{Name: "type", Value: string(t)}
}

func Value(v string) *dom.Property {
	return &dom.Property{Name: "value", Value: v}
}
