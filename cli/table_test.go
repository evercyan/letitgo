package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestTableStyle(t *testing.T) {
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
}

func TestTableCoverage(t *testing.T) {
	assert.NotEmpty(t, Table(numberList).Style(100).Content())
	assert.NotNil(t, Table(1).Content())
	assert.NotNil(t, Table(structPtrList).Content())
	assert.NotNil(t, Table([]int{1}).Content())
	assert.NotNil(t, Table([][]map[string]string{
		{
			{
				"a": "a",
			},
		},
	}).Content())
	assert.NotNil(t, Table([][]int{}).Content())
	assert.NotNil(t, Table(structList).Header([]string{"c1"}).Content())
	assert.NotNil(t, Table(numberList).Header([]string{"c1"}).Content())
	// debug
	// t.FailNow()
}
