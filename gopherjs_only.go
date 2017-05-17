// +build !js

package vecty

var isTest bool

func init() {
	if isTest {
		return
	}
	panic("vecty: only GopherJS compiler is supported")
}
