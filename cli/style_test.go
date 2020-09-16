package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	// 样式
	for a := Reset; a <= CrossedOut; a++ {
		// 背景颜色
		for b := BgBlack; b <= BgWhite; b++ {
			// 字体颜色
			for c := Black; c <= White; c++ {
				text := a.String() + " " + b.String() + " " + c.String()
				fmt.Printf(Style(text).Attr(a).Color(c).BgColor(b).Text() + " ")
				assert.NotEmpty(t, Style(text).Attr(a).Color(c).BgColor(b).Text())
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
	fmt.Println(Style("Coverage").Color(100).BgColor(100).Attr(100).Text())
}
