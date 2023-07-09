package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"medium/m/v2/internal/product/productdomain/productentities"
	"net/http"
)

type Client struct {
	port int
}

func NewClient() *Client {
	return &Client{
		port: 8081,
	}
}

func (c *Client) Create(product *productentities.Product) (*productentities.Product, error) {
	productBytes, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	requestURL := fmt.Sprintf("http://localhost:%d/products", c.port)

	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(productBytes))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Update(product *productentities.Product) (*productentities.Product, error) {
	productBytes, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, product.ID)

	request, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewBuffer(productBytes))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) GetByID(id string) (*productentities.Product, error) {
	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, id)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Search(productType string) ([]*productentities.Product, error) {
	requestURL := fmt.Sprintf("http://localhost:%d/products?type=%s", c.port, productType)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	response, err := DecodeProducts(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Delete(id string) error {
	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, id)

	request, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	_, err = client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func DecodeProduct(r *http.Response) (*productentities.Product, error) {
	product := &productentities.Product{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func DecodeProducts(r *http.Response) ([]*productentities.Product, error) {
	products := []*productentities.Product{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
