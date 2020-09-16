package cli

import (
	"fmt"
)

// 样式
type sAttr int

const (
	Reset        sAttr = iota // 默认
	Bold                      // 加粗
	Faint                     // 弱化
	Italic                    // 斜体
	Underline                 // 下划
	BlinkSlow                 // 慢闪
	BlinkRapid                // 快闪
	ReverseVideo              // 反转
	Concealed                 // 隐藏
	CrossedOut                // 中划
)

func (a sAttr) String() string {
	switch a {
	case Reset:
		return "默认"
	case Bold:
		return "加粗"
	case Faint:
		return "弱化"
	case Italic:
		return "斜体"
	case Underline:
		return "下划"
	case BlinkSlow:
		return "慢闪"
	case BlinkRapid:
		return "快闪"
	case ReverseVideo:
		return "反转"
	case Concealed:
		return "隐藏"
	case CrossedOut:
		return "中划"
	default:
		return ""
	}
}

// 字体颜色
type sColor int

const (
	Black   sColor = iota + 30 // 黑色字体
	Red                        // 红色字体
	Green                      // 绿色字体
	Yellow                     // 黄色字体
	Blue                       // 蓝色字体
	Magenta                    // 洋红字体
	Cyan                       // 青色字体
	White                      // 白色字体
)

func (c sColor) String() string {
	switch c {
	case Black:
		return "黑色字体"
	case Red:
		return "红色字体"
	case Green:
		return "绿色字体"
	case Yellow:
		return "黄色字体"
	case Blue:
		return "蓝色字体"
	case Magenta:
		return "洋红字体"
	case Cyan:
		return "青色字体"
	case White:
		return "白色字体"
	default:
		return ""
	}
}

// 背景颜色
type sBgColor int

const (
	BgBlack   sBgColor = iota + 40 // 黑色背景
	BgRed                          // 红色背景
	BgGreen                        // 绿色背景
	BgYellow                       // 黄色背景
	BgBlue                         // 蓝色背景
	BgMagenta                      // 洋红背景
	BgCyan                         // 青色背景
	BgWhite                        // 白色背景
)

func (bc sBgColor) String() string {
	switch bc {
	case BgBlack:
		return "黑色背景"
	case BgRed:
		return "红色背景"
	case BgGreen:
		return "绿色背景"
	case BgYellow:
		return "黄色背景"
	case BgBlue:
		return "蓝色背景"
	case BgMagenta:
		return "洋红背景"
	case BgCyan:
		return "青色背景"
	case BgWhite:
		return "白色背景"
	default:
		return ""
	}
}

type style struct {
	text    string
	color   sColor
	bgColor sBgColor
	attr    sAttr
}

// Style ...
func Style(text string) *style {
	return &style{
		text: text,
	}
}

func (s *style) Color(color sColor) *style {
	if color.String() != "" {
		s.color = color
	}
	return s
}

func (s *style) BgColor(bgColor sBgColor) *style {
	if bgColor.String() != "" {
		s.bgColor = bgColor
	}
	return s
}

func (s *style) Attr(attr sAttr) *style {
	if attr.String() != "" {
		s.attr = attr
	}
	return s
}

func (s *style) Text() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, s.attr, s.bgColor, s.color, s.text, 0x1B)
}
