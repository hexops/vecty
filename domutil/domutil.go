package domutil

import "github.com/gopherjs/gopherjs/js"

func ReplaceNode(newNode, oldNode *js.Object) {
	if newNode == oldNode {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
