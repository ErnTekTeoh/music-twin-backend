package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func CallHTTPEndpointWithHeaders(url, method string, data []byte) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Authorization": "Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkY0QTI4NTQ5UVMifQ.eyJpYXQiOjE3NjU3MTM1NTcsImV4cCI6MTc4MTI2NTU1NywiaXNzIjoiRks1R01VS1kzSyJ9.FmFl-cAOiwWfS2JY4iMUBOu7fVK8920Z65DBvtkjW4Tcil2Dk-Aq7OjkJWtbiBS09CJDH257fuvVTBU3h5XY9g",
	}
	// Add custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CallHTTPEndpointWithAuth(url, method, authHeader string, data []byte) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return "", err
	}

	// Add Authorization header if provided
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// CallHTTPEndpoint makes an HTTP GET or POST request to the specified URL.
// For GET requests, data should be nil. For POST, pass a []byte payload.
func CallHTTPEndpoint(url string, method string, data []byte) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var req *http.Request
	var err error

	if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
	} else { // default to GET
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
