package sudoku

import (
	"testing"

	"github.com/evercyan/letitgo/cli"
	"github.com/stretchr/testify/assert"
)

func TestSudoku(t *testing.T) {
	// t.FailNow()
}

func TestSudokuGenarate(t *testing.T) {
	cli.Table(Generate(Easy)).Title("Sukudo 生成 - Easy").Output()
	cli.Table(Generate(Medium)).Title("Sukudo 生成 - Medium").Output()
	cli.Table(Generate(Hard)).Title("Sukudo 生成 - Hard").Output()
	cli.Table(Generate(1000)).Title("Sukudo 生成 - 无效难度").Output()
}

func TestSudokuSolve(t *testing.T) {
	data := [9][9]int{
		{0, 0, 9, 7, 4, 8, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 1, 0, 9, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 2, 4, 0},
		{0, 6, 4, 0, 1, 0, 5, 9, 0},
		{0, 9, 8, 0, 0, 0, 3, 0, 0},
		{0, 0, 0, 8, 0, 3, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 0, 2, 7, 5, 9, 0, 0},
	}
	result, err := Solve(data)
	assert.Nil(t, err)
	cli.Table(result).Title("解决 Sukudo").Output()
}

func TestSudokuVerify(t *testing.T) {
	data := [9][9]int{}
	result := Verify(data)
	assert.Equal(t, map[int]map[int]bool{
		0: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		1: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		2: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		3: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		4: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		5: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		6: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		7: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
		8: {0: false, 1: false, 2: false, 3: false, 4: false, 5: false, 6: false, 7: false, 8: false},
	}, result)
}
