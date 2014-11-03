package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Event struct {
	Name string
	Link string
	Desc string
}

func main() {
	generateElemPkg()
	generateEventPkg()
}

func generateElemPkg() {
	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/HTML/Element")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../elem/elem.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.
package elem

import (
  "github.com/neelance/dom"
)
`)

	doc.Find(".index a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("code").Text()
		if name == "" || s.Parent().Is(".obsoleteElement") {
			return
		}
		name = name[1 : len(name)-1]
		if name == "html" || name == "head" || name == "body" {
			return
		}
		desc, _ := s.Attr("title")
		link, _ := s.Attr("href")
		fmt.Fprintf(file, `
// %s
// https://developer.mozilla.org%s
func %s(aspects ...dom.Aspect) *dom.ElemAspect {
  return dom.Elem("%s", aspects...)
}
`, desc, link[6:], capitalize(name), name)
	})
}

func generateEventPkg() {
	doc, err := goquery.NewDocument("https://developer.mozilla.org/docs/Web/Events")
	if err != nil {
		panic(err)
	}

	events := make(map[string]*Event)

	doc.Find(".standard-table").Eq(0).Find("tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		if cols.Length() == 0 || cols.Find(".icon-thumbs-down-alt").Length() != 0 {
			return
		}
		link := cols.Eq(0).Find("a").Eq(0)
		var e Event
		e.Name = link.Text()
		e.Link, _ = link.Attr("href")
		e.Desc = cols.Eq(3).Text()
		events[capitalize(e.Name)] = &e
	})

	var names []string
	for name := range events {
		names = append(names, name)
	}
	sort.Strings(names)

	file, err := os.Create("../event/event.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/docs/Web/Events, licensed under CC-BY-SA 2.5.
package event

import (
  "github.com/neelance/dom"
)
`)

	for _, name := range names {
		e := events[name]
		fmt.Fprintf(file, `
// %s
// https://developer.mozilla.org%s
func %s(f func()) *dom.EventAspect {
  return dom.Event("%s", f)
}
`, e.Desc, e.Link[6:], name, e.Name)
	}
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
