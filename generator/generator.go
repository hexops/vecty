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
	nameMap := map[string]string{
		"bdi":        "BDI",
		"bdo":        "BDO",
		"bgsound":    "BgSound",
		"blockquote": "BlockQuote",
		"colgroup":   "ColGroup",
		"datalist":   "DataList",
		"dl":         "DL",
		"dt":         "DT",
		"fieldset":   "FieldSet",
		"figcaption": "FigCaption",
		"hr":         "HR",
		"iframe":     "IFrame",
		"keygen":     "KeyGen",
		"li":         "LI",
		"menuitem":   "MenuItem",
		"nobr":       "NoBr",
		"noscript":   "NoScript",
		"ol":         "OL",
		"rp":         "RP",
		"rt":         "RT",
		"tbody":      "TBody",
		"textarea":   "TextArea",
		"td":         "TD",
		"tfoot":      "TFoot",
		"th":         "TH",
		"thead":      "THead",
		"tr":         "TR",
		"ul":         "UL",
		"wbr":        "WBr",
	}

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

		funName := nameMap[name]
		if funName == "" {
			funName = capitalize(name)
		}

		desc, _ := s.Attr("title")
		link, _ := s.Attr("href")

		fmt.Fprintf(file, `
// %s
// https://developer.mozilla.org%s
func %s(aspects ...dom.Aspect) *dom.ElemAspect {
	return dom.Elem("%s", aspects...)
}
`, desc, link[6:], funName, name)
	})
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
