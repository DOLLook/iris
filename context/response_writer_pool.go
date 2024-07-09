package context

import (
	"sync"
)

// ResponseWriterPool is the response writer pool, it's used inside router and the framework by itself.
type ResponseWriterPool struct {
	pool *sync.Pool
}

// NewResponseWriterPool creates and returns a new response writer pool.
func NewResponseWriterPool(newFunc func() interface{}) *ResponseWriterPool {
	return &ResponseWriterPool{pool: &sync.Pool{New: newFunc}}
}

// Acquire returns a response writer from pool.
// See Release.
func (c *ResponseWriterPool) Acquire() ResponseWriter {
	w := c.pool.Get().(ResponseWriter)
	return w
}

// Release puts a response writer back to its pool, this function releases its resources.
// See Acquire.
func (c *ResponseWriterPool) Release(responseWriter ResponseWriter) {
	c.pool.Put(responseWriter)
}
