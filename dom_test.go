package vecty

import (
	"fmt"
	"testing"
)

var _ = func() bool {
	isTest = true
	return true
}()

// TODO(slimsag): TestCore; Core.Context
// TODO(slimsag): TestComponent; Component.Render; Component.Context
// TODO(slimsag): TestUnmounter; Unmounter.Unmount
// TODO(slimsag): TestComponentOrHTML
// TODO(slimsag): TestRestorer; Restorer.Restore
// TODO(slimsag): TestHTML; HTML.Restore

func TestTag(t *testing.T) {
	markupCalled := false
	want := "foobar"
	h := Tag(want, markupFunc(func(h *HTML) {
		markupCalled = true
	}))
	if !markupCalled {
		t.Fatal("expected markup to be applied")
	}
	if h.tag != want {
		t.Fatalf("got tag %q want tag %q", h.text, want)
	}
	if h.text != "" {
		t.Fatal("expected no text")
	}
}

func TestText(t *testing.T) {
	markupCalled := false
	want := "Hello world!"
	h := Text(want, markupFunc(func(h *HTML) {
		markupCalled = true
	}))
	if !markupCalled {
		t.Fatal("expected markup to be applied")
	}
	if h.text != want {
		t.Fatalf("got text %q want text %q", h.text, want)
	}
	if h.tag != "" {
		t.Fatal("expected no tag")
	}
}

// TODO(slimsag): TestRerender

// TestRenderBody_ExpectsBody tests that RenderBody panics when something other
// than a "body" tag is rendered by the component.
func TestRenderBody_ExpectsBody(t *testing.T) {
	cases := []struct {
		name      string
		render    *HTML
		wantPanic string
	}{
		{
			name:      "text",
			render:    Text("Hello world!"),
			wantPanic: "vecty: RenderBody expected Component.Render to return a body tag, found \"\"", // TODO(slimsag): error message bug
		},
		{
			name:      "div",
			render:    Tag("div"),
			wantPanic: "vecty: RenderBody expected Component.Render to return a body tag, found \"div\"",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotPanic string
			func() {
				defer func() {
					r := recover()
					if r != nil {
						gotPanic = fmt.Sprint(r)
					}
				}()
				RenderBody(&componentFunc{render: func() *HTML {
					return c.render
				}})
			}()
			if c.wantPanic != gotPanic {
				t.Fatalf("want panic %q got panic %q", c.wantPanic, gotPanic)
			}
		})
	}
}

// TestRenderBody_Standard_loaded tests that RenderBody properly handles the
// standard case of rendering into the "body" tag when the DOM is in a loaded
// state.
func TestRenderBody_Standard_loaded(t *testing.T) {
	body := &mockObject{}
	bodySet := false
	document := &mockObject{
		call: func(name string, args ...interface{}) jsObject {
			if name != "createElement" {
				panic(fmt.Sprintf("expected call to createElement, not %q", name))
			}
			if len(args) != 1 {
				panic("len(args) != 1")
			}
			if args[0].(string) != "body" {
				panic(`args[0].(string) != "body"`)
			}
			return body
		},
		get: map[string]jsObject{
			"readyState": &mockObject{stringValue: "complete"},
		},
		set: func(key string, value interface{}) {
			if key != "body" {
				panic(fmt.Sprintf(`expected document.set "body", not %q`, key))
			}
			if value != body {
				panic(fmt.Sprintf(`expected document.set body value, not %T %+v`, value, value))
			}
			bodySet = true
		},
	}
	global = &mockObject{
		get: map[string]jsObject{
			"document": document,
		},
	}
	restoreCalled := false
	RenderBody(&componentFunc{
		render: func() *HTML {
			return Tag("body")
		},
		restore: func(prev Component) (skip bool) {
			if prev != nil {
				t.Fatal("prev != nil")
			}
			restoreCalled = true
			return false // TODO(slimsag): Make (and test) that skip == true here panics!
		},
	})
	if !restoreCalled {
		t.Fatal("expected Restore to be called")
	}
	if !bodySet {
		t.Fatalf("expected document.body to be set")
	}
}

// TestRenderBody_Standard_loading tests that RenderBody properly handles the
// standard case of rendering into the "body" tag when the DOM is in a loading
// state.
func TestRenderBody_Standard_loading(t *testing.T) {
	body := &mockObject{}
	bodySet := false
	var domLoadedEventListener func()
	document := &mockObject{
		call: func(name string, args ...interface{}) jsObject {
			switch name {
			case "createElement":
				if len(args) != 1 {
					panic("len(args) != 1")
				}
				if args[0].(string) != "body" {
					panic(`args[0].(string) != "body"`)
				}
				return body
			case "addEventListener":
				if len(args) != 2 {
					panic("len(args) != 2")
				}
				if args[0].(string) != "DOMContentLoaded" {
					panic(`args[0].(string) != "DOMContentLoaded"`)
				}
				domLoadedEventListener = args[1].(func())
				return nil
			default:
				panic(fmt.Sprintf("unexpected call to %q", name))
			}
		},
		get: map[string]jsObject{
			"readyState": &mockObject{stringValue: "loading"},
		},
		set: func(key string, value interface{}) {
			if key != "body" {
				panic(fmt.Sprintf(`expected document.set "body", not %q`, key))
			}
			if value != body {
				panic(fmt.Sprintf(`expected document.set body value, not %T %+v`, value, value))
			}
			bodySet = true
		},
	}
	global = &mockObject{
		get: map[string]jsObject{
			"document": document,
		},
	}
	restoreCalled := false
	RenderBody(&componentFunc{
		render: func() *HTML {
			return Tag("body")
		},
		restore: func(prev Component) (skip bool) {
			if prev != nil {
				t.Fatal("prev != nil")
			}
			restoreCalled = true
			return false // TODO(slimsag): Make (and test) that skip == true here panics!
		},
	})
	if !restoreCalled {
		t.Fatal("expected Restore to be called")
	}
	if domLoadedEventListener == nil {
		t.Fatalf("domLoadedEventListener == nil")
	}
	if bodySet {
		t.Fatalf("expected document.body to NOT be set")
	}
	domLoadedEventListener()
	if !bodySet {
		t.Fatalf("expected document.body to be set")
	}
}

func TestSetTitle(t *testing.T) {
	titleSet := ""
	document := &mockObject{
		set: func(key string, value interface{}) {
			if key != "title" {
				panic(fmt.Sprintf(`expected document.set "title", not %q`, key))
			}
			titleSet = value.(string)
		},
	}
	global = &mockObject{
		get: map[string]jsObject{
			"document": document,
		},
	}
	want := "foobar"
	SetTitle(want)
	if titleSet != want {
		t.Fatalf("titleSet is %q, want %q", titleSet, want)
	}
}

func TestAddStylesheet(t *testing.T) {
	linkSet := map[string]interface{}{}
	link := &mockObject{
		set: func(key string, value interface{}) {
			linkSet[key] = value
		},
	}
	appendedToHead := false
	head := &mockObject{
		call: func(name string, args ...interface{}) jsObject {
			switch name {
			case "appendChild":
				if len(args) != 1 {
					panic("len(args) != 1")
				}
				if args[0] != link {
					panic(`args[0] != link`)
				}
				appendedToHead = true
				return nil
			default:
				panic(fmt.Sprintf("unexpected call to %q", name))
			}
		},
	}
	document := &mockObject{
		call: func(name string, args ...interface{}) jsObject {
			switch name {
			case "createElement":
				if len(args) != 1 {
					panic("len(args) != 1")
				}
				if args[0].(string) != "link" {
					panic(`args[0].(string) != "link"`)
				}
				return link
			default:
				panic(fmt.Sprintf("unexpected call to %q", name))
			}
		},
		set: func(key string, value interface{}) {
			if key != "title" {
				panic(fmt.Sprintf(`expected document.set "title", not %q`, key))
			}
		},
		get: map[string]jsObject{
			"head": head,
		},
	}
	global = &mockObject{
		get: map[string]jsObject{
			"document": document,
		},
	}
	url := "https://google.com/foobar.css"
	AddStylesheet(url)
	if !appendedToHead {
		t.Fatal("expected link to be appended to document.head")
	}
	if linkSet["rel"] != "stylesheet" {
		t.Fatal(`linkSet["rel"] != "stylesheet"`)
	}
	if linkSet["href"] != url {
		t.Fatal(`linkSet["href"] != url`)
	}
}

type componentFunc struct {
	Core
	render  func() *HTML
	restore func(prev Component) (skip bool)
}

func (c *componentFunc) Render() *HTML { return c.render() }
func (c *componentFunc) Restore(prev Component) (skip bool) {
	if c.restore != nil {
		return c.restore(prev)
	}
	return
}

type mockObject struct {
	set         func(key string, value interface{})
	get         map[string]jsObject
	call        func(name string, args ...interface{}) jsObject
	stringValue string
	boolValue   bool
}

func (w *mockObject) Set(key string, value interface{})              { w.set(key, value) }
func (w *mockObject) Get(key string) jsObject                        { return w.get[key] }
func (w *mockObject) Call(name string, args ...interface{}) jsObject { return w.call(name, args...) }
func (w *mockObject) String() string                                 { return w.stringValue }
func (w *mockObject) Bool() bool                                     { return w.boolValue }
