package style

import (
	"strconv"

	"github.com/gopherjs/vecty"
)

type Size string

func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

func Color(value string) *dom.Style {
	return &dom.Style{Name: "color", Value: value}
}

func Height(size Size) *dom.Style {
	return &dom.Style{Name: "height", Value: string(size)}
}

func Margin(size Size) *dom.Style {
	return &dom.Style{Name: "margin", Value: string(size)}
}

func Width(size Size) *dom.Style {
	return &dom.Style{Name: "width", Value: string(size)}
}
