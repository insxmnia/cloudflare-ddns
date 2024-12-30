package models

import "net/http"

type MResponse struct {
	Body     []byte
	Status   int
	Request  *http.Request
	Response *http.Response
	Error    error
}
