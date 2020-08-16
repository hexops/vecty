// +build !tinygo

package vecty

func init() {
	if isTest {
		return
	}
	if global() == nil {
		panic("vecty: only WebAssembly, TinyGo, and testing compilation is supported")
	}
	if global().Get("document").IsUndefined() {
		panic("vecty: only running inside a browser is supported")
	}
}

func (h *HTML) tinyGoCannotIterateNilMaps() {}

func tinyGoAssertCopier(c Component) {}
