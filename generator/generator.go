package main

import (
	"fmt"
	"io"
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
	"keygen":     "KeyGen",
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

	fmt.Fprint(file, `// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.
package elem

import (
	"github.com/neelance/dom"
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
func %s(aspects ...dom.Aspect) dom.Aspect {
	return dom.Element("%s", aspects...)
}
`, desc, link, funName, name)
}

func generateEventPkg() {
	nameMap := map[string]string{
		"afterprint":              "AfterPrint",
		"animationend":            "AnimationEnd",
		"animationiteration":      "AnimationIteration",
		"animationstart":          "AnimationStart",
		"audioprocess":            "AudioProcess",
		"beforeprint":             "BeforePrint",
		"beforeunload":            "BeforeUnload",
		"canplay":                 "CanPlay",
		"canplaythrough":          "CanPlayThrough",
		"chargingchange":          "ChargingChange",
		"chargingtimechange":      "ChargingTimeChange",
		"compassneedscalibration": "CompassNeedsCalibration",
		"compositionend":          "CompositionEnd",
		"compositionstart":        "CompositionStart",
		"compositionupdate":       "CompositionUpdate",
		"contextmenu":             "ContextMenu",
		"dblclick":                "DblClick",
		"devicelight":             "DeviceLight",
		"devicemotion":            "DeviceMotion",
		"deviceorientation":       "DeviceOrientation",
		"deviceproximity":         "DeviceProximity",
		"dischargingtimechange":   "DischargingTimeChange",
		"dragend":                 "DragEnd",
		"dragenter":               "DragEnter",
		"dragleave":               "DragLeave",
		"dragover":                "DragOver",
		"dragstart":               "DragStart",
		"durationchange":          "DurationChange",
		"focusin":                 "FocusIn",
		"focusout":                "FocusOut",
		"fullscreenchange":        "FullScreenChange",
		"fullscreenerror":         "FullScreenError",
		"gamepadconnected":        "GamepadConnected",
		"gamepaddisconnected":     "GamepadDisconnected",
		"hashchange":              "HashChange",
		"keydown":                 "KeyDown",
		"keypress":                "KeyPress",
		"keyup":                   "KeyUp",
		"languagechange":          "LanguageChange",
		"levelchange":             "LevelChange",
		"loadeddata":              "LoadedData",
		"loadedmetadata":          "LoadedMetadata",
		"loadend":                 "LoadEnd",
		"loadstart":               "LoadStart",
		"mousedown":               "MouseDown",
		"mouseenter":              "MouseEnter",
		"mouseleave":              "MouseLeave",
		"mousemove":               "MouseMove",
		"mouseout":                "MouseOut",
		"mouseover":               "MouseOver",
		"mouseup":                 "MouseUp",
		"noupdate":                "NoUpdate",
		"orientationchange":       "OrientationChange",
		"pagehide":                "PageHide",
		"pageshow":                "PageShow",
		"pointerlockchange":       "PointerLockChange",
		"pointerlockerror":        "PointerLockError",
		"popstate":                "PopState",
		"ratechange":              "RateChange",
		"readystatechange":        "ReadystateChange",
		"timeupdate":              "TimeUpdate",
		"touchcancel":             "TouchCancel",
		"touchend":                "TouchEnd",
		"touchenter":              "TouchEnter",
		"touchleave":              "TouchLeave",
		"touchmove":               "TouchMove",
		"touchstart":              "TouchStart",
		"transitionend":           "TransitionEnd",
		"updateready":             "UpdateReady",
		"upgradeneeded":           "UpgradeNeeded",
		"userproximity":           "UserProximity",
		"versionchange":           "VersionChange",
		"visibilitychange":        "VisibilityChange",
		"volumechange":            "VolumeChange",
	}

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
		e.Desc = strings.TrimSpace(cols.Eq(3).Text())
		if e.Desc == "" {
			e.Desc = "(no documentation)"
		}

		funName := nameMap[e.Name]
		if funName == "" {
			funName = capitalize(e.Name)
		}

		events[funName] = &e
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

	fmt.Fprint(file, `// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.
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
func %s(f dom.Listener) dom.Aspect {
	return dom.Event("%s", f)
}
`, e.Desc, e.Link[6:], name, e.Name)
	}
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
