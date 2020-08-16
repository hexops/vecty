// +build !js

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

func init() {
	htmlNodeImpl = func(h *HTML) SyscallJSValue {
		if h.node == nil {
			panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
		}
		return h.node.(wrappedObject).j
	}
	funcOfImpl = func(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
		return &jsFuncImpl{
			goFunc: fn,
		}
	}
	valueOfImpl = func(v interface{}) jsObject {
		ts := global().(*objectRecorder).ts
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
}
