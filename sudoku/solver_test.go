package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSudokuSolverCoverage(t *testing.T) {
	data := [9][9]int{}
	data[0] = [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	result, err := new(solver).solve(data)
	assert.NotNil(t, err)
	assert.Equal(t, data, result)
}
