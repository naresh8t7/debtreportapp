package httpclient

import (
	"bytes"
	"debtreportapp/internal/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	DebtsEndpoint        = "/debts"
	PaymentsEndpoint     = "/payments"
	PaymentPlansEndpoint = "/payment_plans"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type DebtsAPI interface {
	GetDebts() ([]types.Debt, error)
	GetPaymentPlans() ([]types.PaymentPlans, error)
	GetPayments() ([]types.Payment, error)
}

type DebtClient struct {
	client  HTTPClient
	baseURL url.URL
}

// NewDebtClient retuns a new Debt API interface
func NewDebtClient(baseURL string) (DebtsAPI, error) {

	if baseURL == "" {
		return nil, errors.New("missing BaseURL")
	}
	baseUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: 15 * time.Second,
		},
	}
	debtClient := &DebtClient{
		client:  httpClient,
		baseURL: *baseUrl,
	}
	return debtClient, nil
}

// request intermediate function to format an HTTP request from an interface
func (c *DebtClient) request(method string, urlIn string, in interface{}, out interface{}) (err error) {
	var br io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		br = bytes.NewReader(b)
	}
	r, err := http.NewRequest(method, urlIn, br)
	if err != nil {
		return
	}
	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}
	if resp == nil {
		return errors.New("got nil response")
	}
	var body []byte
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if err = json.Unmarshal(body, out); err != nil {
			return err
		}
	default:
		return fmt.Errorf("got error response from API %v", string(body))
	}
	return nil
}
