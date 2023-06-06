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
	magentoBaseURL = "https://your-magento-instance.com/rest/V1"
)

type Order struct {
	Id           int    `json:"id"`
	CreatedDate  string `json:"created_date"`
	UpdatedDate  string `json:"updated_date"`
	CustomerCode string `json:"customer_code"`
	CustomerName string `json:"customer_name"`
}

type OrderListResponse struct {
	Items []Order `json:"items"`
	Total int     `json:"total_count"`
}

func GetAllOrders(authToken string, createdDate, updatedDate, orderId, customerCode, customerName string, page, pageSize int) (*OrderListResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/orders", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)

	query := req.URL.Query()
	query.Add("searchCriteria[currentPage]", strconv.Itoa(page))
	query.Add("searchCriteria[pageSize]", strconv.Itoa(pageSize))

	if createdDate != "" {
		query.Add("searchCriteria[filterGroups][0][filters][0][field]", "created_date")
		query.Add("searchCriteria[filterGroups][0][filters][0][value]", createdDate)
		query.Add("searchCriteria[filterGroups][0][filters][0][conditionType]", "eq")
	}

	if updatedDate != "" {
		query.Add("searchCriteria[filterGroups][1][filters][0][field]", "updated_date")
		query.Add("searchCriteria[filterGroups][1][filters][0][value]", updatedDate)
		query.Add("searchCriteria[filterGroups][1][filters][0][conditionType]", "eq")
	}

	if orderId != "" {
		query.Add("searchCriteria[filterGroups][2][filters][0][field]", "order_id")
		query.Add("searchCriteria[filterGroups][2][filters][0][value]", orderId)
		query.Add("searchCriteria[filterGroups][2][filters][0][conditionType]", "eq")
	}

	if customerCode != "" {
		query.Add("searchCriteria[filterGroups][3][filters][0][field]", "customer_code")
		query.Add("searchCriteria[filterGroups][3][filters][0][value]", customerCode)
		query.Add("searchCriteria[filterGroups][3][filters][0][conditionType]", "eq")
	}

	if customerName != "" {
		query.Add("searchCriteria[filterGroups][4][filters][0][field]", "customer_name")
		query.Add("searchCriteria[filterGroups][4][filters][0][value]", customerName)
		query.Add("searchCriteria[filterGroups][4][filters][0][conditionType]", "eq")
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch orders, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orderListResponse OrderListResponse
	err = json.Unmarshal(body, &orderListResponse)
	if err != nil {
		return nil, err
	}

	return &orderListResponse, nil
}

func GetOrderDetails(authToken string, orderId int) (*Order, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/orders/"+strconv.Itoa(orderId), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch order details, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}