package vecty

func removeNode(node jsObject) {
	node.Get("parentNode").Call("removeChild", node)
}

func replaceNode(newNode, oldNode jsObject) {
	if newNode == oldNode {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
