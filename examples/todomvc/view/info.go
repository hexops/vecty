package view

import (
	"github.com/neelance/dom"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/prop"
)

func info() dom.Aspect {
	return elem.Footer(
		prop.Id("info"),

		elem.Paragraph(
			dom.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			dom.Text("Template by "),
			elem.Anchor(
				prop.Href("http://github.com/sindresorhus"),
				dom.Text("Sindre Sorhus"),
			),
		),
		elem.Paragraph(
			dom.Text("Created by "),
			elem.Anchor(
				prop.Href("http://github.com/neelance"),
				dom.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			dom.Text("Part of "),
			elem.Anchor(
				prop.Href("http://todomvc.com"),
				dom.Text("TodoMVC"),
			),
		),
	)
}
