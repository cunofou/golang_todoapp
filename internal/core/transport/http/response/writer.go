package response

import "net/http"

type ResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
    return &ResponseWriter{
        ResponseWriter: w,
        StatusCode:     200,
    }
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
    rw.StatusCode = statusCode
    rw.ResponseWriter.WriteHeader(statusCode)
}