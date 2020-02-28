package vecty

func replaceNode(newNode, oldNode jsObject) {
	if newNode.Equal(oldNode) {
		return
	}
	oldNode.Get("parentNode").Call("replaceChild", newNode, oldNode)
}
