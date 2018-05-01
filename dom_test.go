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
		ts := testSuite(t, "TestHTML_reconcile_std__text_identical")
		defer ts.done()

		init := Text("foobar")
		init.reconcile(nil)

		target := Text("foobar")
		target.reconcile(init)
	})
	t.Run("text_diff", func(t *testing.T) {
		ts := testSuite(t, "TestHTML_reconcile_std__text_diff")
		defer ts.done()

		init := Text("bar")
		init.reconcile(nil)

		target := Text("foo")
		target.reconcile(init)
	})
	t.Run("properties", func(t *testing.T) {
		cases := []struct {
			name        string
			initHTML    *HTML
			targetHTML  *HTML
			sortedLines [][2]int
		}{
			{
				name:        "diff",
				initHTML:    Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Property("a", 3), Property("b", "4foobar"))),
				sortedLines: [][2]int{{3, 4}, {12, 13}},
			},
			{
				name:        "remove",
				initHTML:    Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Property("a", 3))),
				sortedLines: [][2]int{{3, 4}},
			},
			{
				name:        "replaced_elem_diff",
				initHTML:    Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				targetHTML:  Tag("span", Markup(Property("a", 3), Property("b", "4foobar"))),
				sortedLines: [][2]int{{3, 4}, {11, 12}},
			},
			{
				name:        "replaced_elem_shared",
				initHTML:    Tag("div", Markup(Property("a", 1), Property("b", "2foobar"))),
				targetHTML:  Tag("span", Markup(Property("a", 1), Property("b", "4foobar"))),
				sortedLines: [][2]int{{3, 4}, {11, 12}},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				ts := testSuite(t, "TestHTML_reconcile_std__properties__"+tst.name)
				defer ts.multiSortedDone(tst.sortedLines...)

				tst.initHTML.reconcile(nil)
				ts.record("(first reconcile done)")
				tst.targetHTML.reconcile(tst.initHTML)
			})
		}
	})
	t.Run("attributes", func(t *testing.T) {
		cases := []struct {
			name        string
			initHTML    *HTML
			targetHTML  *HTML
			sortedLines [][2]int
		}{
			{
				name:        "diff",
				initHTML:    Tag("div", Markup(Attribute("a", 1), Attribute("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Attribute("a", 3), Attribute("b", "4foobar"))),
				sortedLines: [][2]int{{3, 4}, {12, 13}},
			},
			{
				name:        "remove",
				initHTML:    Tag("div", Markup(Attribute("a", 1), Attribute("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Attribute("a", 3))),
				sortedLines: [][2]int{{3, 4}},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				ts := testSuite(t, "TestHTML_reconcile_std__attributes__"+tst.name)
				defer ts.multiSortedDone(tst.sortedLines...)

				tst.initHTML.reconcile(nil)
				ts.record("(first reconcile done)")
				tst.targetHTML.reconcile(tst.initHTML)
			})
		}
	})
	t.Run("class", func(t *testing.T) {
		cases := []struct {
			name        string
			initHTML    *HTML
			targetHTML  *HTML
			sortedLines [][2]int
		}{
			{
				name:        "multi",
				initHTML:    Tag("div", Markup(Class("a"), Class("b"))),
				targetHTML:  Tag("div", Markup(Class("a"), Class("c"))),
				sortedLines: [][2]int{{4, 5}},
			},
			{
				name:        "diff",
				initHTML:    Tag("div", Markup(Class("a", "b"))),
				targetHTML:  Tag("div", Markup(Class("a", "c"))),
				sortedLines: [][2]int{{4, 5}},
			},
			{
				name:        "remove",
				initHTML:    Tag("div", Markup(Class("a", "b"))),
				targetHTML:  Tag("div", Markup(Class("a"))),
				sortedLines: [][2]int{{4, 5}},
			},
			{
				name:        "map",
				initHTML:    Tag("div", Markup(ClassMap{"a": true, "b": true})),
				targetHTML:  Tag("div", Markup(ClassMap{"a": true})),
				sortedLines: [][2]int{{4, 5}},
			},
			{
				name:        "map_toggle",
				initHTML:    Tag("div", Markup(ClassMap{"a": true, "b": true})),
				targetHTML:  Tag("div", Markup(ClassMap{"a": true, "b": false})),
				sortedLines: [][2]int{{4, 5}},
			},
			{
				name:        "combo",
				initHTML:    Tag("div", Markup(ClassMap{"a": true, "b": true}, Class("c"))),
				targetHTML:  Tag("div", Markup(ClassMap{"a": true, "b": false}, Class("d"))),
				sortedLines: [][2]int{{4, 6}, {11, 12}},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				ts := testSuite(t, "TestHTML_reconcile_std__class__"+tst.name)
				defer ts.multiSortedDone(tst.sortedLines...)

				tst.initHTML.reconcile(nil)
				ts.record("(first reconcile done)")
				tst.targetHTML.reconcile(tst.initHTML)
			})
		}
	})
	t.Run("dataset", func(t *testing.T) {
		cases := []struct {
			name        string
			initHTML    *HTML
			targetHTML  *HTML
			sortedLines [][2]int
		}{
			{
				name:        "diff",
				initHTML:    Tag("div", Markup(Data("a", "1"), Data("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Data("a", "3"), Data("b", "4foobar"))),
				sortedLines: [][2]int{{5, 6}, {14, 15}},
			},
			{
				name:        "remove",
				initHTML:    Tag("div", Markup(Data("a", "1"), Data("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Data("a", "3"))),
				sortedLines: [][2]int{{5, 6}},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				ts := testSuite(t, "TestHTML_reconcile_std__dataset__"+tst.name)
				defer ts.multiSortedDone(tst.sortedLines...)

				tst.initHTML.reconcile(nil)
				ts.record("(first reconcile done)")
				tst.targetHTML.reconcile(tst.initHTML)
			})
		}
	})
	t.Run("style", func(t *testing.T) {
		cases := []struct {
			name        string
			initHTML    *HTML
			targetHTML  *HTML
			sortedLines [][2]int
		}{
			{
				name:        "diff",
				initHTML:    Tag("div", Markup(Style("a", "1"), Style("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Style("a", "3"), Style("b", "4foobar"))),
				sortedLines: [][2]int{{6, 7}, {15, 16}},
			},
			{
				name:        "remove",
				initHTML:    Tag("div", Markup(Style("a", "1"), Style("b", "2foobar"))),
				targetHTML:  Tag("div", Markup(Style("a", "3"))),
				sortedLines: [][2]int{{6, 7}},
			},
		}
		for _, tst := range cases {
			t.Run(tst.name, func(t *testing.T) {
				ts := testSuite(t, "TestHTML_reconcile_std__style__"+tst.name)
				defer ts.multiSortedDone(tst.sortedLines...)

				tst.initHTML.reconcile(nil)
				ts.record("(first reconcile done)")
				tst.targetHTML.reconcile(tst.initHTML)
			})
		}
	})
	t.Run("event_listener", func(t *testing.T) {
		// TODO(pdf): Mock listener functions for equality testing
		ts := testSuite(t, "TestHTML_reconcile_std__event_listener_diff")
		defer ts.done()

		initEventListeners := []Applyer{
			&EventListener{Name: "click"},
			&EventListener{Name: "keydown"},
		}
		prev := Tag("div", Markup(initEventListeners...))
		prev.reconcile(nil)
		ts.record("(expected two added event listeners above)")
		for i, m := range initEventListeners {
			listener := m.(*EventListener)
			if listener.wrapper == nil {
				t.Fatalf("listener %d wrapper == nil: %+v", i, listener)
			}
		}

		targetEventListeners := []Applyer{
			&EventListener{Name: "click"},
		}
		h := Tag("div", Markup(targetEventListeners...))
		h.reconcile(prev)
		ts.record("(expected two removed, one added event listeners above)")
		for i, m := range targetEventListeners {
			listener := m.(*EventListener)
			if listener.wrapper == nil {
				t.Fatalf("listener %d wrapper == nil: %+v", i, listener)
			}
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
		want := "vecty: internal error (only one of HTML.tag or HTML.text may be set)"
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

type persistentComponentBody struct {
	Core
}

func (c *persistentComponentBody) Render() ComponentOrHTML {
	return Tag(
		"body",
		&persistentComponent{},
	)
}

var lastRenderedComponent Component
var renderCount int

type persistentComponent struct {
	Core
}

func (c *persistentComponent) Render() ComponentOrHTML {
	if lastRenderedComponent == nil {
		lastRenderedComponent = c
	} else if lastRenderedComponent != c {
		panic("unexpected last rendered component")
	}
	renderCount++
	return Tag("div")
}

// TestRerender_persistent tests the behavior of rendering persistent
// components.
func TestRerender_persistent(t *testing.T) {
	ts := testSuite(t, "TestRerender_persistent")
	defer ts.done()

	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")

	lastRenderedComponent = nil
	renderCount = 0

	comp := &persistentComponentBody{}
	// Perform the initial render of the component.
	RenderBody(comp)

	if renderCount != 1 {
		t.Fatal("renderCount != 1")
	}

	// Perform a re-render.
	Rerender(comp)

	// Invoke the render callback.
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

	if renderCount != 2 {
		t.Fatal("renderCount != 2")
	}

	// Perform a re-render.
	Rerender(comp)

	// Invoke the render callback.
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

	if renderCount != 3 {
		t.Fatal("renderCount != 3")
	}
}

type persistentComponentBody2 struct {
	Core
}

func (c *persistentComponentBody2) Render() ComponentOrHTML {
	return Tag(
		"body",
		&persistentWrapperComponent{},
	)
}

type persistentWrapperComponent struct {
	Core
}

func (c *persistentWrapperComponent) Render() ComponentOrHTML {
	return &persistentComponent{}
}

// TestRerender_persistent_direct tests the behavior of rendering persistent
// components that are directly returned by Render().
func TestRerender_persistent_direct(t *testing.T) {
	ts := testSuite(t, "TestRerender_persistent_direct")
	defer ts.done()

	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.strings.mock(`global.Get("document").Get("readyState")`, "complete")

	lastRenderedComponent = nil
	renderCount = 0

	comp := &persistentComponentBody2{}
	// Perform the initial render of the component.
	RenderBody(comp)

	if renderCount != 1 {
		t.Fatal("renderCount != 1")
	}

	// Perform a re-render.
	Rerender(comp)

	// Invoke the render callback.
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

	if renderCount != 2 {
		t.Fatal("renderCount != 2")
	}

	// Perform a re-render.
	Rerender(comp)

	// Invoke the render callback.
	ts.ints.mock(`global.Call("requestAnimationFrame", func(float64))`, 0)
	ts.callbacks[`global.Call("requestAnimationFrame", func(float64))`].(func(float64))(0)

	if renderCount != 3 {
		t.Fatal("renderCount != 3")
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
	want := "vecty: RenderBody Component.SkipRender illegally returned true"
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
