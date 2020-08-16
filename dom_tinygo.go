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
