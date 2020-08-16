// +build !js

package vecty

import "strings"

// Stubs for building Vecty under a native GOOS and GOARCH, so that Vecty
// type-checks, lints, auto-completes, and serves documentation under godoc.org
// as with any other normal Go package that is not under GOOS=js and
// GOARCH=wasm.

// SyscallJSValue is an actual syscall/js.Value type under WebAssembly compilation.
//
// It is declared here just for purposes of testing Vecty under native
// 'go test', linting, and serving documentation under godoc.org.
type SyscallJSValue jsObject

// Event represents a DOM event.
type Event struct {
	Value  SyscallJSValue
	Target SyscallJSValue
}

// Node returns the underlying JavaScript Element or TextNode.
//
// It panics if it is called before the DOM node has been attached, i.e. before
// the associated component's Mounter interface would be invoked.
func (h *HTML) Node() SyscallJSValue {
	return htmlNodeImpl(h)
}

// RenderIntoNode renders the given component into the existing HTML element by
// replacing it.
//
// If the Component's Render method does not return an element of the same type,
// an error of type ElementMismatchError is returned.
func RenderIntoNode(node SyscallJSValue, c Component) error {
	return renderIntoNode("RenderIntoNode", node, c)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

var globalValue jsObject

func global() jsObject {
	return globalValue
}

func undefined() wrappedObject {
	return wrappedObject{j: &jsObjectImpl{}}
}

func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
	return funcOfImpl(fn)
}

type jsFuncImpl struct {
	goFunc func(this jsObject, args []jsObject) interface{}
}

func (j *jsFuncImpl) String() string { return "func" }
func (j *jsFuncImpl) Release()       {}

func valueOf(v interface{}) jsObject { return valueOfImpl(v) }

type wrappedObject struct {
	jsObject
	j jsObject
}

type jsObjectImpl struct {
	jsObject
}

func (e *jsObjectImpl) Equal(other jsObject) bool {
	return e == other.(*jsObjectImpl)
}

var (
	htmlNodeImpl = func(h *HTML) SyscallJSValue {
		panic("not implemented on this architecture in non-testing environment")
	}
	funcOfImpl = func(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
		panic("not implemented on this architecture in non-testing environment")
	}
	valueOfImpl = func(v interface{}) jsObject {
		panic("not implemented on this architecture in non-testing environment")
	}
)
