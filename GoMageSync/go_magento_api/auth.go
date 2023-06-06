package go_magento_api

import (
	"net/http"
)

const (
	magentoBaseURL = "https://your-magento-instance-url.com"
	authToken      = "your-auth-token"
)

func createRequest(method, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, magentoBaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}