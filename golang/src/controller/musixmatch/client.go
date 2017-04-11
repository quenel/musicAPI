package musixmatch

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiURL = "http://api.musixmatch.com/ws/1.1/"
	apiKey = "d9e1b756c027e7916253cbe5a99d1e27"
)

const (
	apiKeyParam = "apikey"
)

type ResponseContent struct {
	Message struct {
		Body   json.RawMessage `json:"body"`
		Header struct {
			StatusCode int `json:"status_code"`
		} `json:"header"`
	} `json:"message"`
}

func get(req *http.Request, obj interface{}) error {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response error: %s", err)
	}
	defer resp.Body.Close()

	content := ResponseContent{}
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return fmt.Errorf("could not decode response body, error: %s", err)
	}

	if content.Message.Header.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code response, status code: %d", resp.StatusCode)
	}

	err = json.Unmarshal(content.Message.Body, &obj)
	if err != nil {
		return fmt.Errorf("could not decode response content to type, error : %s", err)
	}

	return nil
}

func buildRequest(path string, params map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, apiURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("could not build path: %s, error: %s", path, err)
	}

	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	for name, param := range params {
		q.Add(name, param)
	}
	q.Add(apiKeyParam, apiKey)
	req.URL.RawQuery = q.Encode()

	return req, nil
}
