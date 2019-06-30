// +build go1.12,!wasm,!js

package vecty

import (
	"fmt"
	"os/exec"
	"reflect"
)

func commandOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	out, _ := cmd.CombinedOutput()
	return string(out), nil
}

func runGoForever() {
	select {}
}

// SyscallJSValue is an actual syscall/js.Value type under WebAssembly and
// GopherJS compilation.
//
// It is declared here just for purposes of testing Vecty under native
// 'go test', linting, and serving documentation under godoc.org.
type SyscallJSValue jsObject

// Event represents a DOM event.
type Event struct {
	Value  SyscallJSValue
	Target SyscallJSValue
}

func newEvent(object, target jsObject) *Event {
	return &Event{
		Value:  object.(wrappedObject).j,
		Target: target.(wrappedObject).j,
	}
}

// Node returns the underlying JavaScript Element or TextNode.
//
// It panics if it is called before the DOM node has been attached, i.e. before
// the associated component's Mounter interface would be invoked.
func (h *HTML) Node() SyscallJSValue {
	if h.node == nil {
		panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
	}
	return h.node.(wrappedObject).j
}

var (
	global    jsObject
	undefined = wrappedObject{&objectRecorder{}, &objectRecorder{}}
)

func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
	return jsFuncImpl{
		goFunc: fn,
	}
}

type jsFuncImpl struct {
	goFunc func(this jsObject, args []jsObject) interface{}
}

func (j jsFuncImpl) String() string { return "func" }
func (j jsFuncImpl) Release()       {}

func valueOf(v interface{}) jsObject {
	ts := global.(*objectRecorder).ts
	name := fmt.Sprintf("valueOf(%v)", v)
	r := &objectRecorder{ts: ts, name: name}
	switch reflect.ValueOf(v).Kind() {
	case reflect.String:
		ts.strings.mock(name, v)
	case reflect.Bool:
		ts.bools.mock(name, v)
	case reflect.Float32, reflect.Float64:
		ts.floats.mock(name, v)
	case reflect.Int:
		ts.ints.mock(name, v)
	default:
	}
	return r
}

type wrappedObject struct {
	*objectRecorder
	j *objectRecorder
}
