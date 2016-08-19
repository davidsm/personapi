package app

import (
	"log"
	"net/http"
)

type logResponseWriter struct {
	w            http.ResponseWriter
	BytesWritten int
	StatusCode   int
}

func (lrw *logResponseWriter) Header() http.Header {
	return lrw.w.Header()
}

func (lrw *logResponseWriter) Write(data []byte) (int, error) {
	written, err := lrw.w.Write(data)
	lrw.BytesWritten += written
	return written, err
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.w.WriteHeader(code)
}

// This seems like a good idea, although I'm not entirely sure
// if it matters in this case
func (lrw *logResponseWriter) Flush() {
	if wf, ok := lrw.w.(http.Flusher); ok {
		wf.Flush()
	}
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	var lrw logResponseWriter
	lrw.w = w
	// Calls to w.Write might result in a call to WriteHeader on w,
	// in which case it won't get intercepted
	// Assume StatusOK given no explicit WriteHeader calls
	lrw.StatusCode = http.StatusOK
	return &lrw
}

type logHandler struct {
	handler http.Handler
	logger  *log.Logger
}

func (lh *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lrw := newLogResponseWriter(w)
	lh.handler.ServeHTTP(lrw, r)

	var bytesWritten int

	// If the method is HEAD, the body will be written to a "null" writer
	// No body will actually be written to the client
	if r.Method == http.MethodHead {
		bytesWritten = 0
	} else {
		bytesWritten = lrw.BytesWritten
	}
	lh.logger.Printf("%s %s %s - %d - %s - %d",
		r.RemoteAddr,
		r.Method,
		r.URL,
		lrw.StatusCode,
		r.Header.Get("User-Agent"),
		bytesWritten,
	)
}
