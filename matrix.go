package matrix_solver

import (
	"fmt"
	"log"
)

// Taken from task description as max possible width/height of input matrix
const maxSize = 182

// OneDimensionalMatrix Encapsulates matrix in form of an array
type OneDimensionalMatrix struct {
	arr [maxSize * maxSize]int
	n   int
	m   int
}

// At returns element at [i][j] position
func (receiver *OneDimensionalMatrix) At(i, j int) int {
	return receiver.arr[receiver.addressOf(i, j)]
}

// Write assigns element at [i][j] position
func (receiver *OneDimensionalMatrix) Write(i, j, val int) {
	receiver.arr[receiver.addressOf(i, j)] = val
}

// Display prints contained matrix in the implied form
func (receiver *OneDimensionalMatrix) Display() {
	for i := 0; i < receiver.n; i++ {
		for j := 0; j < receiver.m; j++ {
			fmt.Printf("%v ", receiver.At(i, j))
		}
		fmt.Println()
	}
}

func (receiver *OneDimensionalMatrix) addressOf(i, j int) int {
	if i > receiver.n || j > receiver.m {
		log.Fatalf("attempt to access element that's out of bounds of matrix")
	}
	return (receiver.m * i) + j
}
