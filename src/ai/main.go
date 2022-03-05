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

func search(n *lib.Node) []*lib.Path {

	pathResultsCh := make(chan *lib.Path, lib.DepthLimit)
	var allPaths []*lib.Path
	var actualPath *lib.Path

	go dfs(888, *n, 1, &lib.Path{Found: false}, pathResultsCh)

	for {

		result, ok := <-pathResultsCh
		if !ok {
			break
		}
		if result.Found {
			allPaths = append(allPaths, actualPath)
		}
	}

	for path := range allPaths {

		fmt.Printf("%v", path)
	}

	return allPaths
}

func dfs(searchVal int64, n lib.Node, depth int64, tentativePath *lib.Path, pathResultsCh chan *lib.Path) {

	v := n.GetVal()
	tentativePath.Path = append(tentativePath.Path, &n)

	if v == searchVal {
		tentativePath.Found = true
	}

	numChildren := int64(len(n.GetChildren()))
	if numChildren == 0 {
		tentativePath.Found = false
	}

	pathResultsCh <- tentativePath

	if depth+1 < lib.DepthLimit {

		for _, child := range n.GetChildren() {

			tentativePath.Found = false
			tentativePath.Path = append(tentativePath.Path, &n)

			if numChildren >= depth/2 {
				go dfs(searchVal, *child, depth+1, tentativePath, pathResultsCh)
			} else {
				dfs(searchVal, *child, depth+1, tentativePath, pathResultsCh)
			}
		}
	}

	return
}
