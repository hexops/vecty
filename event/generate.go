// +build ignore

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
	Spec string
}

func main() {
	// nameMap translates lowercase HTML attribute names from the MDN source
	// into a proper Go style name with MixedCaps and initialisms:
	//
	//  https://github.com/golang/go/wiki/CodeReviewComments#mixed-caps
	//  https://github.com/golang/go/wiki/CodeReviewComments#initialisms
	//
	nameMap := map[string]string{
		"afterprint":               "AfterPrint",
		"animationend":             "AnimationEnd",
		"animationiteration":       "AnimationIteration",
		"animationstart":           "AnimationStart",
		"appinstalled":             "ApplicationInstalled",
		"audioprocess":             "AudioProcess",
		"audioend":                 "AudioEnd",
		"audiostart":               "AudioStart",
		"beforeprint":              "BeforePrint",
		"beforeunload":             "BeforeUnload",
		"canplay":                  "CanPlay",
		"canplaythrough":           "CanPlayThrough",
		"chargingchange":           "ChargingChange",
		"chargingtimechange":       "ChargingTimeChange",
		"compassneedscalibration":  "CompassNeedsCalibration",
		"compositionend":           "CompositionEnd",
		"compositionstart":         "CompositionStart",
		"compositionupdate":        "CompositionUpdate",
		"contextmenu":              "ContextMenu",
		"dblclick":                 "DoubleClick",
		"devicechange":             "DeviceChange",
		"devicelight":              "DeviceLight",
		"devicemotion":             "DeviceMotion",
		"deviceorientation":        "DeviceOrientation",
		"deviceproximity":          "DeviceProximity",
		"dischargingtimechange":    "DischargingTimeChange",
		"dragend":                  "DragEnd",
		"dragenter":                "DragEnter",
		"dragleave":                "DragLeave",
		"dragover":                 "DragOver",
		"dragstart":                "DragStart",
		"durationchange":           "DurationChange",
		"focusin":                  "FocusIn",
		"focusout":                 "FocusOut",
		"fullscreenchange":         "FullScreenChange",
		"fullscreenerror":          "FullScreenError",
		"gamepadconnected":         "GamepadConnected",
		"gamepaddisconnected":      "GamepadDisconnected",
		"gotpointercapture":        "GotPointerCapture",
		"hashchange":               "HashChange",
		"keydown":                  "KeyDown",
		"keypress":                 "KeyPress",
		"keyup":                    "KeyUp",
		"languagechange":           "LanguageChange",
		"levelchange":              "LevelChange",
		"loadeddata":               "LoadedData",
		"loadedmetadata":           "LoadedMetadata",
		"loadend":                  "LoadEnd",
		"loadstart":                "LoadStart",
		"lostpointercapture":       "LostPointerCapture",
		"messageerror":             "MessageError",
		"mousedown":                "MouseDown",
		"mouseenter":               "MouseEnter",
		"mouseleave":               "MouseLeave",
		"mousemove":                "MouseMove",
		"mouseout":                 "MouseOut",
		"mouseover":                "MouseOver",
		"mouseup":                  "MouseUp",
		"noupdate":                 "NoUpdate",
		"nomatch":                  "NoMatch",
		"notificationclick":        "NotificationClick",
		"orientationchange":        "OrientationChange",
		"pagehide":                 "PageHide",
		"pageshow":                 "PageShow",
		"pointercancel":            "PointerCancel",
		"pointerdown":              "PointerDown",
		"pointerenter":             "PointerEnter",
		"pointerleave":             "PointerLeave",
		"pointerlockchange":        "PointerLockChange",
		"pointerlockerror":         "PointerLockError",
		"pointermove":              "PointerMove",
		"pointerout":               "PointerOut",
		"pointerover":              "PointerOver",
		"pointerup":                "PointerUp",
		"popstate":                 "PopState",
		"pushsubscriptionchange":   "PushSubscriptionChange",
		"ratechange":               "RateChange",
		"readystatechange":         "ReadyStateChange",
		"resourcetimingbufferfull": "ResourceTimingBufferFull",
		"selectstart":              "SelectStart",
		"selectionchange":          "SelectionChange",
		"slotchange":               "SlotChange",
		"soundend":                 "SoundEnd",
		"soundstart":               "SoundStart",
		"speechend":                "SpeechEnd",
		"speechstart":              "SpeechStart",
		"timeupdate":               "TimeUpdate",
		"touchcancel":              "TouchCancel",
		"touchend":                 "TouchEnd",
		"touchenter":               "TouchEnter",
		"touchleave":               "TouchLeave",
		"touchmove":                "TouchMove",
		"touchstart":               "TouchStart",
		"transitionend":            "TransitionEnd",
		"updateready":              "UpdateReady",
		"upgradeneeded":            "UpgradeNeeded",
		"userproximity":            "UserProximity",
		"versionchange":            "VersionChange",
		"visibilitychange":         "VisibilityChange",
		"voiceschanged":            "VoicesChanged",
		"volumechange":             "VolumeChange",
		"vrdisplayconnected":       "VRDisplayConnected",
		"vrdisplaydisconnected":    "VRDisplayDisconnected",
		"vrdisplaypresentchange":   "VRDisplayPresentChange",
	}

	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/Events")
	if err != nil {
		panic(err)
	}

	events := make(map[string]*Event)

	doc.Find("#Standard_events+p+.standard-table").Eq(0).Find("tr").Each(func(i int, s *goquery.Selection) {
		cols := s.Find("td")
		if cols.Length() == 0 || cols.Find(".icon-thumbs-down-alt").Length() != 0 {
			return
		}
		link := cols.Eq(0).Find("a").Eq(0)
		var e Event
		e.Name = link.Text()
		e.Link, _ = link.Attr("href")
		e.Desc = strings.TrimSpace(cols.Eq(3).Text())
		e.Spec = strings.TrimSpace(cols.Eq(2).Text())

		funName := nameMap[e.Name]
		if funName == "" {
			funName = capitalize(e.Name)
		}

		if e.Desc != "" {
			e.Desc = fmt.Sprintf("%s is an event fired when %s", funName, lowercase(e.Desc))
		} else {
			e.Desc = "(no documentation)"
		}
		events[funName] = &e
	})

	var names []string
	for name := range events {
		names = append(names, name)
	}
	sort.Strings(names)

	file, err := os.Create("event.gen.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, `//go:generate go run generate.go

// Package event defines markup to bind DOM events.
//
// Generated from "Event reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/Events, licensed under
// CC-BY-SA 2.5.
package event

import "github.com/gopherjs/vecty"
`)

	for _, name := range names {
		e := events[name]
		if e.Spec == "WebVR API" {
			continue // not stabilized
		}
		fmt.Fprintf(file, `%s
//
// https://developer.mozilla.org%s
func %s(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "%s", Listener: listener}
}
`, descToComments(e.Desc), e.Link[6:], name, e.Name)
	}
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func lowercase(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
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
