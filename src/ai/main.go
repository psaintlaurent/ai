package main

import "ai/src/ai/lib"

func main() {

	n := lib.Node{}
	n.AddChild(&n)
}

func dfs(searchVal int64, n *lib.Node) *lib.Path {

	v := n.GetVal()
	var path *lib.Path

	if v == searchVal {

		tmp := make([]*lib.Node, 1)
		tmp = append(tmp, n)
		return &lib.Path{Found: true, Path: tmp}
	} else {

		for _, child := range n.GetChildren() {

			path = dfs(searchVal, child)
			if path.Found == true {
				break
			}
		}

		if path.Found == true {
			path.Path = append(path.Path, n)
			return path
		}
	}

	path.Found = false

	return path
}
