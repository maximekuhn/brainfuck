package parser

type nodeStack struct {
	stack []*Node
}

func newNodeStack() *nodeStack {
	return &nodeStack{stack: make([]*Node, 0)}
}

// push accepts a Node and adds it to the stack.
//
// WARNING: if the input Node is nil, the stack behavior is undefined.
func (ns *nodeStack) push(n *Node) {
	ns.stack = append(ns.stack, n)
}

func (ns *nodeStack) pop() (*Node, bool) {
	if len(ns.stack) == 0 {
		return nil, false
	}
	idx := len(ns.stack) - 1
	n := ns.stack[idx]
	ns.stack = ns.stack[:idx]
	return n, true
}

func (ns *nodeStack) peek() (*Node, bool) {
	if len(ns.stack) == 0 {
		return nil, false
	}
	return ns.stack[len(ns.stack)-1], true
}
