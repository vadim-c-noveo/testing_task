package matrix_solver

import (
	"fmt"
)

func OnesAbsenceErr() error {
	return fmt.Errorf("no 1's found at the input. Skipping to next case")
}

func ParsingError(err error) error {
	return fmt.Errorf("error while parsing: %s. Cannot continue", err)
}

func ConversionError(err error) error {
	return fmt.Errorf("error (%s) while converting", err)
}

func MatrixDimensionsErr(rowLen, m int) error{
	return fmt.Errorf("length of the last row(%d) exceeds expected %d", rowLen, m)
}