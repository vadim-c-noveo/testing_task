/*
	Package matrix_solver implements solution of bitmap matrix distance calculation
	using Breadth First Search algorithm.
*/

package matrix_solver

import (
	"bufio"
	"math"
	"os"
)

func Run() {
	matrix := OneDimensionalMatrix{}
	distanceCalculator := NewDistanceCalculator()
	reader := bufio.NewReader(os.Stdin)
	err := Solve(distanceCalculator, matrix, reader)
	if err != nil {
		panic(err)
	}
}

// DistanceCalculator contains ancillary data needed for calculation
type DistanceCalculator struct {
	moves [4][2]int
}

func NewDistanceCalculator() DistanceCalculator {
	dc := DistanceCalculator{moves: [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}}
	return dc
}

type MatrixProcessor interface {
	processMatrix(*OneDimensionalMatrix) error
}

func (dc DistanceCalculator) processMatrix(matrix *OneDimensionalMatrix) error {
	queue := populateQueue(matrix)

	if len(queue) == 0 {
		return OnesAbsenceErr()
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, move := range dc.moves {
			// let's check neighbours of current
			x, y := current[0]+move[0], current[1]+move[1]
			if x < 0 || x >= matrix.n || y < 0 || y >= matrix.m || matrix.At(x, y) <= matrix.At(current[0], current[1]) {
				// matrix bounds check and skipping of already set smaller distances
				continue
			}

			queue = append(queue, [2]int{x, y})
			matrix.Write(x, y, matrix.At(current[0], current[1]) + 1)
		}
	}
	return nil
}

// Pre-populate queue with locations of whites
func populateQueue(matrix *OneDimensionalMatrix) [][2]int {
	var queue [][2]int
	// let's find all whites(1) and add them to queue
	for i := 0; i < matrix.n; i++ {
		for j := 0; j < matrix.m; j++ {
			if matrix.At(i, j) == 1 {
				matrix.Write(i, j, 0)
				queue = append(queue, [2]int{i, j})
			} else {
				// all blacks are set to max for further comparison
				matrix.Write(i, j, math.MaxInt64)
			}
		}
	}
	return queue
}
