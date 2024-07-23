package testutils

import (
	"github.com/stretchr/testify/mock"
)

// ArgumentCaptor is the argumentCaptor used to capture arguments for mocked methods like Mockito captor.
// See. https://site.mockito.org/javadoc/current/org/mockito/ArgumentCaptor.html
type ArgumentCaptor[T any] interface {
	Capture() interface{}
	GetValue() T
}

type captor[T any] struct {
	target T
}

// NewArgumentCaptor is the constructor of the ArgumentCaptor
func NewArgumentCaptor[T any]() ArgumentCaptor[T] {
	return &captor[T]{}
}

// Capture returns a mock.MatchedBy to test if the argument is the same as the expected and capture it
func (c *captor[T]) Capture() interface{} {
	return mock.MatchedBy(func(param interface{}) bool {
		if param == nil {
			return false
		}

		valueT, ok := param.(T)
		if !ok {
			return false
		}
		c.target = valueT

		return true
	})
}

// GetValue returns the captured value
func (c *captor[T]) GetValue() T {
	return c.target
}
