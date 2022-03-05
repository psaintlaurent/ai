package main

import (
	"ai/src/ai/lib"
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestSearch(t *testing.T) {

	var x, y, i, breadth, randVal int64
	var currentRow, nextRow []*lib.Node
	var allPaths []*lib.Path
	var root *lib.Node
	var searchValSetFlag bool
	const SearchVal = 888

	depth := rand.Int63n(10)
	root = &lib.Node{Val: 0, Visited: false, VisitCount: 0, Children: []*lib.Node{}}
	currentRow = append(currentRow, root)

	for x = 0; x <= depth; x++ {
		for y = 0; y < int64(len(currentRow)); y++ {

			breadth = rand.Int63n(10)

			for i = 0; i < breadth; i++ {

				if x == 5 && searchValSetFlag == false {

					fmt.Printf("Adding the search value: %d", SearchVal)
					currentRow[y].AddChild(&lib.Node{Val: SearchVal, Visited: false, VisitCount: 0, Children: []*lib.Node{}})
					searchValSetFlag = true
				}

				randVal = rand.Int63n(math.MaxInt64)
				fmt.Printf("Adding new node with random value %d at depth %d and breadth %d.", randVal, x, y)
				currentRow[y].AddChild(&lib.Node{Val: randVal, Visited: false, VisitCount: 0, Children: []*lib.Node{}})

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
