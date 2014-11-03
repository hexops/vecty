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
	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/Events")
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
		events[strings.ToUpper(e.Name[:1])+e.Name[1:]] = &e
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

	fmt.Fprint(file, `// Documentation source: 
// "Event reference" (https://developer.mozilla.org/en-US/docs/Web/Events) by Mozilla Contributors, licensed under CC-BY-SA 2.5.
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
`, e.Desc, e.Link, name, e.Name)
	}
}
