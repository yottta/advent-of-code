package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 3, q.Length())
	v, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	assert.Equal(t, 2, q.Length())
	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	assert.Equal(t, 1, q.Length())
	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	assert.Equal(t, 0, q.Length())
	v, ok = q.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
