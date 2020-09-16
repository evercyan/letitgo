package cli

import (
	"errors"
	"fmt"
	"reflect"
)

// 符号
type border struct {
	H  rune // 水平 ─
	V  rune // 垂直 │
	VH rune // 水平垂直(交) ┼
	HU rune // 水平垂直(上) ┴
	HD rune // 水平垂直(下) ┬
	VL rune // 垂直水平(左) ┤
	VR rune // 垂直水平(右) ├
	DL rune // 转弯(下左) ┐
	DR rune // 转弯(下右) ┌
	UL rune // 转弯(上左) ┘
	UR rune // 转弯(上右) └
}

// 边线样式
// ref: http://www.tamasoft.co.jp/en/general-info/unicode.html
var styles = map[tStyle]border{
	Solid:  {'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'},
	Dashed: {'-', '|', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
	Dotted: {'*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*'},
}

// 字符集
var chinese = []rune{0x2E80, 0x9FD0}

var (
	errHeader = errors.New("invalid header length")
	errType   = errors.New("only support slice or array")
)

// --------------------------------

// 边线样式
type tStyle int

const (
	Solid  tStyle = iota // 实线
	Dashed               // 虚线(类 mysql 终端表格)
	Dotted               // 点线
)

func (s tStyle) valid() bool {
	switch s {
	case Solid, Dashed, Dotted:
		return true
	default:
		return false
	}
}

// --------------------------------

type table struct {
	elem   interface{}
	title  string     // 标题
	style  tStyle     // 样式
	header []string   // 头部
	widths []int      // 列宽
	rows   [][]string // 数据
}

// Table 终端打印表格, 支持 []struct, []slice([]int, []string...)
func Table(elem interface{}) *table {
	return &table{
		elem:   elem,
		title:  "",
		style:  Solid,
		header: []string{},
		widths: []int{},
		rows:   [][]string{},
	}
}

// Output ...
func (t *table) Output() {
	if t.title != "" {
		fmt.Println(t.title)
	}
	fmt.Println(t.Content() + "\n")
}

// Content ...
func (t *table) Content() (content string) {
	err := t.parse()
	if err != nil {
		return err.Error()
	}
	defer func() {
		if err := recover(); err != nil {
			content = string(fmt.Sprintf("%v", err))
		}
	}()
	return t.text()
}

// Style ...
func (t *table) Style(s tStyle) *table {
	if s.valid() {
		t.style = s
	}
	return t
}

// Header ...
func (t *table) Header(header []string) *table {
	if len(header) > 0 {
		t.header = header
	}
	return t
}

// Title ...
func (t *table) Title(title string) *table {
	t.title = title
	return t
}

// parse 解析数据, 生成 header, widths, rows
func (t *table) parse() (err error) {
	value := reflect.ValueOf(t.elem)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return errType
	}

	list := make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		list[i] = value.Index(i).Interface()
	}

	for index, item := range list {
		iv, it := reflect.ValueOf(item), reflect.TypeOf(item)
		if iv.Kind() == reflect.Ptr {
			iv = iv.Elem()
			it = it.Elem()
		}
		// 如果设置 Header,  需要校验与字段数量是否匹配
		headerLen := len(t.header)
		if iv.Kind() == reflect.Struct {
			// struct
			if headerLen > 0 && headerLen != iv.NumField() {
				return errHeader
			}
			row := []string{}
			for n := 0; n < iv.NumField(); n++ {
				cn := it.Field(n).Name
				cv := fmt.Sprintf("%+v", iv.FieldByName(cn).Interface())
				row = append(row, cv)
				// 首行解析 tag, 写 header
				if index == 0 {
					ct := it.Field(n).Tag.Get("table")
					if ct == "" {
						ct = cn
					}
					if headerLen == 0 {
						t.header = append(t.header, ct)
					}
					t.widths = append(t.widths, len(ct))
				}
				// 取最大宽度
				if t.widths[n] < len(cv) {
					t.widths[n] = len(cv)
				}
				// 如果有设置 header, 需要处理宽度
				if headerLen > 0 && len(t.header[n]) > t.widths[n] {
					t.widths[n] = len(t.header[n])
				}
			}
			t.rows = append(t.rows, row)
		} else if iv.Kind() == reflect.Slice || iv.Kind() == reflect.Array {
			if headerLen > 0 && headerLen != iv.Len() {
				return errHeader
			}
			row := []string{}
			for n := 0; n < iv.Len(); n++ {
				cv := fmt.Sprintf("%+v", iv.Index(n).Interface())
				row = append(row, cv)
				if index == 0 {
					t.widths = append(t.widths, len(cv))
				}
				if len(cv) > t.widths[n] {
					t.widths[n] = len(cv)
				}
				// 如果有设置 header, 需要处理宽度
				if headerLen > 0 && len(t.header[n]) > t.widths[n] {
					t.widths[n] = len(t.header[n])
				}
			}
			t.rows = append(t.rows, row)
		} else {
			return errType
		}
	}
	return err
}

// text 绘制显示文本
func (t *table) text() (table string) {
	if len(t.rows) == 0 {
		return table
	}
	b := styles[t.style]
	// 除内容区外, 有其余 头部上, 头部中, 头部下, 尾部
	headerT, headerM, headerB, footer := []rune{b.DR}, []rune{b.V}, []rune{b.VR}, []rune{b.UR}
	for i, width := range t.widths {
		// 每列长度 +2, 预留左右两空格位置
		headerT = append(headerT, []rune(t.repeat(b.H, width+2)+string(b.HD))...)
		headerB = append(headerB, []rune(t.repeat(b.H, width+2)+string(b.VH))...)
		footer = append(footer, []rune(t.repeat(b.H, width+2)+string(b.HU))...)

		if len(t.header) > 0 {
			// 有头部的情况下,内容区域可能出现非英文, 解析其实际显示长度
			l := width - t.getDisplayLentgh([]rune(t.header[i])) + 1
			headerM = append(headerM, []rune(" "+t.header[i]+t.repeat(' ', l)+string(b.V))...)
		}
	}
	// 替换最后结束符
	headerT[len(headerT)-1], headerB[len(headerB)-1], footer[len(footer)-1] = b.DL, b.VL, b.UL

	table += string(headerT) + "\n"
	if len(t.header) > 0 {
		// 有头部时加载头部中和头部下
		table += string(headerM) + "\n"
		table += string(headerB) + "\n"
	}

	// 内容区域
	for i, row := range t.rows {
		body := []rune{b.V}
		for i, width := range t.widths {
			l := width - t.getDisplayLentgh([]rune(row[i])) + 1
			body = append(body, []rune(" "+row[i]+t.repeat(' ', l)+string(b.V))...)
		}
		table += string(body) + "\n"
		if len(t.header) == 0 && i != len(t.rows)-1 {
			// 无头部时, 每行追加分隔线
			table += string(headerB) + "\n"
		}
	}

	table += string(footer)
	return table
}

// repeat 重复字符, 补全宽度用
func (t *table) repeat(char rune, num int) string {
	var s = make([]rune, num)
	for i := range s {
		s[i] = char
	}
	return string(s)
}

// getDisplayLentgh 获取显示长度
func (t *table) getDisplayLentgh(r []rune) int {
	length := len(r)
loop:
	for _, v := range r {
		if v >= chinese[0] && v <= chinese[1] {
			length++
			continue loop
		}
	}
	return length
}
