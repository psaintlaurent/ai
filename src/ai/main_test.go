package main

import (
	"ai/src/ai/lib"
	"math"
	"math/rand"
	"testing"
)

func TestMain(m *testing.M) {

	var breadth int64
	var current_row, next_row []*lib.Node
	var root *lib.Node
	var x int64

	depth := rand.Int63n(50)
	current_row = append(current_row, root)


	for x = 0; x <= depth; x++ {

		for y = 0; y<= ; y++ {

			breadth = rand.Int63n(50)


		}

		current_row = next_row
	}

	search(node)

}
