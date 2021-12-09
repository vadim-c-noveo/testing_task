package matrix_solver

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type StringReader interface{
	io.Reader
	ReadString(delim byte) (string, error)
}

// Solve retrieves user input from stdin, runs solver for calculations
//   and prints back results
func Solve(solver MatrixProcessor, matrix OneDimensionalMatrix, reader StringReader) error {
	numberOfCases, err := reader.ReadString('\n')
	if err != nil {
		return ParsingError(err)
	}
	numberOfCasesInt, err := strconv.Atoi(withoutDelim(numberOfCases))
	if err != nil {
		return ConversionError(err)
	}

	for caseIdx := 0; caseIdx < numberOfCasesInt; caseIdx++ {
		err := RetrieveCase(reader, &matrix)
		if err != nil {
			return err
		}
		err = solver.processMatrix(&matrix)
		if err != nil {
			fmt.Println(err)
			continue
		}
		matrix.Display()
		fmt.Println()
	}
	return nil
}

// RetrieveCase collects matrix data via reader, converts it to int and returns it
func RetrieveCase(reader StringReader, matrix *OneDimensionalMatrix) error {
	matrixDimensions, err := reader.ReadString('\n')
	if err != nil {
		return ParsingError(err)
	}
	matrixDimensionsArr := strings.Split(matrixDimensions, " ")

	matrix.n, err = strconv.Atoi(matrixDimensionsArr[0])
	if err != nil {
		return ConversionError(err)
	}
	matrix.m, err = strconv.Atoi(withoutDelim(matrixDimensionsArr[1]))
	if err != nil {
		return ConversionError(err)
	}

	for i := 0; i < matrix.n; i++ {
		err = ReadRow(reader, matrix, i)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadRow reads row of input, writes to matrix
func ReadRow(reader StringReader, matrix *OneDimensionalMatrix, i int) error {
	var tmp int
	row, err := reader.ReadString('\n')
	if err != nil {
		return ParsingError(err)
	}
	if len(withoutDelim(row)) > matrix.m {
		return MatrixDimensionsErr(len(row)-1, matrix.m)
	}

	for j, rn := range withoutDelim(row) {
		tmp, err = strconv.Atoi(string(rn))
		if err != nil {
			return ConversionError(err)
		}
		matrix.Write(i, j, tmp)
	}
	return nil
}

func withoutDelim(str string) string {
	return str[:len(str)-1]
}