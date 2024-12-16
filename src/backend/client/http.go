package client

import (
	"log"
	"net/http"
	"time"
)

type ClientHTTP struct {
	client *http.Client
}

func NewClientHTTP() *ClientHTTP {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	return &ClientHTTP{
		client: client,
	}
}

func (h *ClientHTTP) Get(url string, token string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicf("Error in create request %v", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := h.client.Do(req)
	if err != nil {
		log.Panicf("Error to make request: %v", err.Error())
	}
	return resp
}
