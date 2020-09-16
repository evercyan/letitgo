# letitgo/cli

> 解决日常开发中, 一些终端的特殊样式需求, 如字体颜色样式, 表格等

## Style 字体颜色样式

通常是这样解决的, 但难免不够灵活

```go
fmt.Println(fmt.Sprintf("\033[1;31m错误: %s\033[0m", text))
```

### 字体属性

`基于 ANSI 控制码`

```go
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
```

### 示例

```go
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
```

效果图如下, 部分如 `慢闪` `快闪` `隐藏` 等效果未起作用, 具体以实际为准

![cli-style](https://raw.githubusercontent.com/evercyan/cantor/master/resource/6a/6a1114f0d4db55ddd90dce8eb537e182.png)

---

## Table 渲染表格

- 支持 []struct, []interface 数据类型渲染
- 支持多样式边框
- 支持取 struct tag `table` 作为表格 header
- 支持自定义表格 header, title
- 配合字体颜色组件 Style 使用更香

```go
type user struct {
    Name string `json:"name" table:"名称"`
    Age  int    `json:"age"`
}

var (
    structList = []user{
        {"Stark", 20},
        {"Lannister", 21},
    }
    structPtrList = []*user{
        {
            "Stark",
            20,
        },
        {
            "Lannister",
            21,
        },
    }
    numberList = [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    stringList = [][]string{
        {"a", "bb", "ccc"},
        {"dddd", "eeeee", "ffffff"},
    }
)

// style
for s := Dashed; s <= Dotted; s++ {
    Table(structList).Style(s).Output()
}

// header
Table(structList).Header([]string{"Cooooooooool1", "Col2"}).Style(Solid).Output()
Table(numberList).Header([]string{"Col1", "Col2", "Col3"}).Style(Solid).Output()
Table(stringList).Header([]string{"Col1", "Col2", "Col3"}).Style(Solid).Output()

// color
content := Table(structList).Style(Solid).Content()
fmt.Println(Style(content).Color(Red).Text())
```

![cli-table](https://raw.githubusercontent.com/evercyan/cantor/master/resource/6c/6c3c32f6c87e5caea8239ccfa6db887d.png)
