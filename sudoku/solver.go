package sudoku

import (
	"errors"
)

var (
	// 初始数组, 用 map 可以在 range 时替代 shuffle([]int)
	numbers = map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
	}
)

// solver ...
type solver struct{}

// solve 支持传入空数据进行初始化
func (s *solver) solve(data [9][9]int) ([9][9]int, error) {
	board, err := s.handle(data, 0)
	if err != nil {
		return data, err
	}
	return board, nil
}

// handle 把 board 转换成一维数组后，k 为索引值
func (s *solver) handle(board [9][9]int, k int) ([9][9]int, error) {
	if k == 81 {
		return board, nil
	}

	// 二维数组中的索引值
	r, c := k/9, k%9
	// 元素 >0, 表示已有数字, 处理下一个
	if board[r][c] > 0 {
		return s.handle(board, k+1)
	}

	for number := range numbers {
		if s.isValid(board, r, c, number) {
			board[r][c] = number
			// 每填入一个数后, 递规去解决下一个 key, 直到所有都满足
			if board, err := s.handle(board, k+1); err == nil {
				return board, nil
			}
		}
	}

	return board, errors.New("can not solve")
}

// isValid 检查 num 是否 可放置在 (r, c)
func (s *solver) isValid(board [9][9]int, r int, c int, num int) bool {
	// bi, bj 是 rc 所在块的左上角元素的索引值
	bi, bj := r/3*3, c/3*3
	for n := 0; n < 9; n++ {
		// 行检查, 列检查, 九宫格检查
		if board[r][n] == num || board[n][c] == num || board[bi+n/3][bj+n%3] == num {
			return false
		}
	}
	return true
}
