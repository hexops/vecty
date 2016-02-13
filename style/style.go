package style

import (
	"strconv"

	"github.com/gopherjs/vecty"
)

type Size string

func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

func Color(value string) vecty.Markup {
	return vecty.Style("color", value)
}

func Height(size Size) vecty.Markup {
	return vecty.Style("height", string(size))
}

func Margin(size Size) vecty.Markup {
	return vecty.Style("margin", string(size))
}

func Width(size Size) vecty.Markup {
	return vecty.Style("width", string(size))
}
