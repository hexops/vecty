package vecty

import "testing"

// TODO(slimsag): tests for other Markup

func TestNamespace(t *testing.T) {
	want := "b"
	h := Tag("a", Markup(Namespace(want)))
	if h.namespace != want {
		t.Fatalf("got namespace %q want %q", h.namespace, want)
	}
}
