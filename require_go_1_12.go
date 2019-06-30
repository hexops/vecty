// +build !go1.12

package vecty

// Typechecking will pass except one error:
//
// 	# github.com/gopherjs/vecty
// 	../../require_go_1_12.go:7:2: undefined: VECTY_REQUIRES_GO_1_12_PLUS
//

func init() {
	VECTY_REQUIRES_GO_1_12_PLUS = "Vecty requires Go 1.12+"
}
