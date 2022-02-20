package lib

const DEPTH_LIMIT = 1024

type Node struct {
	val        int64
	visited    bool
	children   []*Node
	visitCount int64
}

type Path struct {
	Found      bool
	DepthLimit int64
	Path       []*Node
}

func (n *Node) Init(val int64) bool {

	n.val = val
	n.visited = false
	n.visitCount = 0

	return true
}

func (n *Node) GetChildren() []*Node {

	return n.children
}

func (n *Node) GetVal() int64 {

	return n.val
}

func (n *Node) AddChild(childNode *Node) int64 {

	n.children = append(n.children, childNode)
	return int64(len(n.children) - 1)
}

func (n *Node) RemoveChild(id *int64) bool {

	ok := false
	if int64(len(n.children)) > *id {
		n.children = append(n.children[0:*id], n.children[*id+1:]...)
		ok = true
	}

	return ok
}

func (n *Node) Child(id *int64) (bool, *Node) {

	ok := false
	var out *Node = nil

	if int64(len(n.children)) > *id {
		out = n.children[*id]
	}

	return ok, out
}
