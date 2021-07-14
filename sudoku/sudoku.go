package sudoku

import (
	"math/rand"
	"time"
)

// Level ...
type Level int

// Level ...
const (
	Easy   Level = iota // 简单
	Medium              // 中等
	Hard                // 困难
)

// Number 对应难度需要重置的数量
func (l Level) Number() int {
	switch l {
	case Easy:
		return 30
	case Medium:
		return 45
	case Hard:
		return 60
	default:
		return 0
	}
}

// Generate ...
func Generate(level Level) (board [9][9]int) {
	number := level.Number()
	if number == 0 {
		return board
	}

	// 生成数据
	board, _ = new(solver).solve([9][9]int{})

	// todo 当前仅随机遮掩, 待优化
	// 根据难度进行遮掩处理
	// 将九宫格转换为 0-80 下标的一维数组, 随机生成隐藏的下标, 置 0 即可
	keys := make([]int, 0)
	for i := 0; i < 81; i++ {
		keys = append(keys, i)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		index := rand.Intn(len(keys))

		board[keys[index]/9][keys[index]%9] = 0
		if index == len(keys)-1 {
			keys = keys[:index]
		} else {
			keys = append(keys[:index], keys[index+1:]...)
		}
	}

	return board
}

// Solve ...
func Solve(board [9][9]int) ([9][9]int, error) {
	return new(solver).solve(board)
}

// Verify 验证数据
func Verify(board [9][9]int) map[int]map[int]bool {
	result := make(map[int]map[int]bool)
	solver := new(solver)
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if _, ok := result[r]; !ok {
				result[r] = make(map[int]bool)
			}
			result[r][c] = board[r][c] > 0 && solver.isValid(board, r, c, board[r][c])
		}
	}
	return result
}
