package main

import (
	"ai/src/ai/lib"
	"fmt"
)

/*
	DFS using channels and goroutines
*/

func main() {

	pathResultsCh := make(chan chan *lib.Path, lib.DEPTH_LIMIT)
	var allPaths []*lib.Path
	var actualPath *lib.Path

	n := lib.Node{}
	n.AddChild(&n)

	go dfs(888, &n, 1, &lib.Path{Found: false}, pathResultsCh)

	for result := range pathResultsCh {

		select {
		case actualPath = <-result:

			if actualPath.Found == true {
				allPaths = append(allPaths, actualPath)
			}

		default:

		}
	}

	for path := range allPaths {

		fmt.Printf("%v", path)
	}

}

func dfs(searchVal int64, n *lib.Node, depth int64, tentativePath *lib.Path, pathResultsCh chan chan *lib.Path) {

	v := n.GetVal()
	tentativePath.Path = append(tentativePath.Path, n)
	ch := make(chan *lib.Path)

	if v == searchVal {
		tentativePath.Found = true
	}

	numChildren := int64(len(n.GetChildren()))
	if numChildren == 0 {
		tentativePath.Found = false
	}

	pathResultsCh <- ch
	ch <- tentativePath

	if depth+1 < lib.DEPTH_LIMIT {

		for _, child := range n.GetChildren() {

			tentativePath.Found = false
			tentativePath.Path = append(tentativePath.Path, n)

			if numChildren >= depth/2 {
				go dfs(searchVal, child, depth+1, tentativePath, pathResultsCh)
			} else {
				dfs(searchVal, child, depth+1, tentativePath, pathResultsCh)
			}
		}
	}

	close(ch)
	return
}
