package parser

type NodeType int

const (
	NodeIncrement NodeType = iota
	NodeDecrement
	NodeNext
	NodePrevious
	NodeOutput
	NodeInput
	NodeLoop
)

type Ast struct {
	Statements []*Node
}

type Node struct {
	Type  NodeType
	Child []*Node // only used for NodeLoop (nil for all other node types)
}
