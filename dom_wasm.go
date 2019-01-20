// +build wasm
// +build go1.12

package vecty

import "syscall/js"

func runGoForever() {
	select {}
}

// Event represents a DOM event.
//
// When compiling under WebAssembly, these values will have syscall/js.Value types.
//
// When compiling under GopherJS, these values will have *gopherjs/js.Object types.
type Event struct {
	js.Value
	Target js.Value
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
//
// When compiling under WebAssembly, the return type will be syscall/js.Value.
//
// When compiling under GopherJS, the return type will be *gopherjs/js.Object.
func (h *HTML) Node() js.Value {
	if h.node == nil {
		panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
	}
	return h.node.(wrappedObject).j
}

var (
	global    = wrapObject(js.Global())
	undefined = wrappedObject{js.Undefined()}
)

func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
	return jsFuncImpl(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wrappedArgs := make([]jsObject, len(args))
		for i, arg := range args {
			wrappedArgs[i] = wrapObject(arg)
		}
		result := fn(wrapObject(this), wrappedArgs)
		if wrapped, ok := result.(wrappedObject); ok {
			return wrapped.j
		}
		return result
	}))
}

type jsFuncImpl js.Func

func (j jsFuncImpl) Object() jsObject { return wrapObject(j.Value) }
func (j jsFuncImpl) Release()         { j.Release() }

func wrapObject(j js.Value) jsObject {
	if j == js.Null() {
		return nil
	}
	if j == js.Undefined() {
		return undefined
	}
	return wrappedObject{j}
}

type wrappedObject struct {
	j js.Value
}

func (w wrappedObject) Set(key string, value interface{}) {
	if v, ok := value.(wrappedObject); ok {
		value = v.j
	}
	w.j.Set(key, value)
}

func (w wrappedObject) Get(key string) jsObject {
	return wrapObject(w.j.Get(key))
}

func (w wrappedObject) Delete(key string) {
	w.j.Call("delete", key)
}

func (w wrappedObject) Call(name string, args ...interface{}) jsObject {
	for i, arg := range args {
		if v, ok := arg.(wrappedObject); ok {
			args[i] = v.j
		}
	}
	return wrapObject(w.j.Call(name, args...))
}

func (w wrappedObject) String() string {
	return w.j.String()
}
func (w wrappedObject) Bool() bool {
	return w.j.Bool()
}

func (w wrappedObject) Int() int {
	return w.j.Int()
}

func (w wrappedObject) Float() float64 {
	return w.j.Float()
}
