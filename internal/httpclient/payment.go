package httpclient

import (
	"debtreportapp/internal/types"
	"net/http"
)

func (c *DebtClient) GetPayments() ([]types.Payment, error) {
	baseURL := c.baseURL
	baseURL.Path += PaymentsEndpoint
	var out []types.Payment
	err := c.request(http.MethodGet, baseURL.String(), nil, &out)
	return out, err
}
