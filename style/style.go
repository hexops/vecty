package style

import (
	"strconv"

	"github.com/neelance/dom"
)

type Size string

func Px(pixels int) Size {
	return Size(strconv.Itoa(pixels) + "px")
}

func Color(value string) dom.Aspect {
	return dom.Style("color", value)
}

func Height(pixels Size) dom.Aspect {
	return dom.Style("height", string(pixels))
}

func Margin(pixels Size) dom.Aspect {
	return dom.Style("margin", string(pixels))
}

func Width(pixels Size) dom.Aspect {
	return dom.Style("width", string(pixels))
}
