package vecty

func init() {
	if isTest {
		return
	}
	// BUG: TinyGo does not allow calling global() inside of init like this.
	//
	//if global() == nil {
	//	panic("vecty: only WebAssembly, TinyGo, and testing compilation is supported")
	//}
	//if global().Get("document").IsUndefined() {
	//	panic("vecty: only running inside a browser is supported")
	//}
}

// BUG: TinyGo does not allow iterating over null maps like these:
//
// 	panic: runtime error: nil pointer dereference
//
// 	main.wasm:1 Uncaught (in promise) RuntimeError: unreachable
// 		at runtime.runtimePanic (http://localhost:8081/main.wasm:wasm-function[61]:0x4472)
// 		at runtime.nilPanic (http://localhost:8081/main.wasm:wasm-function[46]:0x3277)
// 		at runtime.hashmapNext (http://localhost:8081/main.wasm:wasm-function[83]:0x5881)
// 		at (*github.com/gopherjs/vecty.HTML).reconcileProperties (http://localhost:8081/main.wasm:wasm-function[198]:0x117f4)
// 		at (*github.com/gopherjs/vecty.HTML).reconcile (http://localhost:8081/main.wasm:wasm-function[175]:0xda6c)
// 		at github.com/gopherjs/vecty.renderComponent (http://localhost:8081/main.wasm:wasm-function[176]:0xf00d)
// 		at github.com/gopherjs/vecty.renderIntoNode (http://localhost:8081/main.wasm:wasm-function[153]:0xbbe3)
// 		at github.com/gopherjs/vecty.RenderBody (http://localhost:8081/main.wasm:wasm-function[147]:0xb4a8)
// 		at github.com/gopherjs/vecty/example/hellovecty.main (http://localhost:8081/main.wasm:wasm-function[107]:0x6e36)
// 		at runtime.run$1 (http://localhost:8081/main.wasm:wasm-function[55]:0x40c6)
//
func (h *HTML) tinyGoCannotIterateNilMaps() {
	if h.properties == nil {
		h.properties = map[string]interface{}{}
	}
	if h.attributes == nil {
		h.attributes = map[string]interface{}{}
	}
	if h.classes == nil {
		h.classes = map[string]struct{}{}
	}
	if h.dataset == nil {
		h.dataset = map[string]string{}
	}
	if h.styles == nil {
		h.styles = map[string]string{}
	}
}
