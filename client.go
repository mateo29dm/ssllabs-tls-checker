package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.ssllabs.com/api/v2"

func callAnalyze(url string) (*HostResponse, int, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var hostResp HostResponse
	if err := json.Unmarshal(body, &hostResp); err != nil {
		return nil, resp.StatusCode, err
	}

	return &hostResp, resp.StatusCode, nil

}
