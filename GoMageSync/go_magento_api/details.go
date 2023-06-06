package go_magento_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProductDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Sku         string `json:"sku"`
	Description string `json:"description"`
}

type CustomerDetails struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type OrderDetails struct {
	ID          int     `json:"id"`
	OrderID     string  `json:"order_id"`
	CustomerID  int     `json:"customer_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}

func GetProductDetails(authToken string, productID int) (*ProductDetails, error) {
	url := fmt.Sprintf("https://magento.example.com/rest/V1/products/%d", productID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productDetails ProductDetails
	err = json.Unmarshal(body, &productDetails)
	if err != nil {
		return nil, err
	}

	return &productDetails, nil
}

func GetCustomerDetails(authToken string, customerID int) (*CustomerDetails, error) {
	url := fmt.Sprintf("https://magento.example.com/rest/V1/customers/%d", customerID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var customerDetails CustomerDetails
	err = json.Unmarshal(body, &customerDetails)
	if err != nil {
		return nil, err
	}

	return &customerDetails, nil
}

func GetOrderDetails(authToken string, orderID int) (*OrderDetails, error) {
	url := fmt.Sprintf("https://magento.example.com/rest/V1/orders/%d", orderID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orderDetails OrderDetails
	err = json.Unmarshal(body, &orderDetails)
	if err != nil {
		return nil, err
	}

	return &orderDetails, nil
}