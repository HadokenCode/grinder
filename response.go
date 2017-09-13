package grinder

import (
	"net/http"
)

// Response is the standard Grinder response struct
type Response struct {
	writer    http.ResponseWriter
	Status    int
	Size      int64
	Committed bool
}

// NewResponse creates and returns a new Grinder Response struct
func NewResponse(w http.ResponseWriter) (r *Response) {
	return &Response{writer: w}
}

// Write will write the bytes (message) to the client
func (r *Response) Write(b []byte) (n int, err error) {
	n, err = r.writer.Write(b)
	return
}

// WriteHeader writes a header to the response writer
func (r *Response) WriteHeader(code int) {
	r.writer.WriteHeader(code)
}

// Header will return the header information
func (r *Response) Header() http.Header {
	return r.writer.Header()
}
