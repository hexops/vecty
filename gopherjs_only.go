// +build !js

package vecty

func init() {
	if isTest {
		return
	}
	panic("vecty: only GopherJS compiler is supported")
}
