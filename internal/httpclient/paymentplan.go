package httpclient

import (
	"debtreportapp/internal/types"
	"net/http"
)

func (c *DebtClient) GetPaymentPlans() ([]types.PaymentPlans, error) {
	baseURL := c.baseURL
	baseURL.Path += PaymentPlansEndpoint
	var out []types.PaymentPlans
	err := c.request(http.MethodGet, baseURL.String(), nil, &out)
	return out, err
}
