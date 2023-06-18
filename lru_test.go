package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process(t *testing.T) {
	l := New(2)
	l.Set("a", 12)
	l.Set("b", 13)

	assert.Equal(t, l.Get("a"), 12)

	l.Set("c", "hi")

	assert.Equal(t, l.Get("c"), "hi")
	assert.Equal(t, l.Get("a"), nil)

	l.Set("b", 14)
	l.Set("d", 15)

	assert.Equal(t, l.Get("b"), 14)
	assert.Equal(t, l.Get("c"), nil)
}
