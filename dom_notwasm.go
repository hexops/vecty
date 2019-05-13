// +build !wasm

package vecty

import "github.com/gopherjs/gopherjs/js"

func runGoForever() {
	return // GopherJS does not need this, its runtime does not 'exit'.
}

// Event represents a DOM event.
//
// When compiling under WebAssembly, these values will have syscall/js.Value types.
//
// When compiling under GopherJS, these values will have *gopherjs/js.Object types.
type Event struct {
	*js.Object
	Target *js.Object
}

func newEvent(object, target jsObject) *Event {
	return &Event{
		Object: object.(wrappedObject).j,
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
func (h *HTML) Node() *js.Object {
	if h.node == nil {
		panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
	}
	return h.node.(wrappedObject).j
}

var (
	global    = wrapObject(js.Global)
	undefined = wrappedObject{js.Undefined}
)

func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
	return jsFuncImpl(wrappedObject{js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		wrappedArgs := make([]jsObject, len(args))
		for i, arg := range args {
			wrappedArgs[i] = wrapObject(arg)
		}
		result := fn(wrapObject(this), wrappedArgs)
		if wrapped, ok := result.(wrappedObject); ok {
			return wrapped.j
		}
		return result
	})})
}

type jsFuncImpl wrappedObject

func (j jsFuncImpl) Object() jsObject { return wrappedObject(j) }
func (j jsFuncImpl) Release()         {}

func wrapObject(j *js.Object) jsObject {
	if j == nil {
		return nil
	}
	if j == js.Undefined {
		return undefined
	}
	return wrappedObject{j}
}

type wrappedObject struct {
	j *js.Object
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
	w.j.Delete(key)
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
