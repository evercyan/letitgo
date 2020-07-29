package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowFlake(t *testing.T) {
	sf, err := NewSnowflake(1)
	assert.NotNil(t, sf)
	assert.Nil(t, err)
	assert.NotEmpty(t, sf.Generate())
}

func TestSnowFlakeCoverage(t *testing.T) {
	_, err := NewSnowflake(1025)
	assert.NotNil(t, err)
}

// BenchmarkSnowFlake-8   	10129148	       117 ns/op	      32 B/op	       1 allocs/op
func BenchmarkSnowFlake(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sf, _ := NewSnowflake(1)
		sf.Generate()
	}
}
