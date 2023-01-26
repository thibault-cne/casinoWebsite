package roulette

import (
	"bufio"
	"net"
	"net/http"
)

type Writer struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewCustomResponseWriter() *Writer {
	return &Writer{
		header: http.Header{},
	}
}

func (w *Writer) Header() http.Header {
	return w.header
}

func (w *Writer) Write(b []byte) (int, error) {
	w.body = b
	// implement it as per your requirement
	return 0, nil
}

func (w *Writer) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *Writer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// random net Conn
	c := &net.TCPConn{}
	// new bufio readwriter
	rw := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))
	return c, rw, nil
}
