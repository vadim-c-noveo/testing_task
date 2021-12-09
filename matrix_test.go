package matrix_solver

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOneDimensionalMatrix_At(t *testing.T) {
	test := NewMatrix()
	require.Equal(t, 1, test.At(0, 0))
	require.Equal(t, 4, test.At(0, 3))
	require.Equal(t, 6, test.At(1, 1))
	require.Equal(t, 12, test.At(2, 3))
}

func TestOneDimensionalMatrix_Write(t *testing.T) {
	test := NewMatrix()
	test.Write(0, 0, 100)
	test.Write(0, 3, 200)
	test.Write(1, 1, 300)
	test.Write(3, 4, 400)
	require.Equal(t, 100, test.At(0, 0))
	require.Equal(t, 200, test.At(0, 3))
	require.Equal(t, 300, test.At(1, 1))
	require.Equal(t, 400, test.At(3, 4))
}

func NewMatrix() OneDimensionalMatrix {
	return OneDimensionalMatrix{
		arr: [33124]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		n:   3,
		m:   4,
	}
}
