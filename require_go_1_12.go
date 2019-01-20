// +build wasm
// +build !go1.12

package vecty

func init() {
	VECTY_WEB_ASSEMBLY_REQUIRES_GO_1_12_PLUS = "Vecty WebAssembly requires Go 1.12+"
}

// Stubs to ensure typechecking passes except one error:
//
// 	# github.com/gopherjs/vecty
// 	../../require_go_1_12.go:7:2: undefined: VECTY_WEB_ASSEMBLY_REQUIRES_GO_1_12_PLUS
//

var (
	global    jsObject
	undefined jsObject
)

type Event struct{}

func newEvent(object, target jsObject) *Event                           { panic("") }
func funcOf(fn func(this jsObject, args []jsObject) interface{}) jsFunc { panic("") }
func runGoForever()                                                     {}
