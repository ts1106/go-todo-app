package context

import (
	"encoding/json"
	"net/http"
)

const (
	jsonContentType = "application/json; charset=utf-8"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) JSON(statusCode int, data any) {
	c.Writer.Header().Set("Content-Type", jsonContentType)
	c.Writer.WriteHeader(statusCode)
	json.NewEncoder(c.Writer).Encode(data)
}
