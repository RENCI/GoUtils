package Networking

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	client := GetHttpClientWithNoTLSCheck()
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func HttpPost(url string, body map[string]any) ([]byte, error) {
	client := GetHttpClientWithNoTLSCheck()
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resp_body, err
}

func GetHttpClientWithNoTLSCheck() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}
