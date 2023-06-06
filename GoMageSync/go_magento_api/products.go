package go_magento_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	magentoBaseURL = "https://your-magento-instance.com"
)

type ProductFilter struct {
	CreatedDate string
	UpdatedDate string
	ProductName string
	ProductCode string
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProductListResponse struct {
	Items []Product `json:"items"`
}

func GetAllProducts(authToken string, filter ProductFilter, page int, pageSize int) ([]Product, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/rest/V1/products", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("searchCriteria[pageSize]", strconv.Itoa(pageSize))
	query.Add("searchCriteria[currentPage]", strconv.Itoa(page))

	if filter.CreatedDate != "" {
		query.Add("searchCriteria[filter_groups][0][filters][0][field]", "created_at")
		query.Add("searchCriteria[filter_groups][0][filters][0][value]", filter.CreatedDate)
		query.Add("searchCriteria[filter_groups][0][filters][0][condition_type]", "gteq")
	}

	if filter.UpdatedDate != "" {
		query.Add("searchCriteria[filter_groups][1][filters][0][field]", "updated_at")
		query.Add("searchCriteria[filter_groups][1][filters][0][value]", filter.UpdatedDate)
		query.Add("searchCriteria[filter_groups][1][filters][0][condition_type]", "gteq")
	}

	if filter.ProductName != "" {
		query.Add("searchCriteria[filter_groups][2][filters][0][field]", "name")
		query.Add("searchCriteria[filter_groups][2][filters][0][value]", "%"+filter.ProductName+"%")
		query.Add("searchCriteria[filter_groups][2][filters][0][condition_type]", "like")
	}

	if filter.ProductCode != "" {
		query.Add("searchCriteria[filter_groups][3][filters][0][field]", "sku")
		query.Add("searchCriteria[filter_groups][3][filters][0][value]", "%"+filter.ProductCode+"%")
		query.Add("searchCriteria[filter_groups][3][filters][0][condition_type]", "like")
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

	var productListResponse ProductListResponse
	err = json.Unmarshal(body, &productListResponse)
	if err != nil {
		return nil, err
	}

	return productListResponse.Items, nil
}

func GetProductDetails(authToken string, productId int) (*Product, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", magentoBaseURL+"/rest/V1/products/"+strconv.Itoa(productId), nil)
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

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}