package vecty

func replaceNode(newNode, oldNode jsObject) {
	if newNode == oldNode {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
