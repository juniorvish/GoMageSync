package go_magento_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	paymentsEndpoint = "/rest/V1/invoices"
)

type Payment struct {
	EntityID          int    `json:"entity_id"`
	IncrementID       int    `json:"increment_id"`
	OrderID           int    `json:"order_id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	TotalPaid         float64 `json:"total_paid"`
	TotalRefunded     float64 `json:"total_refunded"`
	TotalDue          float64 `json:"total_due"`
}

type PaymentList struct {
	Items []Payment `json:"items"`
}

func GetAllPayments(authToken string, page int, pageSize int) (*PaymentList, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+paymentsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("searchCriteria[currentPage]", strconv.Itoa(page))
	q.Add("searchCriteria[pageSize]", strconv.Itoa(pageSize))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get payments, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var paymentList PaymentList
	err = json.Unmarshal(body, &paymentList)
	if err != nil {
		return nil, err
	}

	return &paymentList, nil
}