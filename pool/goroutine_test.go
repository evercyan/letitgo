package pool

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewGoroutine(t *testing.T) {
	result := make(map[int64][]int)

	g := NewGoroutine(3)
	for i := 0; i < 9; i++ {
		go func(i int) {
			defer g.Done()
			g.Add(1)

			sec := time.Now().Unix()
			if _, ok := result[sec]; !ok {
				result[sec] = []int{}
			}
			result[sec] = append(result[sec], i)

			fmt.Printf("go func: %d, time: %d\n", i, sec)
			time.Sleep(time.Second)
		}(i)
	}
	g.Wait()

	for _, nums := range result {
		assert.Equal(t, 3, len(nums))
	}

	// go func: 8, time: 1617161514
	// go func: 0, time: 1617161514
	// go func: 3, time: 1617161514
	//
	// go func: 6, time: 1617161515
	// go func: 7, time: 1617161515
	// go func: 1, time: 1617161515
	//
	// go func: 2, time: 1617161516
	// go func: 5, time: 1617161516
	// go func: 4, time: 1617161516
}
