package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonstr string = `{
    "name": "hello",
    "detail": {
    	"age": 20,
        "height": "175cm",
        "weight": "60kg"
    },
    "langs": [
        "php",
        "golang",
        "python",
        "shell"
    ]
}`

func TestJson(t *testing.T) {
	name := Json(jsonstr).Key("name").ToString()
	assert.Equal(t, "hello", name)

	age := Json(jsonstr).Key("detail").Key("age").ToInt64()
	assert.Equal(t, int64(20), age)

	lang1 := Json(jsonstr).Key("langs").Index(1).ToString()
	assert.Equal(t, "golang", lang1)

	langs1 := Json(jsonstr).Key("langs").ToJson()
	assert.Equal(t, `["php","golang","python","shell"]`, langs1)

	langs2 := Json(jsonstr).Key("langs").ToArray()
	assert.ElementsMatch(t, []interface{}{"python", "shell", "php", "golang"}, langs2.([]interface{}))

	detail := Json(jsonstr).Key("detail").ToArray()
	assert.Equal(t, map[string]interface{}{
		"age":    float64(20),
		"height": "175cm",
		"weight": "60kg",
	}, detail.(map[string]interface{}))

	assert.Nil(t, Json(jsonstr).Key("name").ToArray())
}

func TestJsonError(t *testing.T) {
	assert.Empty(t, Json("{").ToJson())
	assert.Nil(t, Json("{").Key("name").Value())
	assert.Empty(t, Json("{").Key("name").ToString())
	assert.Equal(t, int64(0), Json("{").Key("name").ToInt64())
	assert.Nil(t, Json(jsonstr).Key("name1").Value())
	assert.Nil(t, Json(jsonstr).Key("langs").Key("name").Value())
	assert.Nil(t, Json(jsonstr).Index(0).Value())
	assert.Nil(t, Json(jsonstr).Key("langs").Index(10).Value())
	assert.Equal(t, int64(0), Json(jsonstr).Key("langs").ToInt64())
}

// BenchmarkJson-8   	  327823	      3381 ns/op	    1472 B/op	      37 allocs/op
func BenchmarkJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Json(jsonstr).Key("name").ToString()
	}
}
