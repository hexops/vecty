package vecty

import (
	"fmt"
)

var (
	errMissingParent = fmt.Errorf("Missing parent node")
	errEmptyElement  = fmt.Errorf("Empty element or node")
)

func removeNode(node jsObject) error {
	parent := node.Get("parentNode")
	// TODO: parent == nil || parent == js.Undefined
	if parent == nil {
		return errMissingParent
	}
	parent.Call("removeChild", node)

	return nil
}

func replaceNode(newNode, oldNode jsObject) error {
	if newNode == oldNode {
		return nil
	}

	parent := oldNode.Get("parentNode")
	// TODO: parent == nil || parent == js.Undefined
	if parent == nil {
		return errMissingParent
	}
	parent.Call("replaceChild", newNode, oldNode)

	return nil
}

func appendHTML(next, parent *HTML) {
	parent.node.Call("appendChild", next.node)
}

func replaceElement(next, prev ComponentOrHTML) error {
	n, p := assertHTML(next), assertHTML(prev)
	if n == nil || n.node == nil || p == nil || p.node == nil {
		return errEmptyElement
	}

	return replaceNode(n.node, p.node)
}

func insertElementBefore(next, prev ComponentOrHTML) error {
	parent, err := parentNode(prev)
	if err != nil {
		return err
	}

	n, p := assertHTML(next), assertHTML(prev)
	if n == nil || n.node == nil {
		return errEmptyElement
	}
	if p == nil {
		p = &HTML{}
	}

	parent.Call("insertBefore", n.node, p.node)
	return nil
}

func insertElementAfter(next, prev ComponentOrHTML) error {
	p := assertHTML(prev)
	if p == nil {
		return errEmptyElement
	}

	return insertElementBefore(next, p.node.Get("nextSibling"))
}

func parentNode(e ComponentOrHTML) (jsObject, error) {
	h := assertHTML(e)
	if h == nil || h.node == nil {
		return nil, errEmptyElement
	}

	parent := h.node.Get(`parentNode`)
	// TODO: parent == nil || parent == js.Undefined
	if parent == nil {
		return nil, errMissingParent
	}

	return parent, nil
}
