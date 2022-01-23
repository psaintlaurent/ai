package lib

type Node struct {
	val        int64
	visited    bool
	children   []*Node
	visitCount int64
}

type Path struct {
	Found bool
	Path  []*Node
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

func (n *Node) AddChild(childNode *Node) int {

	n.children = append(n.children, childNode)
	return len(n.children) - 1
}

func (n *Node) RemoveChild(id int) bool {

	ok := false
	if len(n.children) > id {
		n.children = append(n.children[0:id], n.children[id+1:]...)
		ok = true
	}

	return ok
}

func (n *Node) Child(id *int) (bool, *Node) {

	ok := false
	var out *Node = nil

	if len(n.children) > *id {
		out = n.children[*id]
	}

	return ok, out
}
