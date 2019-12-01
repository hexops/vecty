// +build go1.12,wasm,js js

package vecty

import "syscall/js"

// Event represents a DOM event.
type Event struct {
	js.Value
	Target js.Value
}

// Node returns the underlying JavaScript Element or TextNode.
//
// It panics if it is called before the DOM node has been attached, i.e. before
// the associated component's Mounter interface would be invoked.
func (h *HTML) Node() js.Value {
	if h.node == nil {
		panic("vecty: cannot call (*HTML).Node() before DOM node creation / component mount")
	}
	return h.node.(wrappedObject).j
}

// RenderIntoNode renders the given component into the existing HTML element by
// replacing it.
//
// If the Component's Render method does not return an element of the same type,
// an error of type ElementMismatchError is returned.
func RenderIntoNode(node js.Value, c Component) error {
	return renderIntoNode("RenderIntoNode", wrapObject(node), c)
}

func toLower(s string) string {
	// We must call the prototype method here to workaround a limitation of
	// syscall/js in both Go and GopherJS where we cannot call the
	// `toLowerCase` string method. See https://github.com/golang/go/issues/35917
	return js.Global().Get("String").Get("prototype").Get("toLowerCase").Call("call", js.ValueOf(s)).String()
}

var (
	global    = wrapObject(js.Global())
	undefined = wrappedObject{js.Undefined()}
)

func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc {
	return jsFuncImpl{
		f: js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			wrappedArgs := make([]jsObject, len(args))
			for i, arg := range args {
				wrappedArgs[i] = wrapObject(arg)
			}
			return unwrap(fn(wrapObject(this), wrappedArgs))
		}),
		goFunc: fn,
	}
}

type jsFuncImpl struct {
	f      js.Func
	goFunc func(this jsObject, args []jsObject) interface{}
}

func (j jsFuncImpl) String() string {
	// fmt.Sprint(j) would produce the actual implementation of the function in
	// JS code which differs across WASM/GopherJS/TinyGo so we instead just
	// return an opaque string for testing purposes.
	return "func"
}
func (j jsFuncImpl) Release() { j.f.Release() }

func valueOf(v interface{}) jsObject {
	return wrapObject(js.ValueOf(v))
}

func wrapObject(j js.Value) jsObject {
	if j == js.Null() {
		return nil
	}
	if j == js.Undefined() {
		return undefined
	}
	return wrappedObject{j: j}
}

func unwrap(value interface{}) interface{} {
	if v, ok := value.(wrappedObject); ok {
		return v.j
	}
	if v, ok := value.(jsFuncImpl); ok {
		return v.f
	}
	return value
}

type wrappedObject struct {
	j js.Value
}

func (w wrappedObject) Set(key string, value interface{}) {
	w.j.Set(key, unwrap(value))
}

func (w wrappedObject) Get(key string) jsObject {
	return wrapObject(w.j.Get(key))
}

func (w wrappedObject) Delete(key string) {
	w.j.Call("delete", key)
}

func (w wrappedObject) Call(name string, args ...interface{}) jsObject {
	for i, arg := range args {
		args[i] = unwrap(arg)
	}
	return wrapObject(w.j.Call(name, args...))
}

func (w wrappedObject) String() string {
	return w.j.String()
}

func (w wrappedObject) Truthy() bool {
	return w.j.Truthy()
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
