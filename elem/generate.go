// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// elemNameMap translates lowercase HTML tag names from the MDN source into a
// proper Go style name with MixedCaps and initialisms:
//
//  https://github.com/golang/go/wiki/CodeReviewComments#mixed-caps
//  https://github.com/golang/go/wiki/CodeReviewComments#initialisms
//
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
	"h1":         "Heading1",
	"h2":         "Heading2",
	"h3":         "Heading3",
	"h4":         "Heading4",
	"h5":         "Heading5",
	"h6":         "Heading6",
	"hgroup":     "HeadingsGroup",
	"hr":         "HorizontalRule",
	"i":          "Italic",
	"iframe":     "InlineFrame",
	"img":        "Image",
	"ins":        "InsertedText",
	"kbd":        "KeyboardInput",
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

// Package elem defines markup to create DOM elements.
//
// Generated from "HTML element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under
// CC-BY-SA 2.5.
package elem

import "github.com/gopherjs/vecty"
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
		// Headings are grouped into one description so we generate them
		// individually here.
		if text == "<h1>" {
			writeElem(file, "h1", desc, link)
			writeElem(file, "h2", desc, link)
			writeElem(file, "h3", desc, link)
			writeElem(file, "h4", desc, link)
			writeElem(file, "h5", desc, link)
			writeElem(file, "h6", desc, link)
			return
		}

		name := text[1 : len(text)-1]
		if name == "html" || name == "head" || len(strings.Fields(name)) > 1 {
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

	// Replace a number of 'no-break space' unicode characters which exist in
	// the descriptions with normal spaces.
	desc = strings.Replace(desc, "\u00a0", " ", -1)

	// Reword the description so it is a proper Go comment sentence.
	switch name {
	case "audio":
		// MDN has no description for audio at the time of writing this.
		desc = `Audio is used to embed sound content in documents. It may contain one or more audio sources, represented using the src attribute or the <source> element: the browser will choose the most suitable one. It can also be the destination for streamed media, using a MediaStream.`
	case "h1", "h2", "h3", "h4", "h5", "h6":
		// Headings are grouped into one description
		n, err := strconv.Atoi(string(name[1]))
		if err != nil {
			log.Fatalf("elem: Failed to parse heading number from \"%s\"", name)
		}
		old := "The HTML <h1>–<h6> elements represent six levels of section headings"
		new := fmt.Sprintf("%s represents a level %v section heading", funName, n)
		desc = strings.Replace(desc, old, new, 1)
	default:
		// Descriptions for elements generally read as one of:
		//
		//  The HTML <foobar> element ...
		//  The HTML Full Name element (<foobar>) ...
		//
		// Because these are consistent (sometimes with varying captalization,
		// however) we can exploit that fact to reword the documentation in proper
		// Go style:
		//
		//  Foobar ...
		//
		s := strings.Split(desc, "<"+name+">")[1:]
		desc = strings.Join(s, "<"+name+">")
		exprs := []string{
			`^\s*\)?\s*(e|E)lement\s*`,
			`^\s*\)?\s*`,
		}
		for _, expr := range exprs {
			re := regexp.MustCompile(expr)
			if re.MatchString(desc) {
				desc = re.ReplaceAllString(desc, funName+" ")
				break
			}
		}
		// Suffix fixups
		desc = strings.TrimSpace(desc)
		if !strings.HasSuffix(desc, ".") {
			desc = desc[:len(desc)-1] + "."
		}
	}

	fmt.Fprintf(w, `%s
//
// https://developer.mozilla.org%s
func %s(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return vecty.Tag("%s", markup...)
}
`, descToComments(desc), link, funName, name)
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func descToComments(desc string) string {
	c := ""
	length := 80
	for _, word := range strings.Fields(desc) {
		if length+len(word)+1 > 80 {
			length = 3
			c += "\n//"
		}
		c += " " + word
		length += len(word) + 1
	}
	return c
}
