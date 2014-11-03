// Documentation source: 
// "Event reference" (https://developer.mozilla.org/docs/Web/Events) by Mozilla Contributors, licensed under CC-BY-SA 2.5.
package event

import (
  "github.com/neelance/dom"
)

// A transaction has been aborted.
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
func Abort(f func()) *dom.EventAspect {
  return dom.Event("abort", f)
}

// The associated document has started printing or the print preview has been closed.
// https://developer.mozilla.org/docs/Web/Events/afterprint
func Afterprint(f func()) *dom.EventAspect {
  return dom.Event("afterprint", f)
}

// A CSS animation has completed.
// https://developer.mozilla.org/docs/Web/Events/animationend
func Animationend(f func()) *dom.EventAspect {
  return dom.Event("animationend", f)
}

// A CSS animation is repeated.
// https://developer.mozilla.org/docs/Web/Events/animationiteration
func Animationiteration(f func()) *dom.EventAspect {
  return dom.Event("animationiteration", f)
}

// A CSS animation has started.
// https://developer.mozilla.org/docs/Web/Events/animationstart
func Animationstart(f func()) *dom.EventAspect {
  return dom.Event("animationstart", f)
}

// The input buffer of a ScriptProcessorNode is ready to be processed.
// https://developer.mozilla.org/docs/Web/Events/audioprocess
func Audioprocess(f func()) *dom.EventAspect {
  return dom.Event("audioprocess", f)
}

// The associated document is about to be printed or previewed for printing.
// https://developer.mozilla.org/docs/Web/Events/beforeprint
func Beforeprint(f func()) *dom.EventAspect {
  return dom.Event("beforeprint", f)
}

//  
// https://developer.mozilla.org/docs/Web/Events/beforeunload
func Beforeunload(f func()) *dom.EventAspect {
  return dom.Event("beforeunload", f)
}

// A SMIL animation element begins.
// https://developer.mozilla.org/docs/Web/Events/beginEvent
func BeginEvent(f func()) *dom.EventAspect {
  return dom.Event("beginEvent", f)
}

// An open connection to a database is blocking a versionchange transaction on the same database.
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
func Blocked(f func()) *dom.EventAspect {
  return dom.Event("blocked", f)
}

// An element has lost focus (does not bubble).
// https://developer.mozilla.org/docs/Web/Events/blur
func Blur(f func()) *dom.EventAspect {
  return dom.Event("blur", f)
}

// The resources listed in the manifest have been downloaded, and the application is now cached.
// https://developer.mozilla.org/docs/Web/Events/cached
func Cached(f func()) *dom.EventAspect {
  return dom.Event("cached", f)
}

// The user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
// https://developer.mozilla.org/docs/Web/Events/canplay
func Canplay(f func()) *dom.EventAspect {
  return dom.Event("canplay", f)
}

// The user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
func Canplaythrough(f func()) *dom.EventAspect {
  return dom.Event("canplaythrough", f)
}

// An element loses focus and its value changed since gaining focus.
// https://developer.mozilla.org/docs/Web/Events/change
func Change(f func()) *dom.EventAspect {
  return dom.Event("change", f)
}

// The battery begins or stops charging.
// https://developer.mozilla.org/docs/Web/Events/chargingchange
func Chargingchange(f func()) *dom.EventAspect {
  return dom.Event("chargingchange", f)
}

// The chargingTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
func Chargingtimechange(f func()) *dom.EventAspect {
  return dom.Event("chargingtimechange", f)
}

// The user agent is checking for an update, or attempting to download the cache manifest for the first time.
// https://developer.mozilla.org/docs/Web/Events/checking
func Checking(f func()) *dom.EventAspect {
  return dom.Event("checking", f)
}

// A pointing device button has been pressed and released on an element.
// https://developer.mozilla.org/docs/Web/Events/click
func Click(f func()) *dom.EventAspect {
  return dom.Event("click", f)
}

// A WebSocket connection has been closed.
// https://developer.mozilla.org/docs/Web/Reference/Events/close_websocket
func Close(f func()) *dom.EventAspect {
  return dom.Event("close", f)
}

// The compass used to obtain orientation data is in need of calibration.
// https://developer.mozilla.org/docs/Web/Events/compassneedscalibration
func Compassneedscalibration(f func()) *dom.EventAspect {
  return dom.Event("compassneedscalibration", f)
}

// The rendering of an OfflineAudioContext is terminated.
// https://developer.mozilla.org/docs/Web/Events/complete
func Complete(f func()) *dom.EventAspect {
  return dom.Event("complete", f)
}

// The composition of a passage of text has been completed or canceled.
// https://developer.mozilla.org/docs/Web/Events/compositionend
func Compositionend(f func()) *dom.EventAspect {
  return dom.Event("compositionend", f)
}

// The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition).
// https://developer.mozilla.org/docs/Web/Events/compositionstart
func Compositionstart(f func()) *dom.EventAspect {
  return dom.Event("compositionstart", f)
}

// A character is added to a passage of text being composed.
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
func Compositionupdate(f func()) *dom.EventAspect {
  return dom.Event("compositionupdate", f)
}

// The right button of the mouse is clicked (before the context menu is displayed).
// https://developer.mozilla.org/docs/Web/Events/contextmenu
func Contextmenu(f func()) *dom.EventAspect {
  return dom.Event("contextmenu", f)
}

// The text selection has been added to the clipboard.
// https://developer.mozilla.org/docs/Web/Events/copy
func Copy(f func()) *dom.EventAspect {
  return dom.Event("copy", f)
}

// The text selection has been removed from the document and added to the clipboard.
// https://developer.mozilla.org/docs/Web/Events/cut
func Cut(f func()) *dom.EventAspect {
  return dom.Event("cut", f)
}

// The document has finished loading (but not its dependent resources).
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
func DOMContentLoaded(f func()) *dom.EventAspect {
  return dom.Event("DOMContentLoaded", f)
}

// A pointing device button is clicked twice on an element.
// https://developer.mozilla.org/docs/Web/Events/dblclick
func Dblclick(f func()) *dom.EventAspect {
  return dom.Event("dblclick", f)
}

// Fresh data is available from a light sensor.
// https://developer.mozilla.org/docs/Web/Events/devicelight
func Devicelight(f func()) *dom.EventAspect {
  return dom.Event("devicelight", f)
}

// Fresh data is available from a motion sensor.
// https://developer.mozilla.org/docs/Web/Events/devicemotion
func Devicemotion(f func()) *dom.EventAspect {
  return dom.Event("devicemotion", f)
}

// Fresh data is available from an orientation sensor.
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
func Deviceorientation(f func()) *dom.EventAspect {
  return dom.Event("deviceorientation", f)
}

// Fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object).
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
func Deviceproximity(f func()) *dom.EventAspect {
  return dom.Event("deviceproximity", f)
}

// The dischargingTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
func Dischargingtimechange(f func()) *dom.EventAspect {
  return dom.Event("dischargingtimechange", f)
}

// The user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time.
// https://developer.mozilla.org/docs/Web/Events/downloading
func Downloading(f func()) *dom.EventAspect {
  return dom.Event("downloading", f)
}

// An element or text selection is being dragged (every 350ms).
// https://developer.mozilla.org/docs/Web/Events/drag
func Drag(f func()) *dom.EventAspect {
  return dom.Event("drag", f)
}

// A drag operation is being ended (by releasing a mouse button or hitting the escape key).
// https://developer.mozilla.org/docs/Web/Events/dragend
func Dragend(f func()) *dom.EventAspect {
  return dom.Event("dragend", f)
}

// A dragged element or text selection enters a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/dragenter
func Dragenter(f func()) *dom.EventAspect {
  return dom.Event("dragenter", f)
}

// A dragged element or text selection leaves a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/dragleave
func Dragleave(f func()) *dom.EventAspect {
  return dom.Event("dragleave", f)
}

// An element or text selection is being dragged over a valid drop target (every 350ms).
// https://developer.mozilla.org/docs/Web/Events/dragover
func Dragover(f func()) *dom.EventAspect {
  return dom.Event("dragover", f)
}

// The user starts dragging an element or text selection.
// https://developer.mozilla.org/docs/Web/Events/dragstart
func Dragstart(f func()) *dom.EventAspect {
  return dom.Event("dragstart", f)
}

// An element is dropped on a valid drop target.
// https://developer.mozilla.org/docs/Web/Events/drop
func Drop(f func()) *dom.EventAspect {
  return dom.Event("drop", f)
}

// The duration attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/durationchange
func Durationchange(f func()) *dom.EventAspect {
  return dom.Event("durationchange", f)
}

// The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load() method is called to reload it.
// https://developer.mozilla.org/docs/Web/Events/emptied
func Emptied(f func()) *dom.EventAspect {
  return dom.Event("emptied", f)
}

// A SMIL animation element ends.
// https://developer.mozilla.org/docs/Web/Events/endEvent
func EndEvent(f func()) *dom.EventAspect {
  return dom.Event("endEvent", f)
}

//  
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
func Ended(f func()) *dom.EventAspect {
  return dom.Event("ended", f)
}

// A request caused an error and failed.
// https://developer.mozilla.org/docs/Web/Events/error
func Error(f func()) *dom.EventAspect {
  return dom.Event("error", f)
}

// An element has received focus (does not bubble).
// https://developer.mozilla.org/docs/Web/Events/focus
func Focus(f func()) *dom.EventAspect {
  return dom.Event("focus", f)
}

// An element is about to receive focus (bubbles).
// https://developer.mozilla.org/docs/Web/Events/focusin
func Focusin(f func()) *dom.EventAspect {
  return dom.Event("focusin", f)
}

// An element is about to lose focus (bubbles).
// https://developer.mozilla.org/docs/Web/Events/focusout
func Focusout(f func()) *dom.EventAspect {
  return dom.Event("focusout", f)
}

// An element was turned to fullscreen mode or back to normal mode.
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
func Fullscreenchange(f func()) *dom.EventAspect {
  return dom.Event("fullscreenchange", f)
}

// It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied.
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
func Fullscreenerror(f func()) *dom.EventAspect {
  return dom.Event("fullscreenerror", f)
}

// A gamepad has been connected.
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
func Gamepadconnected(f func()) *dom.EventAspect {
  return dom.Event("gamepadconnected", f)
}

// A gamepad has been disconnected.
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
func Gamepaddisconnected(f func()) *dom.EventAspect {
  return dom.Event("gamepaddisconnected", f)
}

// The fragment identifier of the URL has changed (the part of the URL after the #).
// https://developer.mozilla.org/docs/Web/Events/hashchange
func Hashchange(f func()) *dom.EventAspect {
  return dom.Event("hashchange", f)
}

// The value of an element changes or the content of an element with the attribute contenteditable is modified.
// https://developer.mozilla.org/docs/Web/Events/input
func Input(f func()) *dom.EventAspect {
  return dom.Event("input", f)
}

// A submittable element has been checked and doesn't satisfy its constraints.
// https://developer.mozilla.org/docs/Web/Events/invalid
func Invalid(f func()) *dom.EventAspect {
  return dom.Event("invalid", f)
}

// A key is pressed down.
// https://developer.mozilla.org/docs/Web/Events/keydown
func Keydown(f func()) *dom.EventAspect {
  return dom.Event("keydown", f)
}

// A key is pressed down and that key normally produces a character value (use input instead).
// https://developer.mozilla.org/docs/Web/Events/keypress
func Keypress(f func()) *dom.EventAspect {
  return dom.Event("keypress", f)
}

// A key is released.
// https://developer.mozilla.org/docs/Web/Events/keyup
func Keyup(f func()) *dom.EventAspect {
  return dom.Event("keyup", f)
}

//  
// https://developer.mozilla.org/docs/Web/Events/languagechange
func Languagechange(f func()) *dom.EventAspect {
  return dom.Event("languagechange", f)
}

// The level attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/levelchange
func Levelchange(f func()) *dom.EventAspect {
  return dom.Event("levelchange", f)
}

// Progression has been successful.
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
func Load(f func()) *dom.EventAspect {
  return dom.Event("load", f)
}

// The first frame of the media has finished loading.
// https://developer.mozilla.org/docs/Web/Events/loadeddata
func Loadeddata(f func()) *dom.EventAspect {
  return dom.Event("loadeddata", f)
}

// The metadata has been loaded.
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
func Loadedmetadata(f func()) *dom.EventAspect {
  return dom.Event("loadedmetadata", f)
}

// Progress has stopped (after "error", "abort" or "load" have been dispatched).
// https://developer.mozilla.org/docs/Web/Events/loadend
func Loadend(f func()) *dom.EventAspect {
  return dom.Event("loadend", f)
}

// Progress has begun.
// https://developer.mozilla.org/docs/Web/Events/loadstart
func Loadstart(f func()) *dom.EventAspect {
  return dom.Event("loadstart", f)
}

// A message is received through an event source.
// https://developer.mozilla.org/docs/Web/Reference/Events/message_serversentevents
func Message(f func()) *dom.EventAspect {
  return dom.Event("message", f)
}

// A pointing device button (usually a mouse) is pressed on an element.
// https://developer.mozilla.org/docs/Web/Events/mousedown
func Mousedown(f func()) *dom.EventAspect {
  return dom.Event("mousedown", f)
}

// A pointing device is moved onto the element that has the listener attached.
// https://developer.mozilla.org/docs/Web/Events/mouseenter
func Mouseenter(f func()) *dom.EventAspect {
  return dom.Event("mouseenter", f)
}

// A pointing device is moved off the element that has the listener attached.
// https://developer.mozilla.org/docs/Web/Events/mouseleave
func Mouseleave(f func()) *dom.EventAspect {
  return dom.Event("mouseleave", f)
}

// A pointing device is moved over an element.
// https://developer.mozilla.org/docs/Web/Events/mousemove
func Mousemove(f func()) *dom.EventAspect {
  return dom.Event("mousemove", f)
}

// A pointing device is moved off the element that has the listener attached or off one of its children.
// https://developer.mozilla.org/docs/Web/Events/mouseout
func Mouseout(f func()) *dom.EventAspect {
  return dom.Event("mouseout", f)
}

// A pointing device is moved onto the element that has the listener attached or onto one of its children.
// https://developer.mozilla.org/docs/Web/Events/mouseover
func Mouseover(f func()) *dom.EventAspect {
  return dom.Event("mouseover", f)
}

// A pointing device button is released over an element.
// https://developer.mozilla.org/docs/Web/Events/mouseup
func Mouseup(f func()) *dom.EventAspect {
  return dom.Event("mouseup", f)
}

// The manifest hadn't changed.
// https://developer.mozilla.org/docs/Web/Events/noupdate
func Noupdate(f func()) *dom.EventAspect {
  return dom.Event("noupdate", f)
}

// The manifest was found to have become a 404 or 410 page, so the application cache is being deleted.
// https://developer.mozilla.org/docs/Web/Events/obsolete
func Obsolete(f func()) *dom.EventAspect {
  return dom.Event("obsolete", f)
}

// The browser has lost access to the network.
// https://developer.mozilla.org/docs/Web/Events/offline
func Offline(f func()) *dom.EventAspect {
  return dom.Event("offline", f)
}

// The browser has gained access to the network (but particular websites might be unreachable).
// https://developer.mozilla.org/docs/Web/Events/online
func Online(f func()) *dom.EventAspect {
  return dom.Event("online", f)
}

// An event source connection has been established.
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
func Open(f func()) *dom.EventAspect {
  return dom.Event("open", f)
}

// The orientation of the device (portrait/landscape) has changed
// https://developer.mozilla.org/docs/Web/Events/orientationchange
func Orientationchange(f func()) *dom.EventAspect {
  return dom.Event("orientationchange", f)
}

// A session history entry is being traversed from.
// https://developer.mozilla.org/docs/Web/Events/pagehide
func Pagehide(f func()) *dom.EventAspect {
  return dom.Event("pagehide", f)
}

// A session history entry is being traversed to.
// https://developer.mozilla.org/docs/Web/Events/pageshow
func Pageshow(f func()) *dom.EventAspect {
  return dom.Event("pageshow", f)
}

// Data has been transfered from the system clipboard to the document.
// https://developer.mozilla.org/docs/Web/Events/paste
func Paste(f func()) *dom.EventAspect {
  return dom.Event("paste", f)
}

// Playback has been paused.
// https://developer.mozilla.org/docs/Web/Events/pause
func Pause(f func()) *dom.EventAspect {
  return dom.Event("pause", f)
}

// Playback has begun.
// https://developer.mozilla.org/docs/Web/Events/play
func Play(f func()) *dom.EventAspect {
  return dom.Event("play", f)
}

// Playback is ready to start after having been paused or delayed due to lack of data.
// https://developer.mozilla.org/docs/Web/Events/playing
func Playing(f func()) *dom.EventAspect {
  return dom.Event("playing", f)
}

// The pointer was locked or released.
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
func Pointerlockchange(f func()) *dom.EventAspect {
  return dom.Event("pointerlockchange", f)
}

// It was impossible to lock the pointer for technical reasons or because the permission was denied.
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
func Pointerlockerror(f func()) *dom.EventAspect {
  return dom.Event("pointerlockerror", f)
}

// A session history entry is being navigated to (in certain cases).
// https://developer.mozilla.org/docs/Web/Events/popstate
func Popstate(f func()) *dom.EventAspect {
  return dom.Event("popstate", f)
}

// The user agent is downloading resources listed by the manifest.
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
func Progress(f func()) *dom.EventAspect {
  return dom.Event("progress", f)
}

// The playback rate has changed.
// https://developer.mozilla.org/docs/Web/Events/ratechange
func Ratechange(f func()) *dom.EventAspect {
  return dom.Event("ratechange", f)
}

// The readyState attribute of a document has changed.
// https://developer.mozilla.org/docs/Web/Events/readystatechange
func Readystatechange(f func()) *dom.EventAspect {
  return dom.Event("readystatechange", f)
}

// A SMIL animation element is repeated.
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
func RepeatEvent(f func()) *dom.EventAspect {
  return dom.Event("repeatEvent", f)
}

// A form is reset.
// https://developer.mozilla.org/docs/Web/Events/reset
func Reset(f func()) *dom.EventAspect {
  return dom.Event("reset", f)
}

// The document view has been resized.
// https://developer.mozilla.org/docs/Web/Events/resize
func Resize(f func()) *dom.EventAspect {
  return dom.Event("resize", f)
}

// Page loading has been stopped before the SVG was loaded.
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
func SVGAbort(f func()) *dom.EventAspect {
  return dom.Event("SVGAbort", f)
}

// An error has occurred before the SVG was loaded.
// https://developer.mozilla.org/docs/Web/Events/SVGError
func SVGError(f func()) *dom.EventAspect {
  return dom.Event("SVGError", f)
}

// An SVG document has been loaded and parsed.
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
func SVGLoad(f func()) *dom.EventAspect {
  return dom.Event("SVGLoad", f)
}

// An SVG document is being resized.
// https://developer.mozilla.org/docs/Web/Events/SVGResize
func SVGResize(f func()) *dom.EventAspect {
  return dom.Event("SVGResize", f)
}

// An SVG document is being scrolled.
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
func SVGScroll(f func()) *dom.EventAspect {
  return dom.Event("SVGScroll", f)
}

// An SVG document has been removed from a window or frame.
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
func SVGUnload(f func()) *dom.EventAspect {
  return dom.Event("SVGUnload", f)
}

// An SVG document is being zoomed.
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
func SVGZoom(f func()) *dom.EventAspect {
  return dom.Event("SVGZoom", f)
}

// The document view or an element has been scrolled.
// https://developer.mozilla.org/docs/Web/Events/scroll
func Scroll(f func()) *dom.EventAspect {
  return dom.Event("scroll", f)
}

// A seek operation completed.
// https://developer.mozilla.org/docs/Web/Events/seeked
func Seeked(f func()) *dom.EventAspect {
  return dom.Event("seeked", f)
}

// A seek operation began.
// https://developer.mozilla.org/docs/Web/Events/seeking
func Seeking(f func()) *dom.EventAspect {
  return dom.Event("seeking", f)
}

// Some text is being selected.
// https://developer.mozilla.org/docs/Web/Events/select
func Select(f func()) *dom.EventAspect {
  return dom.Event("select", f)
}

// A contextmenu event was fired on/bubbled to an element that has a contextmenu attribute
// https://developer.mozilla.org/docs/Web/Events/show
func Show(f func()) *dom.EventAspect {
  return dom.Event("show", f)
}

// The user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
// https://developer.mozilla.org/docs/Web/Events/stalled
func Stalled(f func()) *dom.EventAspect {
  return dom.Event("stalled", f)
}

// A storage area (localStorage or sessionStorage) has changed.
// https://developer.mozilla.org/docs/Web/Events/storage
func Storage(f func()) *dom.EventAspect {
  return dom.Event("storage", f)
}

// A form is submitted.
// https://developer.mozilla.org/docs/Web/Events/submit
func Submit(f func()) *dom.EventAspect {
  return dom.Event("submit", f)
}

// A request successfully completed.
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
func Success(f func()) *dom.EventAspect {
  return dom.Event("success", f)
}

// Media data loading has been suspended.
// https://developer.mozilla.org/docs/Web/Events/suspend
func Suspend(f func()) *dom.EventAspect {
  return dom.Event("suspend", f)
}

//  
// https://developer.mozilla.org/docs/Web/Events/timeout
func Timeout(f func()) *dom.EventAspect {
  return dom.Event("timeout", f)
}

// The time indicated by the currentTime attribute has been updated.
// https://developer.mozilla.org/docs/Web/Events/timeupdate
func Timeupdate(f func()) *dom.EventAspect {
  return dom.Event("timeupdate", f)
}

// A touch point has been disrupted in an implementation-specific manners (too many touch points for example).
// https://developer.mozilla.org/docs/Web/Events/touchcancel
func Touchcancel(f func()) *dom.EventAspect {
  return dom.Event("touchcancel", f)
}

// A touch point is removed from the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchend
func Touchend(f func()) *dom.EventAspect {
  return dom.Event("touchend", f)
}

// A touch point is moved onto the interactive area of an element.
// https://developer.mozilla.org/docs/Web/Events/touchenter
func Touchenter(f func()) *dom.EventAspect {
  return dom.Event("touchenter", f)
}

// A touch point is moved off the interactive area of an element.
// https://developer.mozilla.org/docs/Web/Events/touchleave
func Touchleave(f func()) *dom.EventAspect {
  return dom.Event("touchleave", f)
}

// A touch point is moved along the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchmove
func Touchmove(f func()) *dom.EventAspect {
  return dom.Event("touchmove", f)
}

// A touch point is placed on the touch surface.
// https://developer.mozilla.org/docs/Web/Events/touchstart
func Touchstart(f func()) *dom.EventAspect {
  return dom.Event("touchstart", f)
}

// A CSS transition has completed.
// https://developer.mozilla.org/docs/Web/Events/transitionend
func Transitionend(f func()) *dom.EventAspect {
  return dom.Event("transitionend", f)
}

// The document or a dependent resource is being unloaded.
// https://developer.mozilla.org/docs/Web/Events/unload
func Unload(f func()) *dom.EventAspect {
  return dom.Event("unload", f)
}

// The resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache.
// https://developer.mozilla.org/docs/Web/Events/updateready
func Updateready(f func()) *dom.EventAspect {
  return dom.Event("updateready", f)
}

// An attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created.
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
func Upgradeneeded(f func()) *dom.EventAspect {
  return dom.Event("upgradeneeded", f)
}

// Fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not).
// https://developer.mozilla.org/docs/Web/Events/userproximity
func Userproximity(f func()) *dom.EventAspect {
  return dom.Event("userproximity", f)
}

// A versionchange transaction completed.
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
func Versionchange(f func()) *dom.EventAspect {
  return dom.Event("versionchange", f)
}

// The content of a tab has become visible or has been hidden.
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
func Visibilitychange(f func()) *dom.EventAspect {
  return dom.Event("visibilitychange", f)
}

// The volume has changed.
// https://developer.mozilla.org/docs/Web/Events/volumechange
func Volumechange(f func()) *dom.EventAspect {
  return dom.Event("volumechange", f)
}

// Playback has stopped because of a temporary lack of data.
// https://developer.mozilla.org/docs/Web/Events/waiting
func Waiting(f func()) *dom.EventAspect {
  return dom.Event("waiting", f)
}

// A wheel button of a pointing device is rotated in any direction.
// https://developer.mozilla.org/docs/Web/Events/wheel
func Wheel(f func()) *dom.EventAspect {
  return dom.Event("wheel", f)
}
