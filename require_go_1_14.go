// +build !go1.14

package vecty

// Typechecking will pass except one error:
//
// 	# github.com/hexops/vecty
// 	../../require_go_1_14.go:7:2: undefined: VECTY_REQUIRES_GO_1_14_PLUS
//

func init() {
	VECTY_REQUIRES_GO_1_14_PLUS = "Vecty requires Go 1.14+"
}
