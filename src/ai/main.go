package main

import "ai/src/ai/lib"

func main() {

	n := lib.Node{}
	n.AddChild(&n)
}

func dfs(searchVal int64, n *lib.Node) (bool, []*lib.Node) {

	v := n.GetVal()
	if v == searchVal {

		tmp := make([]*lib.Node, 1)
		tmp = append(tmp, n)
		return true, tmp
	} else {

		var found bool = false
		var out []*lib.Node

		for _, child := range n.GetChildren() {

			found, out = dfs(searchVal, child)
			if found == true {
				break
			}
		}

		if found == true {
			out = append(out, n)
		}

		return true, out
	}

	return false, nil
}
