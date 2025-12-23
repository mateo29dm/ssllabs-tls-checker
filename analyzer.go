package main

import (
	"fmt"
	"time"
)

func analyzeHost(host string) (*HostResponse, error) {
	startURL := fmt.Sprintf(
		"%s/analyze?host=%s&startNew=on&all=done",
		baseURL, host,
	)

	resp, _, err := callAnalyze(startURL)
	if err != nil {
		return nil, err
	}

	for {
		fmt.Printf("Status: %s\n", resp.Status)

		switch resp.Status {
		case "READY":
			return resp, nil

		case "ERROR":
			return nil, fmt.Errorf("assessment error: %s", resp.StatusMessage)

		case "DNS":
			time.Sleep(5 * time.Second)

		case "IN_PROGRESS":
			time.Sleep(10 * time.Second)

		default:
			return nil, fmt.Errorf("unknown status: %s", resp.Status)
		}

		pollURL := fmt.Sprintf(
			"%s/analyze?host=%s&all=done",
			baseURL, host,
		)

		resp, _, err = callAnalyze(pollURL)
		if err != nil {
			return nil, err
		}
	}
}
