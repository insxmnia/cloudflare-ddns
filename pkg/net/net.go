package net

import (
	"bytes"
	"cloudflare-ddns/pkg/net/models"
	"io"
	"net/http"
)

/*
This is a customised http request package for easier execution and repitition reduction, like the slogger package.
*/

func CreateRequest(method, url string, body []byte) (*http.Request, error) {
	_request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return _request, nil
}

func ExecuteRequest(request *http.Request) *models.MResponse {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &models.MResponse{
			Status:  response.StatusCode,
			Body:    nil,
			Error:   err,
			Request: request,
		}
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	return &models.MResponse{
		Status:   response.StatusCode,
		Body:     body,
		Error:    nil,
		Request:  request,
		Response: response,
	}
}
