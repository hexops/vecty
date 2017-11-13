package vecty

import (
	"fmt"
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

type testCore struct{ Core }

func (testCore) Render() ComponentOrHTML { return Tag("p") }

type testCorePtr struct{ *Core }

func (testCorePtr) Render() ComponentOrHTML { return Tag("p") }

func TestCore(t *testing.T) {
	// Test that a standard *MyComponent with embedded Core works as we expect.
	t.Run("comp_ptr_and_core", func(t *testing.T) {
		v1 := Tag("v1")
		valid := Component(&testCore{})
		valid.Context().prevRender = v1
		if valid.Context().prevRender != v1 {
			t.Fatal("valid.Context().prevRender != v1")
		}
	})

	// Test that a non-pointer MyComponent with embedded Core does not satisfy
	// the Component interface:
	//
	//  testCore does not implement Component (Context method has pointer receiver)
	//
	t.Run("comp_and_core", func(t *testing.T) {
		isComponent := func(x interface{}) bool {
			_, ok := x.(Component)
			return ok
		}
		if isComponent(testCore{}) {
			t.Fatal("expected !isComponent(testCompCore{})")
		}
	})

	// Test what happens when a user accidentally embeds *Core instead of Core in
	// their component.
	t.Run("comp_ptr_and_core_ptr", func(t *testing.T) {
		v1 := Tag("v1")
		invalid := Component(&testCorePtr{})
		got := recoverStr(func() {
			invalid.Context().prevRender = v1
		})
		// TODO(slimsag): This would happen in user-facing code too. We should
		// create a helper for when we access a component's context, which
		// would panic with a more helpful message.
		want := "runtime error: invalid memory address or nil pointer dereference"
		if got != want {
			t.Fatalf("got panic %q want %q", got, want)
		}
	})
	t.Run("comp_and_core_ptr", func(t *testing.T) {
		v1 := Tag("v1")
		invalid := Component(testCorePtr{})
		got := recoverStr(func() {
			invalid.Context().prevRender = v1
		})
		// TODO(slimsag): This would happen in user-facing code too. We should
		// create a helper for when we access a component's context, which
		// would panic with a more helpful message.
		want := "runtime error: invalid memory address or nil pointer dereference"
		if got != want {
			t.Fatalf("got panic %q want %q", got, want)
		}
	})
}

// TODO(slimsag): TestUnmounter; Unmounter.Unmount

func TestHTML_Node(t *testing.T) {
	// Create a non-nil *js.Object. For 'gopherjs test', &js.Object{} == nil
	// because it is special-cased; but for 'go test' js.Global == nil.
	x := js.Global // used for 'gopherjs test'
	if x == nil {
		x = &js.Object{} // used for 'go test'
	}
	h := &HTML{node: wrapObject(x)}
	if h.Node() != x {
		t.Fatal("h.Node() != x")
	}
}

// TestHTML_reconcile_std tests that (*HTML).reconcile against an old HTML instance
// works as expected (i.e. that it updates nodes correctly).
func TestHTML_reconcile_std(t *testing.T) {
	t.Run("text_identical", func(t *testing.T) {
		h := Text("foobar")
		hNode := &mockObject{}
		h.node = hNode
		prev := Text("foobar")
		prevNode := &mockObject{}
		prev.node = prevNode
		h.reconcile(prev)
		if h.node != prevNode {
			t.Fatal("h.node != prevNode")
		}
	})
	t.Run("text_diff", func(t *testing.T) {
		want := "bar"
		h := Text(want)
		hNode := &mockObject{}
		h.node = hNode
		prev := Text("foo")
		setNodeValue := ""
		prevNode := &mockObject{
			set: func(key string, value interface{}) {
				if key != "nodeValue" {
					panic(`key != "nodeValue"`)
				}
				setNodeValue = value.(string)
			},
		}
		prev.node = prevNode
		h.reconcile(prev)
		if h.node != prevNode {
			t.Fatal("h.node != prevNode")
		}
		if setNodeValue != want {
			t.Fatalf("got %q want %q", setNodeValue, want)
		}
	})
	t.Run("properties", func(t *testing.T) {
		cases := []struct {
			name         string
			initHTML     *HTML
			initResult   string
			targetHTML   *HTML
			targetResult string
		}{
			{
				name:         "diff",
				initHTML:     Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Property("a", 3), Property("b", "4foobar"))),
				targetResult: "a:3 b:4foobar",
			},
			{
				name:         "remove",
				initHTML:     Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Property("a", 3))),
				targetResult: "a:3",
			},
			{
				name:         "replaced_elem_diff",
				initHTML:     Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("span", Markup(Property("a", 3), Property("b", "4foobar"))),
				targetResult: "a:3 b:4foobar",
			},
			{
				name:         "replaced_elem_shared",
				initHTML:     Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("span", Markup(Property("a", 1), Property("b", "4foobar"))),
				targetResult: "a:1 b:4foobar",
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				initSet := make(map[string]interface{})
				targetSet := make(map[string]interface{})
				initElem := &mockObject{
					set: func(key string, value interface{}) {
						initSet[key] = value
					},
					delete: func(key string) {
						delete(initSet, key)
					},
				}
				targetElem := &mockObject{
					set: func(key string, value interface{}) {
						targetSet[key] = value
					},
					delete: func(key string) {
						delete(targetSet, key)
					},
				}
				wrapperFunc := func(obj jsObject) func(string, ...interface{}) jsObject {
					return func(name string, args ...interface{}) jsObject {
						if name != "createElement" {
							panic(fmt.Sprintf("expected call to createElement, not %q", name))
						}
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if args[0].(string) != "div" && args[0].(string) != "span" {
							panic(`args[0].(string) != "div|span"`)
						}
						return obj
					}
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": &mockObject{call: wrapperFunc(initElem)},
					},
				}
				tst.initHTML.reconcile(nil)
				got := sortedMapString(initSet)
				if got != tst.initResult {
					t.Fatalf("got %q want %q", got, tst.initResult)
				}
				matchingTags := tst.initHTML.tag == tst.targetHTML.tag
				if !matchingTags {
					global = &mockObject{
						get: map[string]jsObject{
							"document": &mockObject{call: wrapperFunc(targetElem)},
						},
					}
				}
				tst.targetHTML.reconcile(tst.initHTML)
				if matchingTags {
					got = sortedMapString(initSet)
				} else {
					got = sortedMapString(targetSet)
				}
				if got != tst.targetResult {
					t.Fatalf("got %q want %q", got, tst.targetResult)
				}
			})
		}
	})
	t.Run("attributes", func(t *testing.T) {
		cases := []struct {
			name         string
			initHTML     *HTML
			initResult   string
			targetHTML   *HTML
			targetResult string
		}{
			{
				name:         "diff",
				initHTML:     Tag("div", Markup(Attribute("a", 1), Attribute("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Attribute("a", 3), Attribute("b", "4foobar"))),
				targetResult: "a:3 b:4foobar",
			},
			{
				name:         "remove",
				initHTML:     Tag("div", Markup(Attribute("a", 1), Attribute("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Attribute("a", 3))),
				targetResult: "a:3",
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				set := map[string]interface{}{}
				div := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						switch name {
						case "setAttribute":
							if len(args) != 2 {
								panic("setAttribute: len(args) != 2")
							}
							set[args[0].(string)] = args[1]
						case "removeAttribute":
							if len(args) != 1 {
								panic("removeAttribute: len(args) != 1")
							}
							delete(set, args[0].(string))
						default:
							panic(fmt.Sprintf("expected call to [setAttribute, removeAttribute], not %q", name))
						}
						return nil
					},
				}
				document := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						if name != "createElement" {
							panic(fmt.Sprintf("expected call to createElement, not %q", name))
						}
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if args[0].(string) != "div" {
							panic(`args[0].(string) != "div"`)
						}
						return div
					},
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": document,
					},
				}
				tst.initHTML.reconcile(nil)
				got := sortedMapString(set)
				if got != tst.initResult {
					t.Fatalf("got %q want %q", got, tst.initResult)
				}
				tst.targetHTML.reconcile(tst.initHTML)
				got = sortedMapString(set)
				if got != tst.targetResult {
					t.Fatalf("got %q want %q", got, tst.targetResult)
				}
			})
		}
	})
	t.Run("class", func(t *testing.T) {
		cases := []struct {
			name         string
			initHTML     *HTML
			initResult   string
			targetHTML   *HTML
			targetResult string
		}{
			{
				name:         "multi",
				initHTML:     Tag("div", Markup(Class("a"), Class("b"))),
				initResult:   "a:true b:true",
				targetHTML:   Tag("div", Markup(Class("a"), Class("c"))),
				targetResult: "a:true c:true",
			},
			{
				name:         "diff",
				initHTML:     Tag("div", Markup(Class("a", "b"))),
				initResult:   "a:true b:true",
				targetHTML:   Tag("div", Markup(Class("a", "c"))),
				targetResult: "a:true c:true",
			},
			{
				name:         "remove",
				initHTML:     Tag("div", Markup(Class("a", "b"))),
				initResult:   "a:true b:true",
				targetHTML:   Tag("div", Markup(Class("a"))),
				targetResult: "a:true",
			},
			{
				name:         "map",
				initHTML:     Tag("div", Markup(ClassMap{"a": true, "b": true})),
				initResult:   "a:true b:true",
				targetHTML:   Tag("div", Markup(ClassMap{"a": true})),
				targetResult: "a:true",
			},
			{
				name:         "map_toggle",
				initHTML:     Tag("div", Markup(ClassMap{"a": true, "b": true})),
				initResult:   "a:true b:true",
				targetHTML:   Tag("div", Markup(ClassMap{"a": true, "b": false})),
				targetResult: "a:true",
			},
			{
				name:         "combo",
				initHTML:     Tag("div", Markup(ClassMap{"a": true, "b": true}, Class("c"))),
				initResult:   "a:true b:true c:true",
				targetHTML:   Tag("div", Markup(ClassMap{"a": true, "b": false}, Class("d"))),
				targetResult: "a:true d:true",
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				set := map[string]interface{}{}
				classList := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if _, ok := args[0].(string); !ok {
							panic("args[0].(string) is not string")
						}
						switch name {
						case "add":
							set[args[0].(string)] = true
						case "remove":
							delete(set, args[0].(string))
						default:
							panic(fmt.Sprintf("expected call to add|remove, not %q", name))
						}
						return nil
					},
				}
				div := &mockObject{
					get: map[string]jsObject{
						"classList": classList,
					},
				}
				document := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						if name != "createElement" {
							panic(fmt.Sprintf("expected call to createElement, not %q", name))
						}
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if args[0].(string) != "div" {
							panic(`args[0].(string) != "div"`)
						}
						return div
					},
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": document,
					},
				}
				tst.initHTML.reconcile(nil)
				got := sortedMapString(set)
				if got != tst.initResult {
					t.Fatalf("got %q want %q", got, tst.initResult)
				}
				tst.targetHTML.reconcile(tst.initHTML)
				got = sortedMapString(set)
				if got != tst.targetResult {
					t.Fatalf("got %q want %q", got, tst.targetResult)
				}
			})
		}
	})
	t.Run("dataset", func(t *testing.T) {
		cases := []struct {
			name         string
			initHTML     *HTML
			initResult   string
			targetHTML   *HTML
			targetResult string
		}{
			{
				name:         "diff",
				initHTML:     Tag("div", Markup(Data("a", "1"), Data("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Data("a", "3"), Data("b", "4foobar"))),
				targetResult: "a:3 b:4foobar",
			},
			{
				name:         "remove",
				initHTML:     Tag("div", Markup(Data("a", "1"), Data("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Data("a", "3"))),
				targetResult: "a:3",
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				set := map[string]interface{}{}
				dataset := &mockObject{
					set: func(key string, value interface{}) {
						set[key] = value
					},
					delete: func(key string) {
						delete(set, key)
					},
				}
				div := &mockObject{
					get: map[string]jsObject{
						"dataset": dataset,
					},
				}
				document := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						if name != "createElement" {
							panic(fmt.Sprintf("expected call to createElement, not %q", name))
						}
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if args[0].(string) != "div" {
							panic(`args[0].(string) != "div"`)
						}
						return div
					},
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": document,
					},
				}
				tst.initHTML.reconcile(nil)
				got := sortedMapString(set)
				if got != tst.initResult {
					t.Fatalf("got %q want %q", got, tst.initResult)
				}
				tst.targetHTML.reconcile(tst.initHTML)
				got = sortedMapString(set)
				if got != tst.targetResult {
					t.Fatalf("got %q want %q", got, tst.targetResult)
				}
			})
		}
	})
	t.Run("style", func(t *testing.T) {
		cases := []struct {
			name         string
			initHTML     *HTML
			initResult   string
			targetHTML   *HTML
			targetResult string
		}{
			{
				name:         "diff",
				initHTML:     Tag("div", Markup(Style("a", "1"), Style("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Style("a", "3"), Style("b", "4foobar"))),
				targetResult: "a:3 b:4foobar",
			},
			{
				name:         "remove",
				initHTML:     Tag("div", Markup(Style("a", "1"), Style("b", "2foobar"))),
				initResult:   "a:1 b:2foobar",
				targetHTML:   Tag("div", Markup(Style("a", "3"))),
				targetResult: "a:3",
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				set := map[string]interface{}{}
				style := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						switch name {
						case "setProperty":
							if len(args) != 2 {
								panic("setProperty: len(args) != 2")
							}
							set[args[0].(string)] = args[1]
						case "removeProperty":
							if len(args) != 1 {
								panic("removeProperty: len(args) != 1")
							}
							delete(set, args[0].(string))
						default:
							panic(fmt.Sprintf("expected call to [setProperty, removeProperty], not %q", name))
						}
						return nil
					},
				}
				div := &mockObject{
					get: map[string]jsObject{
						"style": style,
					},
				}
				document := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						if name != "createElement" {
							panic(fmt.Sprintf("expected call to createElement, not %q", name))
						}
						if len(args) != 1 {
							panic("len(args) != 1")
						}
						if args[0].(string) != "div" {
							panic(`args[0].(string) != "div"`)
						}
						return div
					},
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": document,
					},
				}
				tst.initHTML.reconcile(nil)
				got := sortedMapString(set)
				if got != tst.initResult {
					t.Fatalf("got %q want %q", got, tst.initResult)
				}
				tst.targetHTML.reconcile(tst.initHTML)
				got = sortedMapString(set)
				if got != tst.targetResult {
					t.Fatalf("got %q want %q", got, tst.targetResult)
				}
			})
		}
	})
	t.Run("event_listener", func(t *testing.T) {
		// TODO(pdf): Mock listener functions for equality testing
		cases := []struct {
			name                 string
			initEventListeners   []Applyer
			targetEventListeners []Applyer
		}{
			{
				name: "diff",
				initEventListeners: []Applyer{
					&EventListener{Name: "click"},
					&EventListener{Name: "keydown"},
				},
				targetEventListeners: []Applyer{
					&EventListener{Name: "click"},
				},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				addedListeners := map[string]func(*js.Object){}
				div := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						switch name {
						case "addEventListener":
							if len(args) != 2 {
								panic("addEventListener: len(args) != 2")
							}
							addedListeners[args[0].(string)] = args[1].(func(*js.Object))
						case "removeEventListener":
							if len(args) != 2 {
								panic("removeEventListener: len(args) != 2")
							}
							delete(addedListeners, args[0].(string))
						default:
							panic(fmt.Sprintf("unexpected call to %q", name))
						}
						return nil
					},
				}
				document := &mockObject{
					call: func(name string, args ...interface{}) jsObject {
						switch name {
						case "createElement":
							if len(args) != 1 {
								panic("len(args) != 1")
							}
							if args[0].(string) != "div" {
								panic(`args[0].(string) != "div"`)
							}
							return div
						default:
							panic(fmt.Sprintf("unexpected call to %q", name))
						}
					},
				}
				global = &mockObject{
					get: map[string]jsObject{
						"document": document,
					},
				}
				prev := Tag("div", Markup(tst.initEventListeners...))
				prev.reconcile(nil)
				for i, m := range tst.initEventListeners {
					listener := m.(*EventListener)
					if listener.wrapper == nil {
						t.Fatalf("listener %d wrapper == nil: %+v", i, listener)
					}
					if _, ok := addedListeners[listener.Name]; !ok {
						t.Fatalf("listener %d for %q not found: %+v", i, listener.Name, listener)
					}
				}
				if len(tst.initEventListeners) != len(addedListeners) {
					t.Fatalf("listener count mismatch: %d != %d", len(tst.initEventListeners), len(addedListeners))
				}
				h := Tag("div", Markup(tst.targetEventListeners...))
				h.reconcile(prev)
				for i, m := range tst.targetEventListeners {
					listener := m.(*EventListener)
					if listener.wrapper == nil {
						t.Fatalf("listener %d wrapper == nil: %+v", i, listener)
					}
					if _, ok := addedListeners[listener.Name]; !ok {
						t.Fatalf("listener %d for %q not found: %+v", i, listener.Name, listener)
					}
				}
				if len(tst.targetEventListeners) != len(addedListeners) {
					t.Fatalf("listener count mismatch: %d != %d", len(tst.targetEventListeners), len(addedListeners))
				}
			})
		}
	})

	// TODO(pdf): test (*HTML).reconcile child mutations, and value/checked properties
	// TODO(pdf): test multi-pass reconcile of persistent component pointer children, ref: https://github.com/gopherjs/vecty/pull/124
}

// TestHTML_reconcile_nil tests that (*HTML).reconcile(nil) works as expected (i.e.
// that it creates nodes correctly).
func TestHTML_reconcile_nil(t *testing.T) {
	t.Run("one_of_tag_or_text", func(t *testing.T) {
		got := recoverStr(func() {
			h := &HTML{text: "hello", tag: "div"}
			h.reconcile(nil)
		})
		want := "vecty: only one of HTML.tag or HTML.text may be set"
		if got != want {
			t.Fatalf("got panic %q want %q", got, want)
		}
	})
	t.Run("unsafe_text", func(t *testing.T) {
		got := recoverStr(func() {
			h := &HTML{text: "hello", innerHTML: "foobar"}
			h.reconcile(nil)
		})
		want := "vecty: only HTML may have UnsafeHTML attribute"
		if got != want {
			t.Fatalf("got panic %q want %q", got, want)
		}
	})
	t.Run("create_element", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__create_element")
		defer ts.done()

		h := Tag("strong")
		h.reconcile(nil)
	})
	t.Run("create_element_ns", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__create_element_ns")
		defer ts.done()

		h := Tag("strong", Markup(Namespace("foobar")))
		h.reconcile(nil)
	})
	t.Run("create_text_node", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__create_text_node")
		defer ts.done()

		h := Text("hello")
		h.reconcile(nil)
	})
	t.Run("inner_html", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__inner_html")
		defer ts.done()

		h := Tag("div", Markup(UnsafeHTML("<p>hello</p>")))
		h.reconcile(nil)
	})
	t.Run("properties", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__properties")
		defer ts.sortedDone(3, 4)

		h := Tag("div", Markup(Property("a", 1), Property("b", "2foobar")))
		h.reconcile(nil)
	})
	t.Run("attributes", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__attributes")
		defer ts.sortedDone(3, 4)

		h := Tag("div", Markup(Attribute("a", 1), Attribute("b", "2foobar")))
		h.reconcile(nil)
	})
	t.Run("dataset", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__dataset")
		defer ts.sortedDone(5, 6)

		h := Tag("div", Markup(Data("a", "1"), Data("b", "2foobar")))
		h.reconcile(nil)
	})
	t.Run("style", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__style")
		defer ts.sortedDone(6, 7)

		h := Tag("div", Markup(Style("a", "1"), Style("b", "2foobar")))
		h.reconcile(nil)
	})
	t.Run("add_event_listener", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__add_event_listener")
		defer ts.done()

		e0 := &EventListener{Name: "click"}
		e1 := &EventListener{Name: "keydown"}
		h := Tag("div", Markup(e0, e1))
		h.reconcile(nil)
		if e0.wrapper == nil {
			t.Fatal("e0.wrapper == nil")
		}
		if e1.wrapper == nil {
			t.Fatal("e1.wrapper == nil")
		}
	})
	t.Run("children", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__children")
		defer ts.done()

		var compRenderCalls int
		compRender := Tag("div")
		comp := &componentFunc{
			id: "foobar",
			render: func() ComponentOrHTML {
				compRenderCalls++
				return compRender
			},
		}
		h := Tag("div", Tag("div", comp))
		h.reconcile(nil)
		if compRenderCalls != 1 {
			t.Fatal("compRenderCalls != 1")
		}
		if comp.Context().prevRenderComponent.(*componentFunc).id != comp.id {
			t.Fatal("comp.Context().prevRenderComponent.(*componentFunc).id != comp.id")
		}
		if comp.Context().prevRender != compRender {
			t.Fatal("comp.Context().prevRender != compRender")
		}
	})
	t.Run("children_render_nil", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_nil__children_render_nil")
		defer ts.done()

		var compRenderCalls int
		comp := &componentFunc{
			id: "foobar",
			render: func() ComponentOrHTML {
				compRenderCalls++
				return nil
			},
		}
		h := Tag("div", Tag("div", comp))
		h.reconcile(nil)
		if compRenderCalls != 1 {
			t.Fatal("compRenderCalls != 1")
		}
		if comp.Context().prevRenderComponent.(*componentFunc).id != comp.id {
			t.Fatal("comp.Context().prevRenderComponent.(*componentFunc).id != comp.id")
		}
		if comp.Context().prevRender == nil {
			t.Fatal("comp.Context().prevRender == nil")
		}
	})
}

func TestTag(t *testing.T) {
	markupCalled := false
	want := "foobar"
	h := Tag(want, Markup(markupFunc(func(h *HTML) {
		markupCalled = true
	})))
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
	h := Text(want, Markup(markupFunc(func(h *HTML) {
		markupCalled = true
	})))
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

// TestRerender_nil tests that Rerender panics when the component argument is
// nil.
func TestRerender_nil(t *testing.T) {
	gotPanic := ""
	func() {
		defer func() {
			r := recover()
			if r != nil {
				gotPanic = fmt.Sprint(r)
			}
		}()
		Rerender(nil)
	}()
	expected := "vecty: Rerender illegally called with a nil Component argument"
	if gotPanic != expected {
		t.Fatalf("got panic %q expected %q", gotPanic, expected)
	}
}

// TestRerender_no_prevRender tests the behavior of Rerender when there is no
// previous render.
func TestRerender_no_prevRender(t *testing.T) {
	ts := testSuite(t, "TestRerender_no_prevRender")
	defer ts.done()

	got := recoverStr(func() {
		Rerender(&componentFunc{
			render: func() ComponentOrHTML {
				panic("expected no Render call")
			},
			skipRender: func(prev Component) bool {
				panic("expected no SkipRender call")
			},
		})
	})
	want := "vecty: Rerender invoked on Component that has never been rendered"
	if got != want {
		t.Fatalf("got panic %q expected %q", got, want)
	}
}

// TestRerender_identical tests the behavior of Rerender when there is a
// previous render which is identical to the new render.
func TestRerender_identical(t *testing.T) {
	ts := testSuite(t, "TestRerender_identical")
	defer ts.done()

	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")

	// Perform the initial render of the component.
	render := Tag("body")
	var renderCalled, skipRenderCalled int
	comp := &componentFunc{
		id: "original",
		render: func() ComponentOrHTML {
			renderCalled++
			return render
		},
	}
	RenderBody(comp)
	if renderCalled != 1 {
		t.Fatal("renderCalled != 1")
	}
	if comp.Context().prevRender != render {
		t.Fatal("comp.Context().prevRender != render")
	}
	if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
		t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
	}

	// Perform a re-render.
	newRender := Tag("body")
	comp.id = "modified"
	comp.render = func() ComponentOrHTML {
		renderCalled++
		return newRender
	}
	comp.skipRender = func(prev Component) bool {
		if comp.id != "modified" {
			panic(`comp.id != "modified"`)
		}
		if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
			panic(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
		}
		if prev.(*componentFunc).id != "original" {
			panic(`prev.(*componentFunc).id != "original"`)
		}
		skipRenderCalled++
		return false
	}
	Rerender(comp)

	// Invoke the render callback.
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

	if renderCalled != 2 {
		t.Fatal("renderCalled != 2")
	}
	if skipRenderCalled != 1 {
		t.Fatal("skipRenderCalled != 1")
	}
	if comp.Context().prevRender != newRender {
		t.Fatal("comp.Context().prevRender != newRender")
	}
	if comp.Context().prevRenderComponent.(*componentFunc).id != "modified" {
		t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "modified"`)
	}
}

// TestRerender_change tests the behavior of Rerender when there is a
// previous render which is different from the new render.
func TestRerender_change(t *testing.T) {
	cases := []struct {
		name      string
		newRender *HTML
	}{
		{
			name:      "new_child",
			newRender: Tag("body", Tag("div")),
		},
		// TODO(slimsag): bug! nil produces <noscript> and we incorrectly try
		// to replace <body> with it! We should panic & warn the user.
		//{
		//	name:      "nil",
		//	newRender: nil,
		//},
	}
	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			ts := testSuite(t, "TestRerender_change__"+tst.name)
			defer ts.done()

			ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
			ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")

			// Perform the initial render of the component.
			render := Tag("body")
			var renderCalled, skipRenderCalled int
			comp := &componentFunc{
				id: "original",
				render: func() ComponentOrHTML {
					renderCalled++
					return render
				},
			}
			RenderBody(comp)
			ts.record("(expect body to be set now)")
			if renderCalled != 1 {
				t.Fatal("renderCalled != 1")
			}
			if comp.Context().prevRender != render {
				t.Fatal("comp.Context().prevRender != render")
			}
			if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
				t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
			}

			// Perform a re-render.
			comp.id = "modified"
			comp.render = func() ComponentOrHTML {
				renderCalled++
				return tst.newRender
			}
			comp.skipRender = func(prev Component) bool {
				if comp.id != "modified" {
					panic(`comp.id != "modified"`)
				}
				if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
					panic(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
				}
				if prev.(*componentFunc).id != "original" {
					panic(`prev.(*componentFunc).id != "original"`)
				}
				skipRenderCalled++
				return false
			}
			Rerender(comp)

			// Invoke the render callback.
			ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
			ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

			if renderCalled != 2 {
				t.Fatal("renderCalled != 2")
			}
			if skipRenderCalled != 1 {
				t.Fatal("skipRenderCalled != 1")
			}
			if comp.Context().prevRender != tst.newRender {
				t.Fatal("comp.Context().prevRender != tst.newRender")
			}
			if comp.Context().prevRenderComponent.(*componentFunc).id != "modified" {
				t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "modified"`)
			}
		})
	}
}

// TestRerender_Nested tests the behavior of Rerender when there is a
// nested Component that is exchanged for *HTML.
func TestRerender_Nested(t *testing.T) {
	cases := []struct {
		name                     string
		initialRender, newRender ComponentOrHTML
	}{
		{
			name:          "html_to_component",
			initialRender: Tag("body"),
			newRender: &componentFunc{
				render: func() ComponentOrHTML {
					return Tag("body", Tag("div"))
				},
				skipRender: func(Component) bool {
					return false
				},
			},
		},
		{
			name: "component_to_html",
			initialRender: &componentFunc{
				render: func() ComponentOrHTML {
					return Tag("body")
				},
				skipRender: func(Component) bool {
					return false
				},
			},
			newRender: Tag("body", Tag("div")),
		},
	}
	for _, tst := range cases {
		t.Run(tst.name, func(t *testing.T) {
			ts := testSuite(t, "TestRerender_Nested__"+tst.name)
			defer ts.done()

			ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
			ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")

			// Perform the initial render of the component.
			var renderCalled, skipRenderCalled int
			comp := &componentFunc{
				id: "original",
				render: func() ComponentOrHTML {
					renderCalled++
					return tst.initialRender
				},
			}
			RenderBody(comp)
			ts.record("(expect body to be set now)")
			if renderCalled != 1 {
				t.Fatal("renderCalled != 1")
			}
			if comp.Context().prevRender != tst.initialRender {
				t.Fatal("comp.Context().prevRender != render")
			}
			if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
				t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
			}

			// Perform a re-render.
			comp.id = "modified"
			comp.render = func() ComponentOrHTML {
				renderCalled++
				return tst.newRender
			}
			comp.skipRender = func(prev Component) bool {
				if comp.id != "modified" {
					panic(`comp.id != "modified"`)
				}
				if comp.Context().prevRenderComponent.(*componentFunc).id != "original" {
					panic(`comp.Context().prevRenderComponent.(*componentFunc).id != "original"`)
				}
				if prev.(*componentFunc).id != "original" {
					panic(`prev.(*componentFunc).id != "original"`)
				}
				skipRenderCalled++
				return false
			}
			Rerender(comp)

			// Invoke the render callback.
			ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
			ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

			if skipRenderCalled != 1 {
				t.Fatal("skipRenderCalled != 1")
			}
			if comp.Context().prevRender != tst.newRender {
				t.Fatal("comp.Context().prevRender != tst.newRender")
			}
			if comp.Context().prevRenderComponent.(*componentFunc).id != "modified" {
				t.Fatal(`comp.Context().prevRenderComponent.(*componentFunc).id != "modified"`)
			}
		})
	}
}

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
		{
			name:      "nil",
			render:    nil,
			wantPanic: "vecty: RenderBody expected Component.Render to return a body tag, found \"noscript\"",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := testSuite(t, "TestRenderBody_ExpectsBody__"+c.name)
			defer ts.done()

			ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)

			var gotPanic string
			func() {
				defer func() {
					r := recover()
					if r != nil {
						gotPanic = fmt.Sprint(r)
					}
				}()
				RenderBody(&componentFunc{
					render: func() ComponentOrHTML {
						return c.render
					},
					skipRender: func(prev Component) bool { return false },
				})
			}()
			if c.wantPanic != gotPanic {
				t.Fatalf("want panic %q got panic %q", c.wantPanic, gotPanic)
			}
		})
	}
}

// TestRenderBody_RenderSkipper_Skip tests that RenderBody panics when the
// component's SkipRender method returns skip == true.
func TestRenderBody_RenderSkipper_Skip(t *testing.T) {
	ts := testSuite(t, "TestRenderBody_RenderSkipper_Skip")
	defer ts.done()

	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)

	comp := &componentFunc{
		render: func() ComponentOrHTML {
			return Tag("body")
		},
		skipRender: func(prev Component) bool {
			return true
		},
	}
	fakePrevRender := *comp
	comp.Context().prevRenderComponent = &fakePrevRender
	got := recoverStr(func() {
		RenderBody(comp)
	})
	want := "vecty: RenderBody Component.SkipRender returned true"
	if got != want {
		t.Fatalf("got panic %q want %q", got, want)
	}
}

// TestRenderBody_Standard_loaded tests that RenderBody properly handles the
// standard case of rendering into the "body" tag when the DOM is in a loaded
// state.
func TestRenderBody_Standard_loaded(t *testing.T) {
	ts := testSuite(t, "TestRenderBody_Standard_loaded")
	defer ts.done()

	ts.strings.mock(`global.Get("document").Get("readyState")`, "loaded")
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)

	RenderBody(&componentFunc{
		render: func() ComponentOrHTML {
			return Tag("body")
		},
	})
}

// TestRenderBody_Standard_loading tests that RenderBody properly handles the
// standard case of rendering into the "body" tag when the DOM is in a loading
// state.
func TestRenderBody_Standard_loading(t *testing.T) {
	ts := testSuite(t, "TestRenderBody_Standard_loading")
	defer ts.done()

	ts.strings.mock(`global.Get("document").Get("readyState")`, "loading")
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)

	RenderBody(&componentFunc{
		render: func() ComponentOrHTML {
			return Tag("body")
		},
	})

	ts.record("(invoking DOMContentLoaded event listener)")
	ts.callbacks[`global.Get("document").Call("addEventListener", "DOMContentLoaded", func())`].(func())()
}

// TestRenderBody_Nested tests that RenderBody properly handles nested
// Components.
func TestRenderBody_Nested(t *testing.T) {
	ts := testSuite(t, "TestRenderBody_Nested")
	defer ts.done()

	ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)

	RenderBody(&componentFunc{
		render: func() ComponentOrHTML {
			return &componentFunc{
				render: func() ComponentOrHTML {
					return &componentFunc{
						render: func() ComponentOrHTML {
							return Tag("body")
						},
					}
				},
			}
		},
	})
}

// TestSetTitle tests that the SetTitle function performs the correct DOM
// operations.
func TestSetTitle(t *testing.T) {
	ts := testSuite(t, "TestSetTitle")
	defer ts.done()

	SetTitle("foobartitle")
}

// TestAddStylesheet tests that the AddStylesheet performs the correct DOM
// operations.
func TestAddStylesheet(t *testing.T) {
	ts := testSuite(t, "TestAddStylesheet")
	defer ts.done()

	AddStylesheet("https://google.com/foobar.css")
}
