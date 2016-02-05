package style

import (
	"strconv"

	"github.com/gopherjs/vecty"
)

type Size string

func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

func Color(value string) *vecty.Style {
	return &vecty.Style{Name: "color", Value: value}
}

func Height(size Size) *vecty.Style {
	return &vecty.Style{Name: "height", Value: string(size)}
}

func Margin(size Size) *vecty.Style {
	return &vecty.Style{Name: "margin", Value: string(size)}
}

func Width(size Size) *vecty.Style {
	return &vecty.Style{Name: "width", Value: string(size)}
}
