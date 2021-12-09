package matrix_solver

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPopulateQueue(t *testing.T) {
	tests := []struct {
		input OneDimensionalMatrix
		want  [][2]int
	}{
		{OneDimensionalMatrix{[maxSize * maxSize]int{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0}, 3, 4},
			[][2]int{
				{0, 3},
				{1, 2}, {1, 3},
				{2, 1}, {2, 2},
			}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Executing test data #%d", i), func(t *testing.T) {
			got := populateQueue(&tc.input)
			for j, want := range tc.want {
				require.Equalf(t, got, tc.want, "got %v; want %v", got[j], want)
			}
		})
	}
}

func TestProcessMatrix(t *testing.T) {
	tests := []struct {
		input OneDimensionalMatrix
		want  []int
	}{
		{OneDimensionalMatrix{[maxSize * maxSize]int{0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0}, 3, 4},
			[]int{3, 2, 1, 0, 2, 1, 0, 0, 1, 0, 0, 1},
		},
		{
			OneDimensionalMatrix{[maxSize * maxSize]int{0, 0, 0, 1}, 4, 1},
			[]int{3, 2, 1, 0},
		},
		{
			OneDimensionalMatrix{[maxSize * maxSize]int{0, 0, 0, 1}, 1, 4},
			[]int{3, 2, 1, 0},
		},
	}
	distanceCalculator := NewDistanceCalculator()
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Executing test data #%d", i), func(t *testing.T) {
			err := distanceCalculator.processMatrix(&tc.input)
			require.NoError(t, err)
			solution := tc.input.arr[:len(tc.want)]
			require.Equalf(t, tc.want, solution, "got %v; want %v", solution, tc.want)
		})
	}
}

func TestProcessMatrix_ZeroedInput(t *testing.T) {
	got := OneDimensionalMatrix{[maxSize * maxSize]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, 3}
	distanceCalculator := NewDistanceCalculator()

	t.Run("No 1's found in input", func(t *testing.T) {
		err := distanceCalculator.processMatrix(&got)
		expectedError := "no 1's found at the input. Skipping to next case"
		require.Error(t, err, expectedError)
	})
}
