package lib

const DepthLimit = 1024

type Node struct {
	Val        int
	Visited    bool
	Children   []*Node
	VisitCount int
}

type Path struct {
	Found      bool
	DepthLimit int64
	Path       []*Node
}

func (n *Node) Init(val int) bool {

	n = &Node{Val: val, Visited: false, VisitCount: 0, Children: []*Node{}}
	return true
}

func (n *Node) GetChildren() []*Node {

	return n.Children
}

func (n *Node) GetVal() int {

	return n.Val
}

func (n *Node) AddChild(childNode *Node) int {

	n.Children = append(n.Children, childNode)
	return len(n.Children) - 1
}

func (n *Node) RemoveChild(id *int64) bool {

	ok := false
	if int64(len(n.Children)) > *id {
		n.Children = append(n.Children[0:*id], n.Children[*id+1:]...)
		ok = true
	}

	return ok
}

func (n *Node) Child(id *int64) (bool, *Node) {

	ok := false
	var out *Node = nil

	if int64(len(n.Children)) > *id {
		out = n.Children[*id]
	}

	return ok, out
}
