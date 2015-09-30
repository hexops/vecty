package domutil

import "github.com/gopherjs/gopherjs/js"

func RemoveNode(node *js.Object) {
	node.Get("parentNode").Call("removeChild", node)
}

func ReplaceNode(newNode, oldNode *js.Object) {
	if newNode == oldNode {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
