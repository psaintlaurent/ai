package main

import (
	"ai/src/ai/lib"
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {

	var x, y, i, breadth int64
	var currentRow, nextRow []*lib.Node
	var allPaths []*lib.Path
	var root *lib.Node
	var searchValSet bool

	depth := rand.Int63n(50)
	root = &lib.Node{Val: 0, Visited: false, VisitCount: 0, Children: []*lib.Node{}}

	currentRow = append(currentRow, root)

	for x = 0; x <= depth; x++ {
		for y = 0; y < int64(len(currentRow)); y++ {

			breadth = rand.Int63n(50)
			for i = 0; i < breadth; i++ {

				if x == 25 && searchValSet == false {

					currentRow[y].AddChild(&lib.Node{Val: 0, Visited: false, VisitCount: 0, Children: []*lib.Node{}})
					searchValSet = true
				}

				currentRow[y].AddChild(&lib.Node{Val: 888, Visited: false, VisitCount: 0, Children: []*lib.Node{}})
			}
		}
		currentRow = nextRow
	}

	allPaths = search(root)

	if len(allPaths) > 0 {
		t.Logf("Success!")
	} else {
		t.Errorf("Failure!")
	}
}
