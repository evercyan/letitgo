package cmdline

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	// 样式
	for a := Reset; a <= CrossedOut; a++ {
		// 背景颜色
		for b := BgBlack; b <= BgWhite; b++ {
			// 字体颜色
			for c := Black; c <= White; c++ {
				text := a.String() + " " + b.String() + " " + c.String()
				fmt.Println(Style(text).Attr(a).Color(c).BgColor(b).Text())
			}
		}
	}
	fmt.Println(Style("Coverage").Color(100).BgColor(100).Attr(100).Text())

	// debug
	// t.FailNow()
}
