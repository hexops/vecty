package style

import "github.com/neelance/dom"

func Width(pixels int) *dom.StyleAspect {
	return dom.Style("width", pixels)
}

func Height(pixels int) *dom.StyleAspect {
	return dom.Style("height", pixels)
}

func Color(value string) *dom.StyleAspect {
	return dom.Style("color", value)
}
