package main

import (
	"ai/src/ai/lib"
	"fmt"
)

"""
	DFS using channels and goroutines
"""

func main() {

	pathResultsCh := make(chan chan *lib.Path)
	var allPaths []*lib.Path
	var actualPath *lib.Path

	n := lib.Node{}
	n.AddChild(&n)

	go dfs(888, &n, 1, &lib.Path{Found: false}, pathResultsCh)

	for result := range pathResultsCh {

		select {

			case actualPath = <-result:
				allPaths = append(allPaths, actualPath)
			default:
				actualPath = nil
		}

	}

	for path := range allPaths {

		fmt.Printf("%v", path)
	}

}

func dfs(searchVal int64, n *lib.Node, depth int64, tentativePath *lib.Path, pathResultsCh chan chan *lib.Path) {

	v := n.GetVal()

	if v == searchVal {


		tentativePath.Path = append(tentativePath.Path, n)
		tentativePath.Found = true

		ch := make(chan *lib.Path)
		pathResultsCh <-ch
		ch <-tentativePath

	}

	numChildren := int64(len(n.GetChildren()))

	if numChildren == 0 {

		tentativePath.Path = append(tentativePath.Path, n)
		tentativePath.Found = false
		pathResultsCh <- tentativePath
		return
	}

	for _, child := range n.GetChildren() {

		if numChildren >= depth:
			go dfs(searchVal, child, depth+1, )

	}

	if depth == 1 {

		close
	}

	return
}
