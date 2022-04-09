package main

import (
	"ai/src/ai/lib"
	"fmt"
)

/*
	DFS using channels and goroutines
*/

func main() {

	n := lib.Node{}
	n.Init(0)
	n.AddChild(&n)
	search(&n)
}

func search(n *lib.Node) *[]*lib.Path {

	var allPaths *[]*lib.Path
	tmp := make([]*lib.Path, 1)
	allPaths = &tmp

	allPaths = dfs(888, *n, 1, &lib.Path{Found: false}, allPaths)

	for path := range *allPaths {

		fmt.Printf("%v", path)
	}

	return allPaths
}

func dfs(searchVal int, n lib.Node, depth int, tentativePath *lib.Path, pathResultsCh *[]*lib.Path) *[]*lib.Path {

	v := n.GetVal()
	tentativePath.Path = append(tentativePath.Path, &n)

	// Append tentative path to successful path results
	if v == searchVal {
		tentativePath.Found = true
		*pathResultsCh = append(*pathResultsCh, tentativePath)
	}

	tentativePath.Found = false
	if int64(len(n.GetChildren())) > 0 && depth+1 < lib.DepthLimit {

		for _, child := range n.GetChildren() {

			tmp := *tentativePath
			tmp.Path = append(tmp.Path, &n)
			dfs(searchVal, *child, depth+1, &tmp, pathResultsCh)
		}
	}

	return pathResultsCh
}
