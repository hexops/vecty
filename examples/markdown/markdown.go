package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	vecty.SetTitle("Markdown Demo")
	vecty.RenderBody(&PageView{
		Input: `# Markdown Example

This is a live editor, try editing the Markdown on the right of the page.
`,
	})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
	Input string
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() *vecty.HTML {
	return elem.Body(
		// Display a textarea on the right-hand side of the page.
		elem.Div(
			vecty.Style("float", "right"),
			elem.TextArea(
				vecty.Style("font-family", "monospace"),
				vecty.Property("rows", 14),
				vecty.Property("cols", 70),
				vecty.Text(p.Input), // initial textarea text.

				// When input is typed into the textarea, update the local
				// component state and rerender.
				event.Input(func(e *vecty.Event) {
					p.Input = e.Target.Get("value").String()
					vecty.Rerender(p)
				}),
			),
		),

		// Render the markdown.
		&Markdown{Input: p.Input},
	)
}

// Markdown is a simple component which renders the Input markdown as sanitized
// HTML into a div.
type Markdown struct {
	vecty.Core
	Input string
}

// Render implements the vecty.Component interface.
func (m *Markdown) Render() *vecty.HTML {
	// Render the markdown input into HTML using Blackfriday.
	unsafeHTML := blackfriday.MarkdownCommon([]byte(m.Input))

	// Sanitize the HTML.
	safeHTML := string(bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML))

	// Return the HTML, which we guarantee to be safe / sanitized.
	return elem.Div(
		vecty.UnsafeHTML(safeHTML),
	)
}
