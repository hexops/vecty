package style

import "github.com/neelance/dom"

func Width(pixels int) dom.Aspect {
	return dom.Style("width", pixels)
}

func Height(pixels int) dom.Aspect {
	return dom.Style("height", pixels)
}

func Color(value string) dom.Aspect {
	return dom.Style("color", value)
}
