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

func Autofocus(autofocus bool) vecty.Markup {
	return vecty.Property("autofocus", autofocus)
}

func Checked(checked bool) vecty.Markup {
	return vecty.Property("checked", checked)
}

func Class(class string) vecty.Markup {
	return vecty.Property("className", class)
}

func For(id string) vecty.Markup {
	return vecty.Property("htmlFor", id)
}

func Href(url string) vecty.Markup {
	return vecty.Property("href", url)
}

func ID(id string) vecty.Markup {
	return vecty.Property("id", id)
}

func Placeholder(text string) vecty.Markup {
	return vecty.Property("placeholder", text)
}

func Src(url string) vecty.Markup {
	return vecty.Property("src", url)
}

func Type(t InputType) vecty.Markup {
	return vecty.Property("type", string(t))
}

func Value(v string) vecty.Markup {
	return vecty.Property("value", v)
}
