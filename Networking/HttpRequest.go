package Networking

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
)

// One transport for the whole process. Its connection pool is the thing that
// must be shared — a per-call transport leaks idle sockets and their goroutines.
var client = &http.Client{
	Timeout: 5 * time.Minute, // whole request, including body read
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   32, // default is 2 — far too low for one-host workloads
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	},
}

func HttpGet(url string) (int, []byte, error) {
	client := GetHttpClientWithNoTLSCheck()
	resp, err := client.Get(url)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, err
}

func HttpPost(url string, body map[string]any) (int, []byte, error) {
	client := GetHttpClientWithNoTLSCheck()
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return 0, nil, err
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, resp_body, err
}

func HttpDelete(url string) (int, []byte, error) {
	client := GetHttpClientWithNoTLSCheck()

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return 0, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, body, nil
}

func HttpPut(url string, body map[string]any) (int, []byte, error) {
	client := GetHttpClientWithNoTLSCheck()
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, respBody, nil
}

func GetHttpClientWithNoTLSCheck() *http.Client {
	return client
}
