package context

import (
	"sync"
)

// ResponseWriterPool is the response writer pool, it's used inside router and the framework by itself.
type ResponseWriterPool struct {
	pool    *sync.Pool
	enable  bool
	newFunc func() any
}

// NewResponseWriterPool creates and returns a new response writer pool.
func NewResponseWriterPool(newFunc func() interface{}) *ResponseWriterPool {
	return &ResponseWriterPool{pool: &sync.Pool{New: newFunc}, enable: true, newFunc: newFunc}
}

// Acquire returns a response writer from pool.
// See Release.
func (r *ResponseWriterPool) Acquire() ResponseWriter {
	var w ResponseWriter
	if r.enable {
		w = r.pool.Get().(ResponseWriter)
	} else {
		w = r.newFunc().(ResponseWriter)
	}
	return w
}

// Release puts a response writer back to its pool, this function releases its resources.
// See Acquire.
func (r *ResponseWriterPool) Release(responseWriter ResponseWriter) {
	if r.enable {
		r.pool.Put(responseWriter)
	} else {
		// nothing to do
	}
}

// DisablePool disables the pool.
func (r *ResponseWriterPool) DisablePool() {
	r.enable = false
}
