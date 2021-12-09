package matrix_solver

import (
	"bufio"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockReader struct {
	bufio.Reader
	ReadStringFunc func(delim byte) (string, error)
}

func (r *MockReader) ReadString(delim byte) (string, error) {
	if r.ReadStringFunc != nil {
		return r.ReadStringFunc(delim)
	}
	return r.Reader.ReadString(delim)
}

func TestReadRow(t *testing.T) {
	matrix := OneDimensionalMatrix{n: 2, m: 2}
	want := []int{0,0,0,1}

	mockReader := MockReader{ReadStringFunc: func(delim byte) (string, error) {
		return "01\n", nil
	}}
	err := ReadRow(&mockReader, &matrix, 1)
	require.NoError(t, err)
	require.Equal(t, want, matrix.arr[:4])
}

func TestReadRow_MatrixDimensionsErr(t *testing.T) {
	matrix := OneDimensionalMatrix{n: 2, m: 2}

	mockReader := MockReader{ReadStringFunc: func(delim byte) (string, error) {
		return "001\n", nil
	}}
	err := ReadRow(&mockReader, &matrix, 1)
	require.Error(t, err)
}

func TestReadRow_IncorrectInput(t *testing.T) {
	matrix := OneDimensionalMatrix{n: 2, m: 2}

	mockReader := MockReader{ReadStringFunc: func(delim byte) (string, error) {
		return "0A\n", nil
	}}
	err := ReadRow(&mockReader, &matrix, 1)
	require.Error(t, err)
}


