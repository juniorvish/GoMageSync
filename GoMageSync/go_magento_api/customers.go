package go_magento_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	magentoBaseURL = "https://your-magento-instance-url.com"
)

type Customer struct {
	ID           int       `json:"id"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
	CustomerName string    `json:"customer_name"`
	CustomerCode string    `json:"customer_code"`
}

type CustomerFilter struct {
	CreatedDate  *time.Time
	UpdatedDate  *time.Time
	CustomerName *string
	CustomerCode *string
}

type CustomerPagination struct {
	Page  int
	Limit int
}

func GetAllCustomers(authToken string, filter CustomerFilter, pagination CustomerPagination) ([]Customer, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/rest/V1/customers/search", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("searchCriteria[currentPage]", strconv.Itoa(pagination.Page))
	query.Add("searchCriteria[pageSize]", strconv.Itoa(pagination.Limit))

	if filter.CreatedDate != nil {
		query.Add("searchCriteria[filterGroups][0][filters][0][field]", "created_at")
		query.Add("searchCriteria[filterGroups][0][filters][0][value]", filter.CreatedDate.Format(time.RFC3339))
		query.Add("searchCriteria[filterGroups][0][filters][0][conditionType]", "eq")
	}

	if filter.UpdatedDate != nil {
		query.Add("searchCriteria[filterGroups][1][filters][0][field]", "updated_at")
		query.Add("searchCriteria[filterGroups][1][filters][0][value]", filter.UpdatedDate.Format(time.RFC3339))
		query.Add("searchCriteria[filterGroups][1][filters][0][conditionType]", "eq")
	}

	if filter.CustomerName != nil {
		query.Add("searchCriteria[filterGroups][2][filters][0][field]", "name")
		query.Add("searchCriteria[filterGroups][2][filters][0][value]", *filter.CustomerName)
		query.Add("searchCriteria[filterGroups][2][filters][0][conditionType]", "eq")
	}

	if filter.CustomerCode != nil {
		query.Add("searchCriteria[filterGroups][3][filters][0][field]", "customer_code")
		query.Add("searchCriteria[filterGroups][3][filters][0][value]", *filter.CustomerCode)
		query.Add("searchCriteria[filterGroups][3][filters][0][conditionType]", "eq")
	}

	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customersResponse struct {
		Items []Customer `json:"items"`
	}

	err = json.Unmarshal(body, &customersResponse)
	if err != nil {
		return nil, err
	}

	return customersResponse.Items, nil
}

func GetCustomerDetails(authToken string, customerID int) (*Customer, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+fmt.Sprintf("/rest/V1/customers/%d", customerID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}