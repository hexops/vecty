// +build ignore

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var elemNameMap = map[string]string{
	"a":          "Anchor",
	"abbr":       "Abbreviation",
	"b":          "Bold",
	"bdi":        "BidirectionalIsolation",
	"bdo":        "BidirectionalOverride",
	"blockquote": "BlockQuote",
	"br":         "Break",
	"cite":       "Citation",
	"col":        "Column",
	"colgroup":   "ColumnGroup",
	"datalist":   "DataList",
	"dd":         "Description",
	"del":        "DeletedText",
	"dfn":        "Definition",
	"dl":         "DescriptionList",
	"dt":         "DefinitionTerm",
	"em":         "Emphasis",
	"fieldset":   "FieldSet",
	"figcaption": "FigureCaption",
	"h1":         "Header1",
	"h2":         "Header2",
	"h3":         "Header3",
	"h4":         "Header4",
	"h5":         "Header5",
	"h6":         "Header6",
	"hgroup":     "HeadingsGroup",
	"hr":         "HorizontalRule",
	"i":          "Italic",
	"iframe":     "InlineFrame",
	"img":        "Image",
	"ins":        "InsertedText",
	"kbd":        "KeyboardInput",
	"keygen":     "KeyGeneration",
	"li":         "ListItem",
	"menuitem":   "MenuItem",
	"nav":        "Navigation",
	"noframes":   "NoFrames",
	"noscript":   "NoScript",
	"ol":         "OrderedList",
	"optgroup":   "OptionsGroup",
	"p":          "Paragraph",
	"param":      "Parameter",
	"pre":        "Preformatted",
	"q":          "Quote",
	"rp":         "RubyParenthesis",
	"rt":         "RubyText",
	"rtc":        "RubyTextContainer",
	"s":          "Strikethrough",
	"samp":       "Sample",
	"sub":        "Subscript",
	"sup":        "Superscript",
	"tbody":      "TableBody",
	"textarea":   "TextArea",
	"td":         "TableData",
	"tfoot":      "TableFoot",
	"th":         "TableHeader",
	"thead":      "TableHead",
	"tr":         "TableRow",
	"u":          "Underline",
	"ul":         "UnorderedList",
	"var":        "Variable",
	"wbr":        "WordBreakOpportunity",
}

func main() {
	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/HTML/Element")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("elem.gen.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `//go:generate go run generate.go

// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.
package elem

import (
	"github.com/gopherjs/vecty"
)
`)

	doc.Find(".quick-links a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if !strings.HasPrefix(link, "/en-US/docs/Web/HTML/Element/") {
			return
		}

		if s.Parent().Find(".icon-trash, .icon-thumbs-down-alt, .icon-warning-sign").Length() > 0 {
			return
		}

		desc, _ := s.Attr("title")

		text := s.Text()
		if text == "Heading elements" {
			writeElem(file, "h1", desc, link)
			writeElem(file, "h2", desc, link)
			writeElem(file, "h3", desc, link)
			writeElem(file, "h4", desc, link)
			writeElem(file, "h5", desc, link)
			writeElem(file, "h6", desc, link)
			return
		}

		name := text[1 : len(text)-1]
		if name == "html" || name == "head" || name == "body" {
			return
		}

		writeElem(file, name, desc, link)
	})
}

func writeElem(w io.Writer, name, desc, link string) {
	funName := elemNameMap[name]
	if funName == "" {
		funName = capitalize(name)
	}

	fmt.Fprintf(w, `
// %s
// https://developer.mozilla.org%s
func %s(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "%s"}
	vecty.List(markup).Apply(e)
	return e
}
`, desc, link, funName, name)
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
