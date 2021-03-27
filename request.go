package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	method  string
	Request interface{} `json:"request"`
}

type AuthKey struct {
	Key string `json:"authKey"`
}

func (r *Request) jsonUrl(isProd bool) string {
	if isProd {
		return fmt.Sprintf("https://api.a-3.ru/v1/lightapi/%s/json", r.method)
	}
	return fmt.Sprintf("https://apidev.a-3.ru/v1/lightapi/%s/json", r.method)
}

func newA3Request(request interface{}, method string) *Request {
	return &Request{
		method:  method,
		Request: request,
	}
}

func (a3 *A3) request(jsonReq []byte, r *Request) ([]byte, error) {
	req, err := http.NewRequest("POST", r.jsonUrl(a3.config.isProd), bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a3.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a3 *A3) getPenalties(r *Request) (*GetPenaltiesResponse, error) {
	jsonStr, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("getPenalties json marshal error: %v", err)
	}

	responseBody, err := a3.request(jsonStr, r)
	if err != nil {
		return nil, fmt.Errorf("getPenalties error: %v", err)
	}

	response := new(GetPenaltiesResponse)
	if err = json.Unmarshal(responseBody, response); err != nil {
		return nil, fmt.Errorf("getPenalties json unmarshal error: %v", err)
	}

	return response, nil
}