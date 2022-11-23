package api

import "net/http"

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	if e.Message == "" {
		return http.StatusText(e.Code)
	}
	return e.Message
}
