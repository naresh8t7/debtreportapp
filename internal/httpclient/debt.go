package httpclient

import (
	"debtreportapp/internal/types"
	"net/http"
)

func (c *DebtClient) GetDebts() ([]types.Debt, error) {
	baseURL := c.baseURL
	baseURL.Path += DebtsEndpoint
	var out []types.Debt
	err := c.request(http.MethodGet, baseURL.String(), nil, &out)
	return out, err
}
