//go:generate go run generate.go

// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.
package event

import (
	"github.com/gopherjs/vecty"
)

// A transaction has been aborted.
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
func Abort(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "abort", Listener: listener}
}

// The associated document has started printing or the print preview has been closed.
// https://developer.mozilla.org/docs/Web/Events/afterprint
func AfterPrint(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "afterprint", Listener: listener}
}

// A CSS animation has completed.
// https://developer.mozilla.org/docs/Web/Events/animationend
func AnimationEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "animationend", Listener: listener}
}

// A CSS animation is repeated.
// https://developer.mozilla.org/docs/Web/Events/animationiteration
func AnimationIteration(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "animationiteration", Listener: listener}
}

// A CSS animation has started.
// https://developer.mozilla.org/docs/Web/Events/animationstart
func AnimationStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "animationstart", Listener: listener}
}

// The user agent has finished capturing audio for speech recognition.
// https://developer.mozilla.org/docs/Web/Events/audioend
func AudioEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "audioend", Listener: listener}
}

// The input buffer of a ScriptProcessorNode is ready to be processed.
// https://developer.mozilla.org/docs/Web/Events/audioprocess
func AudioProcess(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "audioprocess", Listener: listener}
}

// The user agent has started to capture audio for speech recognition.
// https://developer.mozilla.org/docs/Web/Events/audiostart
func AudioStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "audiostart", Listener: listener}
}

// The associated document is about to be printed or previewed for printing.
// https://developer.mozilla.org/docs/Web/Events/beforeprint
func BeforePrint(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "beforeprint", Listener: listener}
}

// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/beforeunload
func BeforeUnload(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "beforeunload", Listener: listener}
}

// A SMIL animation element begins.
// https://developer.mozilla.org/docs/Web/Events/beginEvent
func BeginEvent(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "beginEvent", Listener: listener}
}

// An open connection to a database is blocking a versionchange transaction on the same database.
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
func Blocked(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "blocked", Listener: listener}
}

// An element has lost focus (does not bubble).
// https://developer.mozilla.org/docs/Web/Events/blur
func Blur(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "blur", Listener: listener}
}

// The spoken utterance reaches a word or sentence boundary
// https://developer.mozilla.org/docs/Web/Events/boundary
func Boundary(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "boundary", Listener: listener}
}

// The resources listed in the manifest have been downloaded, and the application is now cached.
// https://developer.mozilla.org/docs/Web/Events/cached
func Cached(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "cached", Listener: listener}
}

// The user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
// https://developer.mozilla.org/docs/Web/Events/canplay
func CanPlay(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "canplay", Listener: listener}
}

// The user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
func CanPlayThrough(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "canplaythrough", Listener: listener}
}

// The change event is fired for <input>, <select>, and <textarea> elements when a change to the element's value is committed by the user.
// https://developer.mozilla.org/docs/Web/Events/change
func Change(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "change", Listener: listener}
}

// The battery begins or stops charging.
// https://developer.mozilla.org/docs/Web/Events/chargingchange
func ChargingChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "chargingchange", Listener: listener}
}

// The chargingTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
func ChargingTimeChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "chargingtimechange", Listener: listener}
}

// The user agent is checking for an update, or attempting to download the cache manifest for the first time.
// https://developer.mozilla.org/docs/Web/Events/checking
func Checking(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "checking", Listener: listener}
}

// A pointing device button has been pressed and released on an element.
// https://developer.mozilla.org/docs/Web/Events/click
func Click(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "click", Listener: listener}
}

// A WebSocket connection has been closed.
// https://developer.mozilla.org/docs/Web/Reference/Events/close_websocket
func Close(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "close", Listener: listener}
}

// The rendering of an OfflineAudioContext is terminated.
// https://developer.mozilla.org/docs/Web/Events/complete
func Complete(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "complete", Listener: listener}
}

// The composition of a passage of text has been completed or canceled.
// https://developer.mozilla.org/docs/Web/Events/compositionend
func CompositionEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "compositionend", Listener: listener}
}

// The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
// https://developer.mozilla.org/docs/Web/Events/compositionstart
func CompositionStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "compositionstart", Listener: listener}
}

// A character is added to a passage of text being composed.
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
func CompositionUpdate(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "compositionupdate", Listener: listener}
}

// The right button of the mouse is clicked (before the context menu is displayed).
// https://developer.mozilla.org/docs/Web/Events/contextmenu
func ContextMenu(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "contextmenu", Listener: listener}
}

// The text selection has been added to the clipboard.
// https://developer.mozilla.org/docs/Web/Events/copy
func Copy(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "copy", Listener: listener}
}

// The text selection has been removed from the document and added to the clipboard.
// https://developer.mozilla.org/docs/Web/Events/cut
func Cut(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "cut", Listener: listener}
}

// The document has finished loading (but not its dependent resources).
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
func DOMContentLoaded(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "DOMContentLoaded", Listener: listener}
}

// Fresh data is available from a light sensor.
// https://developer.mozilla.org/docs/Web/Events/devicelight
func DeviceLight(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "devicelight", Listener: listener}
}

// Fresh data is available from a motion sensor.
// https://developer.mozilla.org/docs/Web/Events/devicemotion
func DeviceMotion(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "devicemotion", Listener: listener}
}

// Fresh data is available from an orientation sensor.
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
func DeviceOrientation(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "deviceorientation", Listener: listener}
}

// Fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object).
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
func DeviceProximity(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "deviceproximity", Listener: listener}
}

// The dischargingTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
func DischargingTimeChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dischargingtimechange", Listener: listener}
}

// A pointing device button is clicked twice on an element.
// https://developer.mozilla.org/docs/Web/Events/dblclick
func DoubleClick(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dblclick", Listener: listener}
}

// The user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time.
// https://developer.mozilla.org/docs/Web/Events/downloading
func Downloading(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "downloading", Listener: listener}
}

// An element or text selection is being dragged (every 350ms).
// https://developer.mozilla.org/docs/Web/Events/drag
func Drag(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "drag", Listener: listener}
}

// A drag operation is being ended (by releasing a mouse button or hitting the escape key).
// https://developer.mozilla.org/docs/Web/Events/dragend
func DragEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dragend", Listener: listener}
}

// A dragged element or text selection enters a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/dragenter
func DragEnter(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dragenter", Listener: listener}
}

// A dragged element or text selection leaves a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/dragleave
func DragLeave(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dragleave", Listener: listener}
}

// An element or text selection is being dragged over a valid drop target (every 350ms).
// https://developer.mozilla.org/docs/Web/Events/dragover
func DragOver(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dragover", Listener: listener}
}

// The user starts dragging an element or text selection.
// https://developer.mozilla.org/docs/Web/Events/dragstart
func DragStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "dragstart", Listener: listener}
}

// An element is dropped on a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/drop
func Drop(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "drop", Listener: listener}
}

// The duration attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/durationchange
func DurationChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "durationchange", Listener: listener}
}

// The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
// https://developer.mozilla.org/docs/Web/Events/emptied
func Emptied(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "emptied", Listener: listener}
}

// The utterance has finished being spoken.
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
func End(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "end", Listener: listener}
}

// A SMIL animation element ends.
// https://developer.mozilla.org/docs/Web/Events/endEvent
func EndEvent(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "endEvent", Listener: listener}
}

// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
func Ended(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "ended", Listener: listener}
}

// An error occurs that prevents the utterance from being succesfully spoken.
// https://developer.mozilla.org/docs/Web/Events/error_(SpeechSynthesisError)
func Error(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "error", Listener: listener}
}

// An element has received focus (does not bubble).
// https://developer.mozilla.org/docs/Web/Events/focus
func Focus(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "focus", Listener: listener}
}

// An element is about to receive focus (bubbles).
// https://developer.mozilla.org/docs/Web/Events/focusin
func FocusIn(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "focusin", Listener: listener}
}

// An element is about to lose focus (bubbles).
// https://developer.mozilla.org/docs/Web/Events/focusout
func FocusOut(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "focusout", Listener: listener}
}

// An element was turned to fullscreen mode or back to normal mode.
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
func FullScreenChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "fullscreenchange", Listener: listener}
}

// It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
func FullScreenError(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "fullscreenerror", Listener: listener}
}

// A gamepad has been connected.
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
func GamepadConnected(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "gamepadconnected", Listener: listener}
}

// A gamepad has been disconnected.
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
func GamepadDisconnected(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "gamepaddisconnected", Listener: listener}
}

// Element receives pointer capture.
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
func GotPointerCapture(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "gotpointercapture", Listener: listener}
}

// The fragment identifier of the URL has changed (the part of the URL after the #).
// https://developer.mozilla.org/docs/Web/Events/hashchange
func HashChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "hashchange", Listener: listener}
}

// The value of an element changes or the content of an element with the attribute contenteditable is modified.
// https://developer.mozilla.org/docs/Web/Events/input
func Input(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "input", Listener: listener}
}

// A submittable element has been checked and doesn't satisfy its constraints.
// https://developer.mozilla.org/docs/Web/Events/invalid
func Invalid(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "invalid", Listener: listener}
}

// A key is pressed down.
// https://developer.mozilla.org/docs/Web/Events/keydown
func KeyDown(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "keydown", Listener: listener}
}

// A key is pressed down and that key normally produces a character value (use input instead).
// https://developer.mozilla.org/docs/Web/Events/keypress
func KeyPress(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "keypress", Listener: listener}
}

// A key is released.
// https://developer.mozilla.org/docs/Web/Events/keyup
func KeyUp(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "keyup", Listener: listener}
}

// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/languagechange
func LanguageChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "languagechange", Listener: listener}
}

// The level attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/levelchange
func LevelChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "levelchange", Listener: listener}
}

// Progression has been successful.
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
func Load(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "load", Listener: listener}
}

// Progress has stopped (after "error", "abort" or "load" have been dispatched).
// https://developer.mozilla.org/docs/Web/Events/loadend
func LoadEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "loadend", Listener: listener}
}

// Progress has begun.
// https://developer.mozilla.org/docs/Web/Events/loadstart
func LoadStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "loadstart", Listener: listener}
}

// The first frame of the media has finished loading.
// https://developer.mozilla.org/docs/Web/Events/loadeddata
func LoadedData(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "loadeddata", Listener: listener}
}

// The metadata has been loaded.
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
func LoadedMetadata(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "loadedmetadata", Listener: listener}
}

// Element lost pointer capture.
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
func LostPointerCapture(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "lostpointercapture", Listener: listener}
}

// The spoken utterance reaches a named SSML "mark" tag.
// https://developer.mozilla.org/docs/Web/Events/mark
func Mark(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mark", Listener: listener}
}

// A message is received from a service worker, or a message is received in a service worker from another context.
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
func Message(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "message", Listener: listener}
}

// A pointing device button (usually a mouse) is pressed on an element.
// https://developer.mozilla.org/docs/Web/Events/mousedown
func MouseDown(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mousedown", Listener: listener}
}

// A pointing device is moved onto the element that has the listener attached.
// https://developer.mozilla.org/docs/Web/Events/mouseenter
func MouseEnter(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mouseenter", Listener: listener}
}

// A pointing device is moved off the element that has the listener attached.
// https://developer.mozilla.org/docs/Web/Events/mouseleave
func MouseLeave(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mouseleave", Listener: listener}
}

// A pointing device is moved over an element.
// https://developer.mozilla.org/docs/Web/Events/mousemove
func MouseMove(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mousemove", Listener: listener}
}

// A pointing device is moved off the element that has the listener attached or off one of its children.
// https://developer.mozilla.org/docs/Web/Events/mouseout
func MouseOut(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mouseout", Listener: listener}
}

// A pointing device is moved onto the element that has the listener attached or onto one of its children.
// https://developer.mozilla.org/docs/Web/Events/mouseover
func MouseOver(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mouseover", Listener: listener}
}

// A pointing device button is released over an element.
// https://developer.mozilla.org/docs/Web/Events/mouseup
func MouseUp(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "mouseup", Listener: listener}
}

// The speech recognition service returns a final result with no significant recognition.
// https://developer.mozilla.org/docs/Web/Events/nomatch
func NoMatch(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "nomatch", Listener: listener}
}

// The manifest hadn't changed.
// https://developer.mozilla.org/docs/Web/Events/noupdate
func NoUpdate(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "noupdate", Listener: listener}
}

// A system notification spawned by ServiceWorkerRegistration.showNotification() has been clicked.
// https://developer.mozilla.org/docs/Web/Events/notificationclick
func NotificationClick(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "notificationclick", Listener: listener}
}

// The manifest was found to have become a 404 or 410 page, so the application cache is being deleted.
// https://developer.mozilla.org/docs/Web/Events/obsolete
func Obsolete(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "obsolete", Listener: listener}
}

// The browser has lost access to the network.
// https://developer.mozilla.org/docs/Web/Events/offline
func Offline(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "offline", Listener: listener}
}

// The browser has gained access to the network (but particular websites might be unreachable).
// https://developer.mozilla.org/docs/Web/Events/online
func Online(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "online", Listener: listener}
}

// An event source connection has been established.
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
func Open(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "open", Listener: listener}
}

// The orientation of the device (portrait/landscape) has changed
// https://developer.mozilla.org/docs/Web/Events/orientationchange
func OrientationChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "orientationchange", Listener: listener}
}

// A session history entry is being traversed from.
// https://developer.mozilla.org/docs/Web/Events/pagehide
func PageHide(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pagehide", Listener: listener}
}

// A session history entry is being traversed to.
// https://developer.mozilla.org/docs/Web/Events/pageshow
func PageShow(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pageshow", Listener: listener}
}

// Data has been transfered from the system clipboard to the document.
// https://developer.mozilla.org/docs/Web/Events/paste
func Paste(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "paste", Listener: listener}
}

// The utterance is paused part way through.
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
func Pause(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pause", Listener: listener}
}

// Playback has begun.
// https://developer.mozilla.org/docs/Web/Events/play
func Play(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "play", Listener: listener}
}

// Playback is ready to start after having been paused or delayed due to lack of data.
// https://developer.mozilla.org/docs/Web/Events/playing
func Playing(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "playing", Listener: listener}
}

// The pointer is unlikely to produce any more events.
// https://developer.mozilla.org/docs/Web/Events/pointercancel
func PointerCancel(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointercancel", Listener: listener}
}

// The pointer enters the active buttons state.
// https://developer.mozilla.org/docs/Web/Events/pointerdown
func PointerDown(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerdown", Listener: listener}
}

// Pointing device is moved inside the hit-testing boundary.
// https://developer.mozilla.org/docs/Web/Events/pointerenter
func PointerEnter(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerenter", Listener: listener}
}

// Pointing device is moved out of the hit-testing boundary.
// https://developer.mozilla.org/docs/Web/Events/pointerleave
func PointerLeave(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerleave", Listener: listener}
}

// The pointer was locked or released.
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
func PointerLockChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerlockchange", Listener: listener}
}

// It was impossible to lock the pointer for technical reasons or because the permission was denied.
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
func PointerLockError(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerlockerror", Listener: listener}
}

// The pointer changed coordinates.
// https://developer.mozilla.org/docs/Web/Events/pointermove
func PointerMove(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointermove", Listener: listener}
}

// The pointing device moved out of hit-testing boundary or leaves detectable hover range.
// https://developer.mozilla.org/docs/Web/Events/pointerout
func PointerOut(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerout", Listener: listener}
}

// The pointing device is moved into the hit-testing boundary.
// https://developer.mozilla.org/docs/Web/Events/pointerover
func PointerOver(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerover", Listener: listener}
}

// The pointer leaves the active buttons state.
// https://developer.mozilla.org/docs/Web/Events/pointerup
func PointerUp(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pointerup", Listener: listener}
}

// A session history entry is being navigated to (in certain cases).
// https://developer.mozilla.org/docs/Web/Events/popstate
func PopState(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "popstate", Listener: listener}
}

// The user agent is downloading resources listed by the manifest.
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
func Progress(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "progress", Listener: listener}
}

// A Service Worker has received a push message.
// https://developer.mozilla.org/docs/Web/Events/push
func Push(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "push", Listener: listener}
}

// A PushSubscription has expired.
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
func PushSubscriptionChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "pushsubscriptionchange", Listener: listener}
}

// The playback rate has changed.
// https://developer.mozilla.org/docs/Web/Events/ratechange
func RateChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "ratechange", Listener: listener}
}

// The readyState attribute of a document has changed.
// https://developer.mozilla.org/docs/Web/Events/readystatechange
func ReadyStateChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "readystatechange", Listener: listener}
}

// A SMIL animation element is repeated.
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
func RepeatEvent(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "repeatEvent", Listener: listener}
}

// A form is reset.
// https://developer.mozilla.org/docs/Web/Events/reset
func Reset(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "reset", Listener: listener}
}

// The document view has been resized.
// https://developer.mozilla.org/docs/Web/Events/resize
func Resize(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "resize", Listener: listener}
}

// The browser's resource timing buffer is full.
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
func ResourceTimingBufferFull(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "resourcetimingbufferfull", Listener: listener}
}

// The speech recognition service returns a result — a word or phrase has been positively recognized and this has been communicated back to the app.
// https://developer.mozilla.org/docs/Web/Events/result
func Result(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "result", Listener: listener}
}

// A paused utterance is resumed.
// https://developer.mozilla.org/docs/Web/Events/resume
func Resume(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "resume", Listener: listener}
}

// Page loading has been stopped before the SVG was loaded.
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
func SVGAbort(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGAbort", Listener: listener}
}

// An error has occurred before the SVG was loaded.
// https://developer.mozilla.org/docs/Web/Events/SVGError
func SVGError(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGError", Listener: listener}
}

// An SVG document has been loaded and parsed.
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
func SVGLoad(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGLoad", Listener: listener}
}

// An SVG document is being resized.
// https://developer.mozilla.org/docs/Web/Events/SVGResize
func SVGResize(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGResize", Listener: listener}
}

// An SVG document is being scrolled.
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
func SVGScroll(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGScroll", Listener: listener}
}

// An SVG document has been removed from a window or frame.
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
func SVGUnload(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGUnload", Listener: listener}
}

// An SVG document is being zoomed.
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
func SVGZoom(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "SVGZoom", Listener: listener}
}

// The document view or an element has been scrolled.
// https://developer.mozilla.org/docs/Web/Events/scroll
func Scroll(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "scroll", Listener: listener}
}

// A seek operation completed.
// https://developer.mozilla.org/docs/Web/Events/seeked
func Seeked(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "seeked", Listener: listener}
}

// A seek operation began.
// https://developer.mozilla.org/docs/Web/Events/seeking
func Seeking(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "seeking", Listener: listener}
}

// Some text is being selected.
// https://developer.mozilla.org/docs/Web/Events/select
func Select(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "select", Listener: listener}
}

// A selection just started.
// https://developer.mozilla.org/docs/Web/Events/selectstart
func SelectStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "selectstart", Listener: listener}
}

// The selection in the document has been changed.
// https://developer.mozilla.org/docs/Web/Events/selectionchange
func SelectionChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "selectionchange", Listener: listener}
}

// A contextmenu event was fired on/bubbled to an element that has a contextmenu attribute
// https://developer.mozilla.org/docs/Web/Events/show
func Show(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "show", Listener: listener}
}

// Any sound — recognisable speech or not — has stopped being detected.
// https://developer.mozilla.org/docs/Web/Events/soundend
func SoundEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "soundend", Listener: listener}
}

// Any sound — recognisable speech or not — has been detected.
// https://developer.mozilla.org/docs/Web/Events/soundstart
func SoundStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "soundstart", Listener: listener}
}

// Speech recognised by the speech recognition service has stopped being detected.
// https://developer.mozilla.org/docs/Web/Events/speechend
func SpeechEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "speechend", Listener: listener}
}

// Sound that is recognised by the speech recognition service as speech has been detected.
// https://developer.mozilla.org/docs/Web/Events/speechstart
func SpeechStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "speechstart", Listener: listener}
}

// The user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
// https://developer.mozilla.org/docs/Web/Events/stalled
func Stalled(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "stalled", Listener: listener}
}

// The utterance has begun to be spoken.
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
func Start(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "start", Listener: listener}
}

// A storage area (localStorage or sessionStorage) has changed.
// https://developer.mozilla.org/docs/Web/Events/storage
func Storage(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "storage", Listener: listener}
}

// A form is submitted.
// https://developer.mozilla.org/docs/Web/Events/submit
func Submit(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "submit", Listener: listener}
}

// A request successfully completed.
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
func Success(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "success", Listener: listener}
}

// Media data loading has been suspended.
// https://developer.mozilla.org/docs/Web/Events/suspend
func Suspend(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "suspend", Listener: listener}
}

// The time indicated by the currentTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/timeupdate
func TimeUpdate(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "timeupdate", Listener: listener}
}

// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/timeout
func Timeout(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "timeout", Listener: listener}
}

// A touch point has been disrupted in an implementation-specific manners (too many touch points for example).
// https://developer.mozilla.org/docs/Web/Events/touchcancel
func TouchCancel(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchcancel", Listener: listener}
}

// A touch point is removed from the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchend
func TouchEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchend", Listener: listener}
}

// A touch point is moved onto the interactive area of an element.
// https://developer.mozilla.org/docs/Web/Events/touchenter
func TouchEnter(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchenter", Listener: listener}
}

// A touch point is moved off the interactive area of an element.
// https://developer.mozilla.org/docs/Web/Events/touchleave
func TouchLeave(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchleave", Listener: listener}
}

// A touch point is moved along the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchmove
func TouchMove(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchmove", Listener: listener}
}

// A touch point is placed on the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchstart
func TouchStart(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "touchstart", Listener: listener}
}

// A CSS transition has completed.
// https://developer.mozilla.org/docs/Web/Events/transitionend
func TransitionEnd(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "transitionend", Listener: listener}
}

// The document or a dependent resource is being unloaded.
// https://developer.mozilla.org/docs/Web/Events/unload
func Unload(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "unload", Listener: listener}
}

// The resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache.
// https://developer.mozilla.org/docs/Web/Events/updateready
func UpdateReady(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "updateready", Listener: listener}
}

// An attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created.
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
func UpgradeNeeded(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "upgradeneeded", Listener: listener}
}

// Fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not).
// https://developer.mozilla.org/docs/Web/Events/userproximity
func UserProximity(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "userproximity", Listener: listener}
}

// A versionchange transaction completed.
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
func VersionChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "versionchange", Listener: listener}
}

// The content of a tab has become visible or has been hidden.
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
func VisibilityChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "visibilitychange", Listener: listener}
}

// The list of SpeechSynthesisVoice objects that would be returned by the SpeechSynthesis.getVoices() method has changed (when the voiceschanged event fires.)
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
func VoicesChanged(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "voiceschanged", Listener: listener}
}

// The volume has changed.
// https://developer.mozilla.org/docs/Web/Events/volumechange
func VolumeChange(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "volumechange", Listener: listener}
}

// Playback has stopped because of a temporary lack of data.
// https://developer.mozilla.org/docs/Web/Events/waiting
func Waiting(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "waiting", Listener: listener}
}

// A wheel button of a pointing device is rotated in any direction.
// https://developer.mozilla.org/docs/Web/Events/wheel
func Wheel(listener func(*vecty.Event)) *vecty.EventListener {
	return &vecty.EventListener{Name: "wheel", Listener: listener}
}
