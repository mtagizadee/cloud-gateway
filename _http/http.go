package _http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var _client = &http.Client{}

func Post(url string, payload interface{}, headers map[string]string) (*http.Response, error) {
	body, err := body(payload)
	if err != nil {
		return nil, err
	}

	req, err := setup("POST", url, body, headers)
	if err != nil {
		return nil, err
	}

	res, err := send(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := setup("GET", url, nil, headers)
	if err != nil {
		return nil, err
	}

	res, err := send(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func send(req *http.Request) (*http.Response, error) {	
	res, err := _client.Do(req)
	if err != nil {
		return nil, errors.New("could not send request")
	}

	return res, nil
}

func setup(method string, url string, body *bytes.Buffer, headers map[string]string) (*http.Request, error) {
	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	
	if err != nil {
		return nil, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// iterate over headers add add header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

func body(payload interface{}) (*bytes.Buffer, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("could not parse payload")
	}

	return bytes.NewBuffer(body), nil
}

func Read(res *http.Response) (map[string]interface{}, error) {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("could not read response")
	}

	// parse response
	var body map[string]interface{}
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		return nil, errors.New("could not parse response")
	}

	return body, nil
}