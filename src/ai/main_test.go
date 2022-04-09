package main

import (
	"ai/src/ai/lib"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {

	var x, y, i, breadth, randVal int
	var currentRow, nextRow []*lib.Node
	var allPaths *[]*lib.Path
	var root *lib.Node
	var searchValSetFlag bool
	const SearchVal = 888
	var max, min int

	rand.Seed(time.Now().UnixNano())
	depth := rand.Intn(10)
	depth = rand.Intn(max-min+1) + min

	root = &lib.Node{Val: 0, Visited: false, VisitCount: 0, Children: []*lib.Node{}}
	currentRow = append(currentRow, root)

	for x = 0; x <= depth; x++ {

		for y = 0; y < len(currentRow); y++ {

			breadth = rand.Intn(10)
			breadth = rand.Intn(max-min+1) + min

			for i = 0; i < breadth; i++ {

				if x == 5 && searchValSetFlag == false {

					fmt.Printf("Adding the search value: %d", SearchVal)
					tmp := &lib.Node{Val: SearchVal, Visited: false, VisitCount: 0, Children: []*lib.Node{}}
					currentRow[y].AddChild(tmp)
					nextRow = append(nextRow, tmp)
					searchValSetFlag = true
				}

				fmt.Printf("Adding new node with random value %d at depth %d and breadth %d. \n", randVal, x, i)
				tmp := &lib.Node{Val: rand.Intn(math.MaxInt64), Visited: false, VisitCount: 0, Children: []*lib.Node{}}
				currentRow[y].AddChild(tmp)
				nextRow = append(nextRow, tmp)

			}
		}
		currentRow = nextRow
	}

	allPaths = search(root)

	if len(*allPaths) > 0 {
		t.Logf("Success!")
	} else {
		t.Errorf("Failure!")
	}
}
